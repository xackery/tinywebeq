package npc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"io/fs"
	"net/http"
	"strconv"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/template"
	"github.com/xackery/tinywebeq/tlog"
)

// Peek handles npc peek requests
func Peek(templates fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var id int64

		if !config.Get().Npc.IsEnabled {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		tlog.Debugf("peek: %s", r.URL.String())

		strID := r.URL.Query().Get("id")
		if len(strID) > 0 {
			id, err = strconv.ParseInt(strID, 10, 64)
			if err != nil {
				tlog.Errorf("strconv.Atoi: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		tlog.Debugf("peekRender: id: %d", id)

		err = peekRender(ctx, templates, id, w)
		if err != nil {
			if err.Error() == "npc not found" {
				http.Error(w, "Not Found", http.StatusNotFound)
				return
			}
			tlog.Errorf("peekRender: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func peekRender(ctx context.Context, templates fs.FS, id int64, w http.ResponseWriter) error {
	npc, err := store.NpcByNpcID(ctx, id)
	if err != nil {
		return fmt.Errorf("store.NpcByNpcID: %w", err)
	}

	if npc.AttackSpeed == 0 {
		npc.AttackSpeed = 100
	}
	if npc.AttackSpeed < 0 {
		npc.AttackSpeed = 100 - npc.AttackSpeed
	}

	var npcLoot *model.NpcLoot
	if npc.LoottableID > 0 {
		npcLoot, err = store.NpcLootByNpcID(ctx, int64(npc.LoottableID))
		if err != nil {
			return fmt.Errorf("store.NpcLootByNpcID: %w", err)
		}

	}
	var npcMerchant *model.NpcMerchant
	if npc.MerchantID > 0 {
		npcMerchant, err = store.NpcMerchantByNpcID(ctx, int64(npc.MerchantID))
		if err != nil {
			return fmt.Errorf("store.NpcMerchantByNpcID: %w", err)
		}
	}

	var npcFaction *model.NpcFaction
	if npc.NpcFactionID > 0 {
		npcFaction, err = store.NpcFactionByFactionID(ctx, int64(npc.NpcFactionID))
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("store.NpcFactionByFactionID: %w", err)
		}
	}

	var npcSpell *model.NpcSpell
	if npc.NpcSpellsID > 0 {
		npcSpell, err = store.NpcSpellByNpcSpellsID(ctx, int64(npc.NpcSpellsID))
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("store.NpcSpellByNpcSpellsID: %w", err)
		}
	}

	npcQuest, err := store.NpcQuestByNpcID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("store.NpcQuestByNpcID: %w", err)
	}

	npcSpawn, err := store.NpcSpawnByNpcID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("store.NpcSpawnByNpcID: %w", err)
	}

	data := struct {
		Site               site.BaseData
		Library            *library.Library
		NpcInfo            []string
		Store              *store.Store
		IsNpcSearchEnabled bool
		Npc                *model.Npc
		NpcLoot            *model.NpcLoot
		NpcMerchant        *model.NpcMerchant
		NpcFaction         *model.NpcFaction
		NpcSpell           *model.NpcSpell
		NpcQuest           *model.NpcQuest
		NpcSpawn           *model.NpcSpawn
	}{
		Site:               site.BaseDataInit("NPC"),
		Library:            library.Instance(),
		IsNpcSearchEnabled: config.Get().Npc.Search.IsEnabled,
		Store:              store.Instance(),
		Npc:                npc,
		NpcLoot:            npcLoot,
		NpcMerchant:        npcMerchant,
		NpcFaction:         npcFaction,
		NpcSpell:           npcSpell,
		NpcQuest:           npcQuest,
		NpcSpawn:           npcSpawn,
	}

	view, err := template.Compile("npc", "npc/peek.go.tmpl", templates)
	if err != nil {
		return fmt.Errorf("template.Compile: %w", err)
	}

	if err = view.ExecuteTemplate(w, "peek.go.tmpl", data); err != nil {
		return fmt.Errorf("peekTemplate.Execute: %w", err)
	}

	return nil
}
