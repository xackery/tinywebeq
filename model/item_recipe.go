package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"

	"github.com/xackery/tinywebeq/tlog"
)

type ItemRecipe struct {
	ItemID          int64
	Entries         []*ItemRecipeEntry
	CacheKey        string `db:"key"`
	CacheExpiration int64
}

type ItemRecipeEntry struct {
	RecipeID       int32
	RecipeName     string
	Tradeskill     int16
	Trivial        int16
	ItemID         int32
	IsContainer    bool
	ComponentCount int8
	SuccessCount   int8
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

func (t *ItemRecipe) Identifier() string {
	return "item_recipe"
}

func (t *ItemRecipe) Key() string {
	return t.CacheKey
}

func (t *ItemRecipe) SetKey(key string) {
	t.CacheKey = key
}

func (t *ItemRecipe) SetExpiration(expiration int64) {
	t.CacheExpiration = expiration
}

func (t *ItemRecipe) Expiration() int64 {
	return t.CacheExpiration
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
