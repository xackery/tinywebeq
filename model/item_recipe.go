package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"

	"github.com/xackery/tinywebeq/tlog"
)

type ItemRecipe struct {
	Entries    []*ItemRecipeEntry
	key        string
	expiration int64
}

type ItemRecipeEntry struct {
	ItemID         int    `db:"item_id"`
	RecipeID       int    `db:"recipe_id"`
	RecipeName     string `db:"recipe_name"`
	Tradeskill     int    `db:"tradeskill"`
	Trivial        int    `db:"trivial"`
	IsContainer    int    `db:"is_container"`
	ComponentCount int    `db:"component_count"`
	SuccessCount   int    `db:"success_count"`
}

func (t *ItemRecipe) Identifier() string {
	return "item_recipe"
}

func (t *ItemRecipe) Key() string {
	return t.key
}

func (t *ItemRecipe) SetKey(key string) {
	t.key = key
}

func (t *ItemRecipe) SetExpiration(expiration int64) {
	t.expiration = expiration
}

func (t *ItemRecipe) Expiration() int64 {
	return t.expiration
}

func (t *ItemRecipe) Serialize() string {
	return serialize(t)
}

func (t *ItemRecipe) Deserialize(data string) error {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return fmt.Errorf("base64 decode: %w", err)
	}
	buf := bytes.NewBuffer(decoded)
	d := gob.NewDecoder(buf)

	err = d.Decode(&t)
	if err != nil {
		return fmt.Errorf("gob decode: %w", err)
	}
	return nil
}

func (t *ItemRecipe) ComponentEntries() []*ItemRecipeEntry {
	used := make([]*ItemRecipeEntry, 0)
	for _, entry := range t.Entries {
		if entry.ComponentCount > 0 {
			used = append(used, entry)
		}
	}
	tlog.Debugf("used requested")
	return used
}

func (t *ItemRecipe) SuccessEntries() []*ItemRecipeEntry {
	reward := make([]*ItemRecipeEntry, 0)
	for _, entry := range t.Entries {
		if entry.SuccessCount > 0 {
			reward = append(reward, entry)
		}
	}
	return reward
}
