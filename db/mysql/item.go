package mysql

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/models"
)

func (b *Mysql) ItemByID(ctx context.Context, id int32) (*models.Item, error) {
	if b.query == nil {
		return nil, fmt.Errorf("query is nil")
	}
	var err error
	item := &models.Item{}
	cItem, err := b.query.ItemByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("item by id: %w", err)
	}
	item.DecodeItem(cItem)
	return item, nil
}

func (b *Mysql) ItemDiscoveredOnlyByID(ctx context.Context, itemID uint32) (*models.Item, error) {
	var err error
	item := &models.Item{}
	cItem, err := b.query.ItemDiscoveredOnlyByID(ctx, itemID)
	if err != nil {
		return nil, fmt.Errorf("item discovered only by id: %w", err)
	}
	item.DecodeDiscoveredItem(cItem)
	return item, nil
}

func (b *Mysql) ItemsAll(ctx context.Context) ([]*models.Item, error) {
	var err error
	var items []*models.Item
	cItems, err := b.query.ItemsAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("items all: %w", err)
	}
	for _, cItem := range cItems {
		item := &models.Item{}
		item.DecodeItemsAll(cItem)
		items = append(items, item)
	}
	return items, nil
}

func (b *Mysql) ItemSearchByName(ctx context.Context, name string) ([]*models.ItemSearch, error) {
	var err error
	items := []*models.ItemSearch{}
	cItems, err := b.query.ItemSearchByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("item search by name: %w", err)
	}
	for _, row := range cItems {
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

		item := &models.ItemSearch{
			ID:    int64(row.ID),
			Name:  row.Name,
			Level: 0,
		}
		items = append(items, item)
	}
	return items, nil
}
