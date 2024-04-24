package npc

import (
	"context"
	"fmt"

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

func fetchNpc(ctx context.Context, id int) (*model.Npc, error) {
	path := fmt.Sprintf("npc/%d.yaml", id)
	cacheData, ok := cache.Read(ctx, path)
	if ok {
		cacheNpc, ok := cacheData.(*model.Npc)
		if !ok {
			return nil, fmt.Errorf("cache read: invalid type, wanted *model.Npc, got %T", cacheData)
		}
		if cacheNpc != nil {
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
	fmt.Println(npc)

	err = cache.Write(ctx, path, npc)
	if err != nil {
		return nil, fmt.Errorf("cache write: %w", err)
	}

	return npc, nil
}
