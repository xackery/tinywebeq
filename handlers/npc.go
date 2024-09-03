package handlers

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"time"

	"go.uber.org/zap"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/image"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/models"
	"github.com/xackery/tinywebeq/site"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"
)

type NpcDecorated struct {
	Npc         *models.Npc
	NpcLoot     *models.NpcLoot
	NpcMerchant *models.NpcMerchant
	NpcSpawn    *models.NpcSpawn
	NpcFaction  *models.NpcFaction
	NpcSpell    *models.NpcSpell
	NpcQuest    *models.NpcQuest
}

func decorateNPC(ctx context.Context, npc *models.Npc) (*NpcDecorated, error) {
	var (
		err         error
		npcLoot     *models.NpcLoot
		npcMerchant *models.NpcMerchant
		npcFaction  *models.NpcFaction
		npcSpell    *models.NpcSpell
	)

	if npc.LoottableID > 0 {
		npcLoot, err = store.NpcLootByNpcID(ctx, int64(npc.LoottableID))
		if err != nil {
			return nil, fmt.Errorf("store.NpcLootByNpcID: %w", err)
		}

	}

	if npc.MerchantID > 0 {
		npcMerchant, err = store.NpcMerchantByNpcID(ctx, int64(npc.MerchantID))
		if err != nil {
			return nil, fmt.Errorf("store.NpcMerchantByNpcID: %w", err)
		}
	}

	if npc.NpcFactionID > 0 {
		npcFaction, err = store.NpcFactionByFactionID(ctx, int64(npc.NpcFactionID))
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("store.NpcFactionByFactionID: %w", err)
		}
	}

	if npc.NpcSpellsID > 0 {
		npcSpell, err = store.NpcSpellByNpcSpellsID(ctx, int64(npc.NpcSpellsID))
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return nil, fmt.Errorf("store.NpcSpellByNpcSpellsID: %w", err)
		}
	}

	npcQuest, err := store.NpcQuestByNpcID(ctx, int64(npc.ID))
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("store.NpcQuestByNpcID: %w", err)
	}

	npcSpawn, err := store.NpcSpawnByNpcID(ctx, int64(npc.ID))
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("store.NpcSpawnByNpcID: %w", err)
	}

	data := NpcDecorated{
		Npc:         npc,
		NpcLoot:     npcLoot,
		NpcMerchant: npcMerchant,
		NpcSpawn:    npcSpawn,
		NpcFaction:  npcFaction,
		NpcSpell:    npcSpell,
		NpcQuest:    npcQuest,
	}

	return &data, nil
}

func (h *Handlers) NPCIndex() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var name string
		var err error
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if !config.Get().Npc.Search.IsEnabled {
			http.NotFound(w, r)
			return
		}

		tlog.Debugf("search: %s", r.URL.String())

		type Response struct {
			Code    int                 `json:"code"`
			Message string              `json:"message"`
			Npcs    []*models.NpcSearch `json:"npcs,omitempty"`
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
		result, err := store.NpcSearchByName(ctx, name)
		if err != nil {
			tlog.Errorf("store.NpcSearchByName: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}

		resp := Response{
			Code:    200,
			Message: fmt.Sprintf("There are %d npcs found", len(result)),
			Npcs:    result,
		}

		err = json.NewEncoder(w).Encode(resp)
		if err != nil {
			tlog.Errorf("json.NewEncoder: %v", err)
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
			return
		}
	}
}

func (h *Handlers) NPCView() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		npc := r.Context().Value(ContextKeyNPC).(*models.Npc)

		if npc.AttackSpeed == 0 {
			npc.AttackSpeed = 100
		}
		if npc.AttackSpeed < 0 {
			npc.AttackSpeed = 100 - npc.AttackSpeed
		}

		dNPC, err := decorateNPC(r.Context(), npc)
		if err != nil {
			h.serverErrorResponse(w, err)
		}

		data := struct {
			Site               site.BaseData
			Library            *library.Library
			NpcInfo            []string
			IsNpcSearchEnabled bool
			*NpcDecorated
		}{
			Site:               site.BaseDataInit(npc.Name),
			Library:            library.Instance(),
			IsNpcSearchEnabled: config.Get().Npc.Search.IsEnabled,
			NpcDecorated:       dNPC,
		}

		if config.Get().Npc.Preview.IsEnabled {
			data.Site.ImageURL = fmt.Sprintf("/npcs/preview.png?id=%d", int64(npc.ID))
		}

		h.render(w, "npc", "view.go.tmpl", "content.go.tmpl", data)
	}
}

