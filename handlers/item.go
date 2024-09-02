package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"go.uber.org/zap"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/template"
)

// ItemIndex returns the index page for item functionality.
func (h *Handlers) ItemIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		view, err := template.Compile("item", "item/index.go.tmpl", h.templates)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		if view == nil {
			return
		}

		data := struct {
			Site                site.BaseData
			Library             *library.Library
			Store               *store.Store
			IsItemSearchEnabled bool
		}{
			Site:                site.BaseDataInit("Item Index"),
			Library:             library.Instance(),
			IsItemSearchEnabled: config.Get().Item.Search.IsEnabled,
			Store:               store.Instance(),
		}

		if err = view.ExecuteTemplate(w, "content.go.tmpl", data); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	}
}

func (h *Handlers) ItemView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			err error
			id  int64
		)

		h.logger.Debug("view: %s", r.URL.String())

		id, err = strconv.ParseInt(r.PathValue("itemID"), 10, 64)
		if err != nil {
			h.serverErrorResponse(w, fmt.Errorf("invalid itemID: %s", r.PathValue("itemID")))
			return
		}

		item, err := store.ItemByItemID(r.Context(), id)
		if err != nil {
			h.serverErrorResponse(w, fmt.Errorf("error fetching item from store: %w", err))
			return
		}

		// JSON API Views
		if r.Header.Get("Accept") == "application/json" {
			viewJSON(w, nil, item)
			return
		}

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		itemQuest, err := store.ItemQuestByItemID(ctx, int64(item.ItemID))
		if err != nil {
			h.logger.Debug("Ignoring err store.ItemQuestByItemID", zap.Error(err))
		}

		itemRecipe, err := store.ItemRecipeByItemID(ctx, int64(item.ItemID))
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			h.logger.Debug("Ignoring err store.ItemRecipeByItemID: %v", zap.Error(err))
		}

		data := struct {
			Site                site.BaseData
			Item                *model.Item
			Library             *library.Library
			Store               *store.Store
			IsItemSearchEnabled bool
			Quests              *model.ItemQuest
			Recipes             *model.ItemRecipe
		}{
			Site:                site.BaseDataInit(item.Name),
			Item:                item,
			Library:             library.Instance(),
			IsItemSearchEnabled: config.Get().Item.Search.IsEnabled,
			Quests:              itemQuest,
			Recipes:             itemRecipe,
			Store:               store.Instance(),
		}

		if config.Get().Item.Preview.IsEnabled {
			data.Site.ImageURL = fmt.Sprintf("/items/preview.png?id=%d", item.ItemID)
		}

		view, err := template.Compile("item", "item/view.go.tmpl", h.templates)
		if err != nil {
			h.serverErrorResponse(w, fmt.Errorf("template.Compile: %w", err))
			return
		}

		err = view.ExecuteTemplate(w, "content.go.tmpl", data)
		if err != nil {
			h.logger.Error("viewTemplate.Execute", zap.Error(err))
			return
		}
	}
}
