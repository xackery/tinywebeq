package npc

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

func fetchNpcLoot(ctx context.Context, id int) (*model.NpcLoot, error) {
	path := fmt.Sprintf("npc_loot/%d.yaml", id)
	cacheData, ok := cache.Read(ctx, path)
	if ok {
		cacheNpc, ok := cacheData.(*model.NpcLoot)
		if !ok {
			return nil, fmt.Errorf("cache read: invalid type, wanted *model.NpcLoot, got %T", cacheData)
		}
		if cacheNpc != nil {
			return cacheNpc, nil
		}
	}

	discoveredFrom := ""
	if config.Get().Item.IsDiscoveredOnly {
		discoveredFrom = ", discovered_items di"
	}
	discoveredAnd := ""
	if config.Get().Item.IsDiscoveredOnly {
		discoveredAnd = "\nAND di.item_id = i.id"
	}

	query := fmt.Sprintf(`SELECT i.id, i.name, i.itemtype, lde.chance, lte.probability, lte.lootdrop_id, lte.multiplier
FROM items i, loottable_entries lte, lootdrop_entries lde%s
WHERE lte.loottable_id = :id
AND lte.lootdrop_id = lde.lootdrop_id
AND lde.item_id = i.id%s`, discoveredFrom, discoveredAnd)

	rows, err := db.Query(ctx,
		query,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return nil, fmt.Errorf("query npcs: %w", err)
	}
	defer rows.Close()

	npcLoot := &model.NpcLoot{}

	for rows.Next() {
		npcLootEntry := &model.NpcLootEntry{}
		err = rows.StructScan(npcLootEntry)
		if err != nil {
			return nil, fmt.Errorf("rows.StructScan: %w", err)
		}
		npcLoot.Entries = append(npcLoot.Entries, npcLootEntry)
	}

	err = cache.Write(ctx, path, npcLoot)
	if err != nil {
		return nil, fmt.Errorf("cache write: %w", err)
	}

	return npcLoot, nil
}
