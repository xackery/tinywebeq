package npc

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

func fetchNpcSpawn(ctx context.Context, id int) (*model.NpcSpawn, error) {
	path := fmt.Sprintf("npc_spawn/%d.yaml", id)
	cacheData, src, ok := cache.Read(ctx, path)
	if ok {
		cacheNpc, ok := cacheData.(*model.NpcSpawn)
		if !ok {
			return nil, fmt.Errorf("cache read: invalid type, wanted *model.NpcSpawn, got %T", cacheData)
		}
		if cacheNpc != nil {
			if src != cache.SourceCacheMemory {
				err := cache.WriteMemoryCache(ctx, path, cacheNpc)
				if err != nil {
					return nil, fmt.Errorf("cache write: %w", err)
				}
			}

			return cacheNpc, nil
		}
	}

	query := `SELECT z.long_name, z.short_name, s2.x, s2.y, s2.z, sg.name AS spawngroup, sg.id AS spawngroupid, s2.respawntime
FROM zone z, spawnentry se, spawn2 s2, spawngroup sg
WHERE se.npcID = :id
AND se.spawngroupID = s2.spawngroupID
AND s2.zone = z.short_name
AND se.spawngroupID = sg.id
ORDER BY z.long_name`

	rows, err := db.Query(ctx,
		query,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return nil, fmt.Errorf("query npcs: %w", err)
	}
	defer rows.Close()

	npcSpawn := &model.NpcSpawn{}

	for rows.Next() {
		npcSpawnEntry := &model.NpcSpawnEntry{}
		err = rows.StructScan(npcSpawnEntry)
		if err != nil {
			return nil, fmt.Errorf("rows.StructScan: %w", err)
		}
		npcSpawn.Entries = append(npcSpawn.Entries, npcSpawnEntry)
	}

	err = cache.Write(ctx, path, npcSpawn)
	if err != nil {
		return nil, fmt.Errorf("cache write: %w", err)
	}

	return npcSpawn, nil
}
