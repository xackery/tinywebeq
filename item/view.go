package item

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"text/template"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"

	"github.com/xackery/tinywebeq/site"
)

var (
	viewTemplate *template.Template
)

// https://allaclone.wayfarershaven.com/?a=item&id=1004
func viewInit() error {
	var err error
	viewTemplate = template.New("view")
	viewTemplate, err = viewTemplate.ParseFS(site.TemplateFS(),
		"item/view.go.tpl",      // data
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

// View handles item view requests
func View(w http.ResponseWriter, r *http.Request) {
	var (
		err error
		id  int64
	)

	tlog.Debugf("view: %s", r.URL.String())

	id, err = strconv.ParseInt(r.PathValue("itemID"), 10, 64)
	if err != nil {
		tlog.Errorf("strconv.ParseInt: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	item, err := store.ItemByItemID(r.Context(), id)
	if err != nil {
		tlog.Errorf("store.ItemByItemID: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	// JSON API Views
	if r.Header.Get("Accept") == "application/json" {
		viewJSON(w, nil, item)
		return
	}

	ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
	defer cancel()

	tlog.Debugf("viewRender: id: %d", id)

	err = viewRender(ctx, item, w)
	if err != nil {
		tlog.Errorf("viewRender: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
	tlog.Debugf("viewRender: id: %d done", id)
}

func viewJSON(w http.ResponseWriter, headers http.Header, data any) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		tlog.Errorf("json.Encode: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func viewRender(ctx context.Context, item *model.Item, w http.ResponseWriter) error {
	itemQuest, err := store.ItemQuestByItemID(ctx, int64(item.ItemID))
	if err != nil {
		tlog.Debugf("Ignoring err store.ItemQuestByItemID: %v", err)
	}

	itemRecipe, err := store.ItemRecipeByItemID(ctx, int64(item.ItemID))
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		tlog.Debugf("Ignoring err store.ItemRecipeByItemID: %v", err)
	}

	type TemplateData struct {
		Site                site.BaseData
		Item                *model.Item
		Library             *library.Library
		Store               *store.Store
		IsItemSearchEnabled bool
		ItemQuest           *model.ItemQuest
		ItemRecipe          *model.ItemRecipe
	}

	data := TemplateData{
		Site:                site.BaseDataInit(item.Name),
		Item:                item,
		Library:             library.Instance(),
		IsItemSearchEnabled: config.Get().Item.Search.IsEnabled,
		ItemQuest:           itemQuest,
		ItemRecipe:          itemRecipe,
		Store:               store.Instance(),
	}

	if config.Get().Item.Preview.IsEnabled {
		data.Site.ImageURL = fmt.Sprintf("/items/preview.png?id=%d", item.ItemID)
	}

	err = viewTemplate.ExecuteTemplate(w, "content.go.tpl", data)
	if err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	return nil
}
