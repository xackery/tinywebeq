package store

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

// ItemQuestByItemID fetches item quest by item id, first by memory, then by file
func ItemQuestByItemID(ctx context.Context, itemID int64) (*model.ItemQuest, error) {
	itemQuest, err := itemQuestByItemID(ctx, itemID)
	if err != nil {
		return nil, fmt.Errorf("fetch: %w", err)
	}
	for _, entry := range itemQuest.Entries {
		if entry.ZoneID > 0 {
			entry.Zone, err = ZoneByZoneIDNumber(ctx, int64(entry.ZoneID))
			if err != nil {
				return nil, fmt.Errorf("fetch zone %d: %w", entry.ZoneID, err)
			}
		}
	}
	return itemQuest, nil
}

func itemQuestByItemID(ctx context.Context, itemID int64) (*model.ItemQuest, error) {
	itemQuest, err := db.BBolt.ItemQuestByItemID(ctx, itemID)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return nil, fmt.Errorf("query item: %w", err)
	}
	if itemQuest == nil {
		return nil, fmt.Errorf("item not found")
	}
	return itemQuest, nil
}
