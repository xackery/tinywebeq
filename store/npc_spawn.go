package store

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

// NpcSpawnByNpcID fetches item spawn by item id, first by memory, then by cache, then by database
func NpcSpawnByNpcID(ctx context.Context, npcID int64) (*model.NpcSpawn, error) {

	npcSpawn, err := db.Mysql.NpcSpawnByNpcID(ctx, npcID)
	if err != nil {
		return nil, fmt.Errorf("query npcspawn: %w", err)
	}
	return npcSpawn, nil

}
