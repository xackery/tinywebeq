package db

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/db/bbolt"
)

var (
	BBolt *bbolt.BBolt
)

func initBbolt(ctx context.Context) error {
	BBolt = bbolt.New("cache/bolt.db")
	buckets := []string{
		"quest",
		"item_quest",
		"item_recipe",
		"npc_quest",
		"quest",
	}
	for _, bucket := range buckets {
		err := BBolt.BucketCreate(ctx, bucket)
		if err != nil {
			return fmt.Errorf("bucket create %s: %w", bucket, err)
		}
	}
	return nil
}
