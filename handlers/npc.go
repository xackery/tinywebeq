package handlers

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/models"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/template"
)

func (h *Handlers) NPCIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		h.render(w, "npc", "index.go.tmpl", "content.go.tmpl", nil)
	}
}

func (h *Handlers) NPCView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		npc := r.Context().Value(ContextKeyNPC).(*models.Npc)

		var npcLoot *models.NpcLoot
		if npc.LoottableID > 0 {
			npcLoot, err = store.NpcLootByNpcID(r.Context(), int64(npc.LoottableID))
			if err != nil {
				h.serverErrorResponse(w, fmt.Errorf("store.NpcLootByNpcID: %w", err))
			}

		}
		var npcMerchant *models.NpcMerchant
		if npc.MerchantID > 0 {
			npcMerchant, err = store.NpcMerchantByNpcID(r.Context(), int64(npc.MerchantID))
			if err != nil {
				h.serverErrorResponse(w, fmt.Errorf("store.NpcMerchantByNpcID: %w", err))
			}
		}

		var npcFaction *models.NpcFaction
		if npc.NpcFactionID > 0 {
			npcFaction, err = store.NpcFactionByFactionID(r.Context(), int64(npc.NpcFactionID))
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				h.serverErrorResponse(w, fmt.Errorf("store.NpcFactionByFactionID: %w", err))
			}
		}

		var npcSpell *models.NpcSpell
		if npc.NpcSpellsID > 0 {
			npcSpell, err = store.NpcSpellByNpcSpellsID(r.Context(), int64(npc.NpcSpellsID))
			if err != nil && !errors.Is(err, sql.ErrNoRows) {
				h.serverErrorResponse(w, fmt.Errorf("store.NpcSpellByNpcSpellsID: %w", err))
			}
		}

		npcQuest, err := store.NpcQuestByNpcID(r.Context(), int64(npc.ID))
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			h.serverErrorResponse(w, fmt.Errorf("store.NpcQuestByNpcID: %w", err))
		}

		npcSpawn, err := store.NpcSpawnByNpcID(r.Context(), int64(npc.ID))
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			h.serverErrorResponse(w, fmt.Errorf("store.NpcSpawnByNpcID: %w", err))
		}

		data := struct {
			Site               site.BaseData
			Npc                *models.Npc
			Library            *library.Library
			NpcInfo            []string
			IsNpcSearchEnabled bool
			NpcLoot            *models.NpcLoot
			NpcMerchant        *models.NpcMerchant
			NpcSpawn           *models.NpcSpawn
			NpcFaction         *models.NpcFaction
			NpcSpell           *models.NpcSpell
			NpcQuest           *models.NpcQuest
		}{
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
			data.Site.ImageURL = fmt.Sprintf("/npcs/preview.png?id=%d", int64(npc.ID))
		}

		view, err := template.Compile("npc", "npc/view.go.tmpl", h.templates)
		if err != nil {
			h.serverErrorResponse(w, fmt.Errorf("template.Compile: %w", err))
		}

		if err = view.ExecuteTemplate(w, "content.go.tmpl", data); err != nil {
			h.logger.Error(fmt.Errorf("viewTemplate.Execute: %w", err))
		}
	}
}
