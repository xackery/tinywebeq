package store

import (
	"context"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/models"
)

// NpcQuestByNpcID fetches item quest by item id, first by memory, then by cache, then by database
func NpcQuestByNpcID(ctx context.Context, npcID int64) (*models.NpcQuest, error) {
	npcQuest, err := db.BBolt.NpcQuestByNpcID(ctx, npcID)
	if err != nil {
		return nil, err
	}

	return npcQuest, nil
}
