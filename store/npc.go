package store

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

func NpcByNpcID(ctx context.Context, npcID int64) (*model.Npc, error) {
	npc, err := db.Mysql.NpcByNpcID(ctx, npcID)
	if err != nil {
		return nil, fmt.Errorf("query npc: %w", err)
	}
	return npc, nil
}
