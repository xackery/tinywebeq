package npc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
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
	viewTemplate *template.Template
)

func viewInit() error {
	var err error
	viewTemplate = template.New("view")
	viewTemplate, err = viewTemplate.ParseFS(site.TemplateFS(),
		"npc/view.go.tpl",       // data
		"head.go.tpl",           // head
		"header.go.tpl",         // header
		"sidebar.go.tpl",        // sidebar
		"footer.go.tpl",         // footer
		"layout/content.go.tpl", // layout (requires footer, header, head, data)
	)
	if err != nil {
		return fmt.Errorf("template.ParseFS: %w", err)
	}

	return nil
}

// View handles npc view requests
func View(w http.ResponseWriter, r *http.Request) {
	var err error
	var id int64

	if !config.Get().Npc.IsEnabled {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}

	tlog.Debugf("view: %s", r.URL.String())

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

	tlog.Debugf("viewRender: id: %d", id)

	err = viewRender(ctx, id, w)
	if err != nil {
		if err.Error() == "npc not found" {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
		tlog.Errorf("viewRender: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tlog.Debugf("viewRender: id: %d done", id)
}

func viewRender(ctx context.Context, id int64, w http.ResponseWriter) error {

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

	type TemplateData struct {
		Site               site.BaseData
		Npc                *model.Npc
		Library            *library.Library
		NpcInfo            []string
		IsNpcSearchEnabled bool
		NpcLoot            *model.NpcLoot
		NpcMerchant        *model.NpcMerchant
		NpcSpawn           *model.NpcSpawn
		NpcFaction         *model.NpcFaction
		NpcSpell           *model.NpcSpell
		NpcQuest           *model.NpcQuest
	}

	data := TemplateData{
		Site:               site.BaseDataInit(npc.Name),
		Npc:                npc,
		Library:            library.Instance(),
		IsNpcSearchEnabled: config.Get().Npc.Search.IsEnabled,
		NpcLoot:            npcLoot,
		NpcMerchant:        npcMerchant,
		NpcSpawn:           npcSpawn,
		NpcFaction:         npcFaction,
		NpcSpell:           npcSpell,
		NpcQuest:           npcQuest,
	}
	if config.Get().Npc.Preview.IsEnabled {
		data.Site.ImageURL = fmt.Sprintf("/npcs/preview.png?id=%d", id)
	}

	err = viewTemplate.ExecuteTemplate(w, "content.go.tpl", data)
	if err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	return nil
}
