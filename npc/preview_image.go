package npc

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/image"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/tlog"
)

// Preview handles npc preview requests
func PreviewImage(w http.ResponseWriter, r *http.Request) {
	var err error
	var id int
	if !config.Get().Npc.IsEnabled {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

	tlog.Debugf("previewImage: %s", r.URL.String())

	strID := r.URL.Query().Get("id")
	if len(strID) > 0 {
		id, err = strconv.Atoi(strID)
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

func previewImageRender(ctx context.Context, id int, w http.ResponseWriter) error {
	npc, err := FetchNpc(ctx, id)
	if err != nil {
		return fmt.Errorf("fetchNpc: %w", err)
	}
	var npcLoot *model.NpcLoot
	if npc.Loottableid > 0 {
		npcLoot, err = fetchNpcLoot(ctx, npc.Loottableid)
		if err != nil {
			return fmt.Errorf("fetchNpcLoot: %w", err)
		}
	}
	var npcMerchant *model.NpcMerchant
	if npc.Merchantid > 0 {
		npcMerchant, err = fetchNpcMerchant(ctx, npc.Merchantid)
		if err != nil {
			return fmt.Errorf("fetchNpcMerchant: %w", err)
		}
	}

	var npcSpell *model.NpcSpell
	if npc.Npcspellsid > 0 {
		npcSpell, err = fetchNpcSpell(ctx, npc.Npcspellsid, npc.Level)
		if err != nil {
			return fmt.Errorf("fetchNpcSpell: %w", err)
		}
	}

	if npc.Attackspeed == 0 {
		npc.Attackspeed = 100
	}
	if npc.Attackspeed < 0 {
		npc.Attackspeed = 100 - npc.Attackspeed
	}

	tags := ""
	if npc.Lastname.Valid && npc.Lastname.String != "" {
		tags += fmt.Sprintf("(%s) ", npc.Lastname.String)
	}
	if npcMerchant != nil {
		tags += "Merchant, "
	}
	if npc.Rarespawn > 0 {
		tags += "Rare "
	}

	if len(tags) > 0 {
		tags = tags[:len(tags)-1]
	}

	lines := []string{
		fmt.Sprintf("%s %s", npc.CleanName(), tags),
		fmt.Sprintf("Lvl %d %s %s", npc.Level, npc.RaceStr(), npc.ClassStr()),

		fmt.Sprintf("%d HP, %d-%d DMG @ %0.1f%%", npc.Hp, npc.Mindmg, npc.Maxdmg, npc.Attackspeed),
		npc.NpcSpecialAttacksStr(),
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
			_, spellLines := library.SpellInfo(entry.Spellid, npc.Level)
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

	data, err := image.GenerateNpcPreview(npc.Race, lines)
	if err != nil {
		return fmt.Errorf("GenerateNpcPreview: %w", err)
	}

	_, err = w.Write(data)
	if err != nil {
		return fmt.Errorf("w.Write: %w", err)
	}

	return nil
}
