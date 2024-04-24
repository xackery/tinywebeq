package npc

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/site"
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
	var id int

	tlog.Debugf("view: %s", r.URL.String())

	strID := r.URL.Query().Get("id")
	if len(strID) > 0 {
		id, err = strconv.Atoi(strID)
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

func viewRender(ctx context.Context, id int, w http.ResponseWriter) error {

	npc, err := fetchNpc(ctx, id)
	if err != nil {
		return fmt.Errorf("fetchNpc: %w", err)
	}

	var npcLoot *model.NpcLoot
	if npc.Loottableid > 0 {
		npcLoot, err = fetchNpcLoot(ctx, npc.Loottableid)
		if err != nil {
			return fmt.Errorf("fetchNpcLoot: %w", err)
		}
	}
	var npcMerchant *model.NpcMerchant
	if npc.Merchantid > 0 {
		npcMerchant, err = fetchNpcMerchant(ctx, npc.Merchantid)
		if err != nil {
			return fmt.Errorf("fetchNpcMerchant: %w", err)
		}
	}

	var npcFaction *model.NpcFaction
	if npc.Npcfactionid > 0 {
		npcFaction, err = fetchNpcFaction(ctx, npc.Npcfactionid)
		if err != nil {
			return fmt.Errorf("fetchNpcFaction: %w", err)
		}
	}

	npcSpawn, err := fetchNpcSpawn(ctx, id)
	if err != nil {
		return fmt.Errorf("fetchNpcSpawn: %w", err)
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
	}

	data := TemplateData{
		Site:               site.BaseDataInit("Npc View"),
		Npc:                npc,
		Library:            library.Instance(),
		IsNpcSearchEnabled: config.Get().Npc.Search.IsEnabled,
		NpcLoot:            npcLoot,
		NpcMerchant:        npcMerchant,
		NpcSpawn:           npcSpawn,
		NpcFaction:         npcFaction,
	}

	err = viewTemplate.ExecuteTemplate(w, "content.go.tpl", data)
	if err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	return nil
}
