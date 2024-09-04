package mysql

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/models"
)

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
