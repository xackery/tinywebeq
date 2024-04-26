package quest

import (
	"context"
	"fmt"
	"strings"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/item"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/model"
	"github.com/xackery/tinywebeq/npc"
)

var (
	isInitialized bool
)

func Init() error {
	if isInitialized {
		return nil
	}
	isInitialized = true
	err := viewInit()
	if err != nil {
		return fmt.Errorf("viewInit: %w", err)
	}

	err = searchInit()
	if err != nil {
		return fmt.Errorf("searchInit: %w", err)
	}

	err = questInit()
	if err != nil {
		return fmt.Errorf("questInit: %w", err)
	}
	return nil
}

func FetchQuest(ctx context.Context, id int) (*model.Quest, error) {
	path := fmt.Sprintf("quest/%d.yaml", id)
	cacheData, src, ok := cache.Read(ctx, path)
	if !ok {
		return nil, fmt.Errorf("not found")
	}

	cacheQuest, ok := cacheData.(*model.Quest)
	if !ok {
		return nil, fmt.Errorf("cache read: invalid type, wanted *model.Quest, got %T", cacheData)
	}
	if cacheQuest == nil {
		return nil, fmt.Errorf("cache read: nil")
	}

	isUpdated := false
	for _, entry := range cacheQuest.Entries {
		if entry.NpcID == 0 {
			npc, err := npc.FetchNpcByName(ctx, entry.NpcName, entry.ZoneID)
			if err != nil {
				return nil, fmt.Errorf("fetchNpcByName: %w", err)
			}
			entry.NpcID = npc.ID
			zone := library.ZoneByID(npc.ZoneID())
			if zone != nil {
				entry.Expansion = zone.Expansion
			}
			isUpdated = true
		}
		if entry.ItemName == "" {
			item, err := item.FetchItem(ctx, entry.ItemID)
			if err != nil {
				return nil, fmt.Errorf("fetchItem %d: %w", entry.ItemID, err)
			}
			entry.ItemName = item.Name
			isUpdated = true
		}
	}

	if isUpdated {
		err := cache.WriteSqlite(ctx, path, cacheQuest)
		if err != nil {
			return nil, fmt.Errorf("write sqlite cache: %w", err)
		}

	}

	if src != cache.SourceCacheMemory {
		err := cache.WriteMemoryCache(ctx, path, cacheQuest)
		if err != nil {
			return nil, fmt.Errorf("cache write: %w", err)
		}
	}

	return cacheQuest, nil
}

func FetchQuestByName(ctx context.Context, name string, zoneID int) (*model.Quest, error) {
	name = strings.ReplaceAll(name, "-", "`")

	query := fmt.Sprintf(`SELECT id, name, attack_speed, class, hp, lastname, level, loottable_id, maxdmg, merchant_id, mindmg, quest_faction_id, quest_spells_id, questspecialattks, race, trackable, rare_spawn
FROM quest_types
WHERE name=:name
AND id < %d
AND id > %d 
LIMIT 1`, zoneID*1000+1000, zoneID*1000-1)

	rows, err := db.Query(ctx,
		query,
		map[string]interface{}{
			"name": name,
		})
	if err != nil {
		return nil, fmt.Errorf("query quests: %w", err)
	}
	defer rows.Close()

	quest := &model.Quest{}

	if !rows.Next() {
		return nil, fmt.Errorf("not found")
	}

	err = rows.StructScan(quest)
	if err != nil {
		return nil, fmt.Errorf("rows.StructScan: %w", err)
	}

	path := fmt.Sprintf("quest/%d.yaml", quest.ID)

	err = cache.Write(ctx, path, quest)
	if err != nil {
		return nil, fmt.Errorf("cache write: %w", err)
	}

	return quest, nil
}
