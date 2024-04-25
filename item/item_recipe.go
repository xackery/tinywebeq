package item

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/model"
)

func fetchItemRecipe(ctx context.Context, id int) (*model.ItemRecipe, error) {
	path := fmt.Sprintf("item_recipe/%d.yaml", id)
	cacheData, src, ok := cache.Read(ctx, path)
	if ok {
		cacheItemRecipe := cacheData.(*model.ItemRecipe)

		if cacheItemRecipe != nil {

			if src != cache.SourceCacheMemory {
				err := cache.WriteMemoryCache(ctx, path, cacheItemRecipe)
				if err != nil {
					return nil, fmt.Errorf("cache write: %w", err)
				}
			}

			return cacheItemRecipe, nil
		}
	}

	return nil, fmt.Errorf("not found")
}
