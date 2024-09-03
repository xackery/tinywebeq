package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"go.uber.org/zap"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/image"
	"github.com/xackery/tinywebeq/models"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"
)

func (h *Handlers) IndexSpells() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var name string
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if !config.Get().Spell.Search.IsEnabled {
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}

		tlog.Debugf("search: %s", r.URL.String())

		type Response struct {
			Code    int                   `json:"code"`
			Message string                `json:"message"`
			Spells  []*models.SpellSearch `json:"spells,omitempty"`
		}

		name = r.URL.Query().Get("name")
		if len(name) < 1 {
			resp := Response{
				Code:    400,
				Message: "Invalid name",
			}
			err = json.NewEncoder(w).Encode(resp)
			if err != nil {
				tlog.Errorf("json.NewEncoder: %v", err)
				http.Error(w, "Internal Server Error", http.StatusInternalServerError)
				return
			}
			return

		}

		// turn %20 and other HTML to normal string formatting
		name, err = url.QueryUnescape(name)
		if err != nil {
			tlog.Errorf("url.QueryUnescape: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
		tlog.Debugf("search: name: %s", name)
		results, err := store.SpellSearchByName(ctx, name)
		if err != nil {
			tlog.Errorf("library.SpellSearchByName: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		resp := Response{
			Code:    200,
			Message: fmt.Sprintf("There are %d spells found", len(results)),
			Spells:  results,
		}

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			tlog.Errorf("json.NewEncoder: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handlers) ViewSpell() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		spell := r.Context().Value(ContextKeySpell).(*models.Spell)

		_, info := store.SpellInfo(spell.ID, 0)

		data := struct {
			Site                 site.BaseData
			Spell                *models.Spell
			SpellInfo            []string
			IsSpellSearchEnabled bool
		}{
			Site:                 site.BaseDataInit(spell.Name),
			Spell:                spell,
			SpellInfo:            info,
			IsSpellSearchEnabled: config.Get().Spell.Search.IsEnabled,
		}

		h.render(w, "spell", "view.go.tmpl", "content.go.tmpl", data)
	}
}

func (h *Handlers) GenerateSpellImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		spell := r.Context().Value(ContextKeySpell).(*models.Spell)

		spellIcon, lines := store.SpellInfo(spell.ID, 0)
		if len(lines) == 0 {
			h.serverErrorResponse(w, fmt.Errorf("no spell info found"))
			return
		}

		data, err := image.GenerateSpellPreview(spellIcon, lines)
		if err != nil {
			h.serverErrorResponse(w, fmt.Errorf("GenerateSpellPreview: %w", err))
			return
		}

		_, err = w.Write(data)
		if err != nil {
			h.logger.Error("error writing image", zap.Error(err))
		}
	}
}
