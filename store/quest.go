package store

import (
	"context"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/models"
)

// QuestNextQuestID returns the next available quest id
func QuestNextQuestID(ctx context.Context) (int64, error) {
	return db.BBolt.QuestNextID(ctx)
}

// QuestByID returns a quest by id
func QuestByQuestID(ctx context.Context, questID int64) (*models.Quest, error) {
	return db.BBolt.QuestByQuestID(ctx, questID)
}
