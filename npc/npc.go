package npc

import (
	"context"
	"fmt"
	"strings"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
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
	return nil
}

func FetchNpc(ctx context.Context, id int) (*model.Npc, error) {
	path := fmt.Sprintf("npc/%d.yaml", id)
	cacheData, src, ok := cache.Read(ctx, path)
	if ok {
		cacheNpc, ok := cacheData.(*model.Npc)
		if !ok {
			return nil, fmt.Errorf("cache read: invalid type, wanted *model.Npc, got %T", cacheData)
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

	query := "SELECT id, name, attack_speed, class, hp, lastname, level, loottable_id, maxdmg, merchant_id, mindmg, npc_faction_id, npc_spells_id, npcspecialattks, race, trackable, rare_spawn FROM npc_types WHERE id=:id LIMIT 1"

	rows, err := db.Query(ctx,
		query,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return nil, fmt.Errorf("query npcs: %w", err)
	}
	defer rows.Close()

	npc := &model.Npc{}

	if !rows.Next() {
		return nil, fmt.Errorf("not found")
	}

	err = rows.StructScan(npc)
	if err != nil {
		return nil, fmt.Errorf("rows.StructScan: %w", err)
	}

	err = cache.Write(ctx, path, npc)
	if err != nil {
		return nil, fmt.Errorf("cache write: %w", err)
	}

	return npc, nil
}

func FetchNpcByName(ctx context.Context, name string, zoneID int) (*model.Npc, error) {
	name = strings.ReplaceAll(name, "-", "`")

	query := fmt.Sprintf(`SELECT id, name, attack_speed, class, hp, lastname, level, loottable_id, maxdmg, merchant_id, mindmg, npc_faction_id, npc_spells_id, npcspecialattks, race, trackable, rare_spawn
FROM npc_types
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
		return nil, fmt.Errorf("query npcs: %w", err)
	}
	defer rows.Close()

	npc := &model.Npc{}

	if !rows.Next() {
		return nil, fmt.Errorf("not found")
	}

	err = rows.StructScan(npc)
	if err != nil {
		return nil, fmt.Errorf("rows.StructScan: %w", err)
	}

	path := fmt.Sprintf("npc/%d.yaml", npc.ID)

	err = cache.Write(ctx, path, npc)
	if err != nil {
		return nil, fmt.Errorf("cache write: %w", err)
	}

	return npc, nil
}
