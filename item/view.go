package item

import (
	"context"
	"database/sql"
	"encoding/json"
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

// View handles item view requests
func View(templates fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

		err = viewRender(ctx, templates, item, w)
		if err != nil {
			tlog.Errorf("viewRender: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tlog.Debugf("viewRender: id: %d done", id)
	}
}

func viewJSON(w http.ResponseWriter, headers http.Header, data any) {
	w.Header().Set("Content-Type", "application/json")
	enc := json.NewEncoder(w)
	if err := enc.Encode(data); err != nil {
		tlog.Errorf("json.Encode: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func viewRender(ctx context.Context, templates fs.FS, item *model.Item, w http.ResponseWriter) error {
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

	view, err := template.Compile("item", "item/view.go.tmpl", templates)
	if err != nil {
		tlog.Errorf("template.Compile: %v", err)
		return fmt.Errorf("template.Compile: %v", err)
	}

	err = view.ExecuteTemplate(w, "content.go.tmpl", data)
	if err != nil {
		return fmt.Errorf("viewTemplate.Execute: %w", err)
	}

	return nil
}
