package store

import (
	"context"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/models"
)

// ItemRecipeByItemID fetches item recipe by item id, first by memory, then by cache, then by database
func ItemRecipeByItemID(ctx context.Context, itemID int64) (*models.ItemRecipe, error) {

	itemRecipe, err := db.BBolt.ItemRecipeByItemID(ctx, itemID)
	if err != nil {
		return nil, err
	}

	return itemRecipe, nil
}
