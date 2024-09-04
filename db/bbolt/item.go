package bbolt

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	bolt "go.etcd.io/bbolt"

	"github.com/xackery/tinywebeq/models"
)

func (b *BBolt) ItemQuestByItemID(ctx context.Context, itemID int64) (*models.ItemQuest, error) {
	db, err := b.Open()
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	itemQuest := &models.ItemQuest{}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("item_quest"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		v := b.Get(itob(itemID))
		if v == nil {
			return nil
		}
		return gob.NewDecoder(bytes.NewReader(v)).Decode(itemQuest)
	})
	if err != nil {
		return nil, fmt.Errorf("view: %w", err)
	}
	return itemQuest, nil
}

func (b *BBolt) ItemQuestReplace(ctx context.Context, itemID int64, itemQuest *models.ItemQuest) error {
	db, err := b.Open()
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("item_quest"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		var buf bytes.Buffer
		err := gob.NewEncoder(&buf).Encode(itemQuest)
		if err != nil {
			return fmt.Errorf("encode: %w", err)
		}
		return b.Put(itob(itemID), buf.Bytes())
	})
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	return nil
}

func (b *BBolt) ItemQuestTruncate(ctx context.Context) error {
	db, err := b.Open()
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte("item_quest"))
		if err != nil {
			return fmt.Errorf("delete bucket: %w", err)
		}
		_, err = tx.CreateBucket([]byte("item_quest"))
		if err != nil {
			return fmt.Errorf("create bucket: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	return nil
}

func (b *BBolt) ItemRecipeReplace(ctx context.Context, itemID int64, itemRecipe *models.ItemRecipe) error {
	db, err := b.Open()
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("item_recipe"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		var buf bytes.Buffer
		err := gob.NewEncoder(&buf).Encode(itemRecipe)
		if err != nil {
			return fmt.Errorf("encode: %w", err)
		}
		return b.Put(itob(itemID), buf.Bytes())
	})
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	return nil
}

func (b *BBolt) ItemRecipeByItemID(ctx context.Context, itemID int64) (*models.ItemRecipe, error) {
	db, err := b.Open()
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	itemRecipe := &models.ItemRecipe{}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("item_recipe"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		v := b.Get(itob(itemID))
		if v == nil {
			return nil
		}
		return gob.NewDecoder(bytes.NewReader(v)).Decode(itemRecipe)
	})
	if err != nil {
		return nil, fmt.Errorf("view: %w", err)
	}
	return itemRecipe, nil
}

func (b *BBolt) ItemRecipeTruncate(ctx context.Context) error {
	db, err := b.Open()
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte("item_recipe"))
		if err != nil {
			return fmt.Errorf("delete bucket: %w", err)
		}
		_, err = tx.CreateBucket([]byte("item_recipe"))
		if err != nil {
			return fmt.Errorf("create bucket: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	return nil
}
