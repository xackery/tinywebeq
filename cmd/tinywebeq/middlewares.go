package main

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"

	"github.com/xackery/tinywebeq/handlers"
	"github.com/xackery/tinywebeq/store"
)

func (app *application) itemContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "itemID"), 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		item, err := store.ItemByItemID(r.Context(), id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotAcceptable)
			return
		}

		ctx := context.WithValue(r.Context(), handlers.ContextKeyItem, item)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func (app *application) npcContext(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "npcID"), 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		npc, err := store.NpcByNpcID(r.Context(), id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotAcceptable)
			return
		}

		ctx := context.WithValue(r.Context(), handlers.ContextKeyNPC, npc)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
