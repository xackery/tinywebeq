package item

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

// Peek handles item peek requests
func Peek(templates fs.FS) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var id int64

		if !config.Get().Item.IsEnabled {
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
			if err.Error() == "item not found" {
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

	item, err := store.ItemByItemID(ctx, id)
	if err != nil {
		return fmt.Errorf("store.ItemByItemID: %w", err)
	}

	itemQuest, err := store.ItemQuestByItemID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("store.ItemQuestByItemID: %w", err)
	}

	itemRecipe, err := store.ItemRecipeByItemID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("store.ItemRecipeByItemID: %w", err)
	}

	data := struct {
		Site                site.BaseData
		Library             *library.Library
		ItemInfo            []string
		IsItemSearchEnabled bool
		Store               *store.Store
		Item                *model.Item
		ItemRecipe          *model.ItemRecipe
		ItemQuest           *model.ItemQuest
	}{
		Site:                site.BaseDataInit("ITEM"),
		Library:             library.Instance(),
		IsItemSearchEnabled: config.Get().Item.Search.IsEnabled,
		Store:               store.Instance(),
		Item:                item,
		ItemRecipe:          itemRecipe,
		ItemQuest:           itemQuest,
	}

	view, err := template.Compile("item", "item/peek.go.tmpl", templates)
	if err != nil {
		return fmt.Errorf("template.Compile: %w", err)
	}

	if err = view.ExecuteTemplate(w, "peek.go.tmpl", data); err != nil {
		return fmt.Errorf("peekTemplate.Execute: %w", err)
	}

	return nil
}
