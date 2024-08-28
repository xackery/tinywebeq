package store

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

func ItemByItemID(ctx context.Context, itemID int64) (*model.Item, error) {

	if config.Get().Item.IsDiscoveredOnly {
		item, err := db.Mysql.ItemDiscoveredOnlyByID(ctx, uint32(itemID))
		if err != nil {
			return nil, fmt.Errorf("query items: %w", err)
		}
		return item, nil
	}
	item, err := db.Mysql.ItemByID(ctx, int32(itemID))
	if err != nil {
		return nil, fmt.Errorf("query items: %w", err)
	}
	return item, nil
}
