package npc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	peekTemplate *template.Template
)

func peekInit() error {
	var err error
	peekTemplate = template.New("peek")
	peekTemplate, err = peekTemplate.ParseFS(site.TemplateFS(),
		"npc/peek.go.tpl", // data
	)
	if err != nil {
		return fmt.Errorf("template.ParseFS: %w", err)
	}

	return nil
}

// Peek handles npc peek requests
func Peek(w http.ResponseWriter, r *http.Request) {
	var err error
	var ids []int64

	if !config.Get().Npc.IsEnabled {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tlog.Debugf("peek: %s", r.URL.String())

	strID := r.URL.Query().Get("id")
	if len(strID) > 0 {
		id, err := strconv.ParseInt(strID, 10, 64)
		if err != nil {
			tlog.Errorf("strconv.Atoi: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		ids = append(ids, id)
	}
	strIDs := r.URL.Query().Get("ids")
	if len(strIDs) > 0 {
		for _, idStr := range strings.Split(strIDs, ",") {
			id, err := strconv.ParseInt(idStr, 10, 64)
			if err != nil {
				tlog.Errorf("strconv.Atoi: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			ids = append(ids, id)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tlog.Debugf("peekRender: ids: %+v", ids)

	err = peekRender(ctx, ids, w)
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

func peekRender(ctx context.Context, ids []int64, w http.ResponseWriter) error {

	type NpcEntries struct {
		Npc         *model.Npc
		NpcLoot     *model.NpcLoot
		NpcMerchant *model.NpcMerchant
		NpcFaction  *model.NpcFaction
		NpcSpell    *model.NpcSpell
		NpcQuest    *model.NpcQuest
		NpcSpawn    *model.NpcSpawn
	}
	npcEntries := []*NpcEntries{}
	for _, id := range ids {
		npc, err := store.NpcByNpcID(ctx, id)
		if err != nil {
			return fmt.Errorf("store.NpcByNpcID: %w", err)
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

		npcEntries = append(npcEntries, &NpcEntries{
			Npc:         npc,
			NpcLoot:     npcLoot,
			NpcMerchant: npcMerchant,
			NpcFaction:  npcFaction,
			NpcSpell:    npcSpell,
			NpcQuest:    npcQuest,
			NpcSpawn:    npcSpawn,
		})
	}
	type TemplateData struct {
		Site               site.BaseData
		Entries            []*NpcEntries
		Library            *library.Library
		NpcInfo            []string
		IsNpcSearchEnabled bool
	}

	data := TemplateData{
		Site:               site.BaseDataInit("NPC"),
		Library:            library.Instance(),
		IsNpcSearchEnabled: config.Get().Npc.Search.IsEnabled,
		Entries:            npcEntries,
	}

	err := peekTemplate.ExecuteTemplate(w, "peek.go.tpl", data)
	if err != nil {
		return fmt.Errorf("peekTemplate.Execute: %w", err)
	}

	return nil
}
