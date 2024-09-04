package handlers

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/image"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/models"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/template"
)

//// Search handles item search requests
//func Search(w http.ResponseWriter, r *http.Request) {
//	var name string
//	var err error
//	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
//	defer cancel()
//
//	if !config.Get().Item.Search.IsEnabled {
//		http.Error(w, "Not Found", http.StatusNotFound)
//		return
//	}
//
//	tlog.Debugf("search: %s", r.URL.String())
//
//	type Response struct {
//		Code    int                 `json:"code"`
//		Message string              `json:"message"`
//		Items   []*model.ItemSearch `json:"items,omitempty"`
//	}
//
//	name = r.URL.Query().Get("name")
//	if len(name) < 1 {
//		resp := Response{
//			Code:    400,
//			Message: "Invalid name",
//		}
//		err = json.NewEncoder(w).Encode(resp)
//		if err != nil {
//			tlog.Errorf("json.NewEncoder: %v", err)
//			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//			return
//		}
//		return
//
//	}
//
//	// turn %20 and other HTML to normal string formatting
//	name, err = url.QueryUnescape(name)
//	if err != nil {
//		tlog.Errorf("url.QueryUnescape: %v", err)
//		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//		return
//	}
//	tlog.Debugf("search: name: %s", name)
//	result, err := store.ItemSearchByName(ctx, name)
//	if err != nil {
//		tlog.Errorf("library.ItemSearchByName: %v", err)
//		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//		return
//	}
//
//	resp := Response{
//		Code:    200,
//		Message: fmt.Sprintf("There are %d items found", len(result)),
//		Items:   result,
//	}
//
//	err = json.NewEncoder(w).Encode(resp)
//	if err != nil {
//		tlog.Errorf("json.NewEncoder: %v", err)
//		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
//		return
//	}
//
//}

type ItemDecorated struct {
	Item    *models.Item
	Quests  *models.ItemQuest
	Recipes *models.ItemRecipe
}

func decorateItem(ctx context.Context, item *models.Item) (*ItemDecorated, error) {
	itemQuest, err := store.ItemQuestByItemID(ctx, int64(item.ID))
	if err != nil {
		// TODO: better error handling here
	}

	itemRecipe, err := store.ItemRecipeByItemID(ctx, int64(item.ID))
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		// TODO: better error handling here
	}

	return &ItemDecorated{
		Item:    item,
		Quests:  itemQuest,
		Recipes: itemRecipe,
	}, nil
}

// IndexItems returns the index page for item functionality.
func (h *Handlers) IndexItems() http.HandlerFunc {
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

		h.render(w, "item", "index.go.tmpl", "content.go.tmpl", data)
	}
}

func (h *Handlers) ViewItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var item *models.Item

		switch r.Context().Value(ContextKeyItem).(type) {
		case *models.Item:
			item = r.Context().Value(ContextKeyItem).(*models.Item)
		case *models.DiscoveredItem:
			item = r.Context().Value(ContextKeyItem).(*models.DiscoveredItem).Item
		}

		// JSON API Views
		if r.Header.Get("Accept") == "application/json" {
			viewJSON(w, nil, item)
			return
		}

		dItem, err := decorateItem(r.Context(), item)
		if err != nil {
			h.serverErrorResponse(w, err)
		}

		data := struct {
			Site                site.BaseData
			Library             *library.Library
			Store               *store.Store
			IsItemSearchEnabled bool
			*ItemDecorated
		}{
			Site:                site.BaseDataInit(item.Name),
			Library:             library.Instance(),
			IsItemSearchEnabled: config.Get().Item.Search.IsEnabled,
			Store:               store.Instance(),
			ItemDecorated:       dItem,
		}

		h.render(w, "item", "view.go.tmpl", "content.go.tmpl", data)
	}
}

func (h *Handlers) PeekItem() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !config.Get().Item.IsEnabled {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		item := r.Context().Value(ContextKeyItem).(*models.Item)

		dItem, err := decorateItem(r.Context(), item)
		if err != nil {
			h.serverErrorResponse(w, err)
		}

		data := struct {
			Site                site.BaseData
			Library             *library.Library
			ItemInfo            []string
			IsItemSearchEnabled bool
			Store               *store.Store
			*ItemDecorated
		}{
			Site:                site.BaseDataInit("ITEM"),
			Library:             library.Instance(),
			IsItemSearchEnabled: config.Get().Item.Search.IsEnabled,
			Store:               store.Instance(),
			ItemDecorated:       dItem,
		}

		h.render(w, "item", "peek.go.tmpl", "peek.go.tmpl", data)
	}
}

// GenerateItemImage handles item preview requests
func (h *Handlers) GenerateItemImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error
		var id int64
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		item := r.Context().Value(ContextKeyItem).(*models.Item)

		ctx, cancel := context.WithTimeout(r.Context(), 5*time.Second)
		defer cancel()

		itemQuest, err := store.ItemQuestByItemID(ctx, id)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			fmt.Printf("error: %v\n", err)
			h.serverErrorResponse(w, fmt.Errorf("store.ItemQuestByItemID: %w", err))
			return
		}

		itemRecipe, err := store.ItemRecipeByItemID(ctx, id)
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			h.serverErrorResponse(w, fmt.Errorf("store.ItemRecipeByItemID: %w", err))
			return
		}

		data, err := image.GenerateItemPreview(item, itemQuest, itemRecipe)
		if err != nil {
			h.serverErrorResponse(w, fmt.Errorf("GenerateItemPreview: %w", err))
		}

		w.Header().Set("Content-Type", "image/png")
		if _, err = w.Write(data); err != nil {
			h.serverErrorResponse(w, err)
		}
	}
}
