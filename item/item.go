package item

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

var (
	isInitialied bool
)

func Init() error {
	if isInitialied {
		return nil
	}
	isInitialied = true
	err := viewInit()
	if err != nil {
		return fmt.Errorf("viewInit: %w", err)
	}
	err = searchInit()
	if err != nil {
		return fmt.Errorf("searchInit: %w", err)
	}
	return nil
}

func fetchItem(ctx context.Context, id int) (*model.Item, error) {
	path := fmt.Sprintf("item/%d.yaml", id)
	cacheData, src, ok := cache.Read(ctx, path)
	if ok {
		cacheItem := cacheData.(*model.Item)
		if cacheItem != nil {
			if src != cache.SourceCacheMemory {
				err := cache.WriteMemoryCache(ctx, path, cacheItem)
				if err != nil {
					return nil, fmt.Errorf("cache write: %w", err)
				}
			}

			return cacheItem, nil
		}
	}

	query := "SELECT * FROM items WHERE id=:id LIMIT 1"
	if config.Get().Item.IsDiscoveredOnly {
		query += "SELECT * FROM items, discovered items WHERE items.id=:id AND discovered_items.item_id=:id LIMIT 1"
	}

	rows, err := db.Query(ctx,
		query,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return nil, fmt.Errorf("query items: %w", err)
	}
	defer rows.Close()

	item := &model.Item{}

	if !rows.Next() {
		return nil, fmt.Errorf("not found")
	}

	err = rows.StructScan(item)
	if err != nil {
		return nil, fmt.Errorf("rows.StructScan: %w", err)
	}

	err = cache.Write(ctx, path, item)
	if err != nil {
		return nil, fmt.Errorf("cache write: %w", err)
	}

	return item, nil
}
