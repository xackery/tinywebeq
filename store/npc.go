package store

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/models"
)

func NpcByNpcID(ctx context.Context, npcID int64) (*models.Npc, error) {
	npc, err := db.Mysql.NpcByNpcID(ctx, npcID)
	if err != nil {
		return nil, fmt.Errorf("query npc: %w", err)
	}

	return npc, nil
}
