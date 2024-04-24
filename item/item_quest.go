package item

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/model"
)

func fetchItemQuest(ctx context.Context, id int) (*model.ItemQuest, error) {
	path := fmt.Sprintf("item_quest/%d.yaml", id)
	cacheData, ok := cache.Read(ctx, path)
	if ok {
		cacheItemQuest := cacheData.(*model.ItemQuest)

		if cacheItemQuest != nil {
			return cacheItemQuest, nil
		}
	}

	return nil, fmt.Errorf("not found")
}
