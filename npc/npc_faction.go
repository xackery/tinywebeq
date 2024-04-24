package npc

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

func fetchNpcFaction(ctx context.Context, id int) (*model.NpcFaction, error) {
	path := fmt.Sprintf("npc_faction/%d.yaml", id)
	cacheData, ok := cache.Read(ctx, path)
	if ok {
		cacheNpc, ok := cacheData.(*model.NpcFaction)
		if !ok {
			return nil, fmt.Errorf("cache read: invalid type, wanted *model.NpcFaction, got %T", cacheData)
		}
		if cacheNpc != nil {
			return cacheNpc, nil
		}
	}

	query := `SELECT fl.name, fl.id, fe.value
FROM faction_list fl, npc_faction_entries fe
WHERE fe.npc_faction_id = :id
AND fe.faction_id = fl.id
GROUP BY fl.id
ORDER BY fe.value DESC`

	rows, err := db.Query(ctx,
		query,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return nil, fmt.Errorf("query npcs: %w", err)
	}
	defer rows.Close()

	npcFaction := &model.NpcFaction{}

	for rows.Next() {
		npcFactionEntry := &model.NpcFactionEntry{}
		err = rows.StructScan(npcFactionEntry)
		if err != nil {
			return nil, fmt.Errorf("rows.StructScan: %w", err)
		}
		npcFaction.Entries = append(npcFaction.Entries, npcFactionEntry)
	}

	err = cache.Write(ctx, path, npcFaction)
	if err != nil {
		return nil, fmt.Errorf("cache write: %w", err)
	}

	return npcFaction, nil
}
