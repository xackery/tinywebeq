package item

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/model"
)

func fetchItemQuest(ctx context.Context, id int) (*model.ItemQuest, error) {
	path := fmt.Sprintf("item_quest/%d.yaml", id)
	cacheData, src, ok := cache.Read(ctx, path)
	if ok {
		cacheItemQuest := cacheData.(*model.ItemQuest)

		if cacheItemQuest != nil {
			if src != cache.SourceCacheMemory {
				err := cache.WriteMemoryCache(ctx, path, cacheItemQuest)
				if err != nil {
					return nil, fmt.Errorf("cache write: %w", err)
				}
			}

			return cacheItemQuest, nil
		}
	}

	return nil, fmt.Errorf("not found")
}
