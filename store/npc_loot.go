package store

import (
	"context"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

// NpcLootByNpcID fetches item loot by item id, first by memory, then by cache, then by database
func NpcLootByNpcID(ctx context.Context, loottableID int64) (*model.NpcLoot, error) {

	npcLoot, err := db.Mysql.NpcLootByLootTableID(ctx, loottableID)
	if err != nil {
		return nil, err
	}

	return npcLoot, nil
}
