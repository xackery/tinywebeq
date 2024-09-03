package store

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/models"
)

var (
	itemSearchMux sync.RWMutex
	itemSearch    map[string]*models.ItemSearch
)

func InitItemSearch(ctx context.Context) error {
	return initItemSearch(ctx, true)
}

func initItemSearch(ctx context.Context, isForcedEnabled bool) error {

	if !config.Get().Item.Search.IsEnabled && !isForcedEnabled {
		return nil
	}

	if !config.Get().Item.Search.IsMemorySearchEnabled && !isForcedEnabled {
		return nil
	}

	totalCount := 0

	itemSearchMux.Lock()
	defer itemSearchMux.Unlock()

	itemSearch = make(map[string]*models.ItemSearch)

	rows, err := db.Mysql.ItemsAll(ctx)
	if err != nil {
		return fmt.Errorf("itemsAll: %w", err)
	}

	for _, row := range rows {
		totalCount++

		level := int32(0)
		if row.Hp > 5 && level < 5 {
			level = 5
		}
		ratio := float64(row.Damage) / float64(row.Delay)
		if ratio > 1.5 && level < 10 {
			level = 10
		}
		if row.Mana > 5 && level < 15 {
			level = 15
		}

		if row.Reqlevel > 0 && level < row.Reqlevel {
			level = row.Reqlevel
		}
		if row.Reclevel > 0 && level < row.Reclevel {
			level = row.Reclevel
		}
		if row.Ac > 0 {
			if row.Ac < 5 && level < 5 {
				level = 5
			}
			if row.Ac < 10 && level < 10 {
				level = 10
			}
			if row.Ac < 15 && level < 15 {
				level = 15
			}
			if row.Ac < 20 && level < 20 {
				level = 20
			}
		}

		if level > int32(config.Get().MaxLevel) {
			continue
		}

		itemSearch[row.Name] = &models.ItemSearch{
			ID:    int64(row.ID),
			Name:  row.Name,
			Level: int64(level),
		}
	}

	return nil
}

func ItemSearchByName(ctx context.Context, name string) ([]*models.ItemSearch, error) {
	if !config.Get().Item.Search.IsEnabled {
		return nil, fmt.Errorf("item search is disabled")
	}

	if !config.Get().Item.Search.IsMemorySearchEnabled {
		return db.Mysql.ItemSearchByName(ctx, name)
	}

	itemSearchMux.RLock()
	defer itemSearchMux.RUnlock()

	var items []*models.ItemSearch

	item, ok := itemSearch[name]
	if ok {
		items = append(items, item)
		return items, nil
	}

	names := strings.Split(name, " ")
	for _, item := range itemSearch {
		for _, n := range names {
			if strings.Contains(strings.ToLower(item.Name), strings.ToLower(n)) {
				items = append(items, item)
				break
			}
		}
	}

	return items, nil
}
