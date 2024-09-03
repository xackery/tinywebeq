package main

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/xackery/tinywebeq/handlers"
	"github.com/xackery/tinywebeq/repo"
	"github.com/xackery/tinywebeq/store"
)

func (app *application) itemCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "itemID"), 10, 64)
		if err != nil {
			if errors.Is(err, strconv.ErrSyntax) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			app.logger.Error(zap.Error(err))
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		item, err := app.db.DiscoveredItemByID(r.Context(), id)
		if err != nil {
			if errors.Is(err, repo.ErrNotFound) {
				http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
				return
			}

			app.logger.Error(zap.Error(err))
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			return
		}

		ctx := context.WithValue(r.Context(), handlers.ContextKeyItem, item)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func npcCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "npcID"), 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		npc, err := store.NpcByNpcID(r.Context(), id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), handlers.ContextKeyNPC, npc)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func playerCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "playerID"), 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		player, err := store.PlayerByCharacterID(r.Context(), id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), handlers.ContextKeyPlayer, player)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func questCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "questID"), 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		quest, err := store.QuestByQuestID(r.Context(), id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if quest == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), handlers.ContextKeyQuest, quest)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func spellCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "spellID"), 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		spell := store.SpellByID(int32(id))
		if spell == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), handlers.ContextKeySpell, spell)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func zoneCtx(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		id, err := strconv.ParseInt(chi.URLParam(r, "zoneID"), 10, 64)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		zone, err := store.ZoneByZoneIDNumber(r.Context(), id)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}

		if zone == nil {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			return
		}

		ctx := context.WithValue(r.Context(), handlers.ContextKeyZone, zone)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
