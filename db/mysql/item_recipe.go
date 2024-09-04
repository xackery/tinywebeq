package mysql

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/models"
)

func (b *Mysql) ItemRecipeAll(ctx context.Context) ([]*models.ItemRecipe, error) {
	var err error
	var itemRecipes []*models.ItemRecipe
	cItemRecipes, err := b.query.ItemRecipeAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("item recipe all: %w", err)
	}

	itemRecipe := &models.ItemRecipe{}
	itemID := int32(0)
	for _, row := range cItemRecipes {
		ire := &models.ItemRecipeEntry{}
		ire.DecodeItemRecipeEntry(row)

		if itemID == 0 {
			itemID = ire.ItemID
			itemRecipe.ItemID = int64(itemID)
		}

		if itemID == ire.ItemID {
			itemRecipe.Entries = append(itemRecipe.Entries, ire)
			continue
		}

		// start a new recipe entry
		itemID = ire.ItemID
		itemRecipe = &models.ItemRecipe{
			ItemID: int64(ire.ItemID),
		}
		itemRecipe.Entries = append(itemRecipe.Entries, ire)
	}
	return itemRecipes, nil
}
