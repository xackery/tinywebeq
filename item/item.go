package item

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
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
	return nil
}

func fetchItem(ctx context.Context, id int) (*db.Item, error) {
	path := fmt.Sprintf("item/%d.yaml", id)
	ok, cacheData := cache.Read(path)
	if ok {
		cacheItem := cacheData.(*db.Item)
		if cacheItem != nil {
			return cacheItem, nil
		}
	}

	query := "SELECT * FROM items WHERE id=:id LIMIT 1"
	if config.Get().IsDiscoveredOnly {
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

	item := &db.Item{}

	if !rows.Next() {
		return nil, fmt.Errorf("not found")
	}

	err = rows.StructScan(item)
	if err != nil {
		return nil, fmt.Errorf("rows.StructScan: %w", err)
	}

	err = cache.Write(path, item)
	if err != nil {
		return nil, fmt.Errorf("cache write: %w", err)
	}

	return item, nil
}
