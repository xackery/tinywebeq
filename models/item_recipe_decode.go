package models

import "github.com/xackery/tinywebeq/db/mysql/storage/mysqlc"

func (t *ItemRecipeEntry) DecodeItemRecipeEntry(in mysqlc.ItemRecipeAllRow) {
	t.RecipeID = in.RecipeID
	t.RecipeName = in.RecipeName
	t.Tradeskill = in.Tradeskill
	t.Trivial = in.Trivial
	t.ItemID = in.ItemID
	t.IsContainer = in.IsContainer
	t.ComponentCount = in.ComponentCount
	t.SuccessCount = in.SuccessCount
}