func (h *Handlers) NPCPeek() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var err error

		npc := r.Context().Value(ContextKeyNPC).(*models.Npc)

		dNPC, err := decorateNPC(r.Context(), npc)
		if err != nil {
			h.serverErrorResponse(w, err)
		}

		data := struct {
			Site               site.BaseData
			Library            *library.Library
			NpcInfo            []string
			Store              *store.Store
			IsNpcSearchEnabled bool
			*NpcDecorated
		}{
			Site:               site.BaseDataInit("NPC"),
			Library:            library.Instance(),
			IsNpcSearchEnabled: config.Get().Npc.Search.IsEnabled,
			Store:              store.Instance(),
			NpcDecorated:       dNPC,
		}

		h.render(w, "npc", "peek.go.tmpl", "content.go.tmpl", data)
	}
}

func (h *Handlers) NPCImage() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		npc := r.Context().Value(ContextKeyNPC).(*models.Npc)
		dNPC, err := decorateNPC(r.Context(), npc)
		if err != nil {
			h.serverErrorResponse(w, err)
		}

		tags := ""
		if npc.Lastname.Valid && npc.Lastname.String != "" {
			tags += fmt.Sprintf("(%s) ", npc.Lastname.String)
		}

		if dNPC.NpcMerchant != nil {
			tags += "Merchant, "
		}
		if npc.RareSpawn > 0 {
			tags += "Rare "
		}

		if len(tags) > 0 {
			tags = tags[:len(tags)-1]
		}

		lines := []string{
			fmt.Sprintf("%s %s", npc.CleanName(), tags),
			fmt.Sprintf("ID: %d", npc.ID),
			fmt.Sprintf("Lvl %d %s %s", npc.Level, npc.RaceStr(), npc.ClassStr()),

			fmt.Sprintf("%d HP, %d-%d DMG @ %0.1f%%", npc.Hp, npc.Mindmg, npc.Maxdmg, npc.AttackSpeed),
			npc.NpcSpecialAttacksStr(),
		}

		if len(dNPC.NpcQuest.Entries) > 0 {
			lines = append(lines, fmt.Sprintf("Starts %d quests", len(dNPC.NpcQuest.Entries)))
		}

		if len(dNPC.NpcSpawn.Entries) > 0 {
			lines = append(lines, fmt.Sprintf("Spawns at %d locations", len(dNPC.NpcSpawn.Entries)))
		}

		lines = append(lines, "")

		if dNPC.NpcLoot != nil {
			lines = append(lines, fmt.Sprintf("Drops %d items", len(dNPC.NpcLoot.Entries)))
		}

		if dNPC.NpcMerchant != nil {
			lines = append(lines, fmt.Sprintf("Sells %d items", len(dNPC.NpcMerchant.Entries)))
		}

		if dNPC.NpcSpell != nil {
			for i, entry := range dNPC.NpcSpell.Entries {
				if i > 1 {
					break
				}
				_, spellLines := store.SpellInfo(int32(entry.Spellid), int32(npc.Level))
				isSlot := false
				for _, line := range spellLines {
					if strings.HasPrefix(line, "ID: ") {
						continue
					}
					if strings.HasPrefix(line, "Recovery Time: ") {
						continue
					}
					if strings.HasPrefix(line, "Mana: ") {
						continue
					}
					if strings.HasPrefix(line, "Slot") {
						isSlot = true
					}
					if isSlot && !strings.HasPrefix(line, "Slot") {
						break
					}
					if len(line) == 0 {
						continue
					}
					lines = append(lines, line)
				}
			}
			if len(dNPC.NpcSpell.Entries) > 4 {
				lines = append(lines, fmt.Sprintf("... and %d more spells", len(dNPC.NpcSpell.Entries)-4))
			}
		}

		data, err := image.GenerateNpcPreview(int32(npc.Race), lines)
		if err != nil {
			h.serverErrorResponse(w, err)
		}

		_, err = w.Write(data)
		if err != nil {
			h.logger.Error("error writing image", zap.Error(err))
		}
	}
}
