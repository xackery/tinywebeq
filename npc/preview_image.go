package npc

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/image"
	"github.com/xackery/tinywebeq/models"
	"github.com/xackery/tinywebeq/store"
	"github.com/xackery/tinywebeq/tlog"
)

// Preview handles npc preview requests
func PreviewImage(w http.ResponseWriter, r *http.Request) {
	var err error
	var id int64
	if !config.Get().Npc.IsEnabled {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	tlog.Debugf("previewImage: %s", r.URL.String())

	strID := r.URL.Query().Get("id")
	if len(strID) > 0 {
		id, err = strconv.ParseInt(strID, 10, 64)
		if err != nil {
			tlog.Errorf("strconv.Atoi: %v", err)
			http.Error(w, "Not Found", http.StatusNotFound)
			return
		}
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	tlog.Debugf("previewImageRender: id: %d", id)

	err = previewImageRender(ctx, id, w)
	if err != nil {
		tlog.Errorf("previewImageRender: %v", err)
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	tlog.Debugf("previewImageRender: id: %d done", id)
}

func previewImageRender(ctx context.Context, id int64, w http.ResponseWriter) error {

	npc, err := store.NpcByNpcID(ctx, id)
	if err != nil {
		return fmt.Errorf("store.NpcByNpcID: %w", err)
	}

	var npcLoot *models.NpcLoot
	if npc.LoottableID > 0 {
		npcLoot, err = store.NpcLootByNpcID(ctx, int64(npc.LoottableID))
		if err != nil {
			return fmt.Errorf("store.NpcLootByNpcID: %w", err)
		}

	}
	var npcMerchant *models.NpcMerchant
	if npc.MerchantID > 0 {
		npcMerchant, err = store.NpcMerchantByNpcID(ctx, int64(npc.MerchantID))
		if err != nil {
			return fmt.Errorf("store.NpcMerchantByNpcID: %w", err)
		}
	}

	var npcSpell *models.NpcSpell
	if npc.NpcSpellsID > 0 {
		npcSpell, err = store.NpcSpellByNpcSpellsID(ctx, int64(npc.NpcSpellsID))
		if err != nil && !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("store.NpcSpellByNpcSpellsID: %w", err)
		}
	}

	npcQuest, err := store.NpcQuestByNpcID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("store.NpcQuestByNpcID: %w", err)
	}

	npcSpawn, err := store.NpcSpawnByNpcID(ctx, id)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("store.NpcSpawnByNpcID: %w", err)
	}

	if npc.AttackSpeed == 0 {
		npc.AttackSpeed = 100
	}
	if npc.AttackSpeed < 0 {
		npc.AttackSpeed = 100 - npc.AttackSpeed
	}

	tags := ""
	if npc.Lastname.Valid && npc.Lastname.String != "" {
		tags += fmt.Sprintf("(%s) ", npc.Lastname.String)
	}
	if npcMerchant != nil {
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

	if len(npcQuest.Entries) > 0 {
		lines = append(lines, fmt.Sprintf("Starts %d quests", len(npcQuest.Entries)))
	}

	if len(npcSpawn.Entries) > 0 {
		lines = append(lines, fmt.Sprintf("Spawns at %d locations", len(npcSpawn.Entries)))
	}

	lines = append(lines, "")

	if npcLoot != nil {
		lines = append(lines, fmt.Sprintf("Drops %d items", len(npcLoot.Entries)))
	}

	if npcMerchant != nil {
		lines = append(lines, fmt.Sprintf("Sells %d items", len(npcMerchant.Entries)))
	}

	if npcSpell != nil {
		for i, entry := range npcSpell.Entries {
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
		if len(npcSpell.Entries) > 4 {
			lines = append(lines, fmt.Sprintf("... and %d more spells", len(npcSpell.Entries)-4))
		}
	}

	data, err := image.GenerateNpcPreview(int32(npc.Race), lines)
	if err != nil {
		return fmt.Errorf("GenerateNpcPreview: %w", err)
	}

	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("w.Write: %w", err)
	}

	return nil
}
