package item

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/model"
)

func fetchItemScore(ctx context.Context, id int) (*model.ItemScore, error) {
	path := fmt.Sprintf("item_score/%d.yaml", id)
	cacheData, src, ok := cache.Read(ctx, path)
	if ok {
		cacheItemScore := cacheData.(*model.ItemScore)

		if cacheItemScore != nil {

			if src != cache.SourceCacheMemory {
				err := cache.WriteMemoryCache(ctx, path, cacheItemScore)
				if err != nil {
					return nil, fmt.Errorf("cache write: %w", err)
				}
			}

			return cacheItemScore, nil
		}
	}

	return nil, fmt.Errorf("not found")
}
