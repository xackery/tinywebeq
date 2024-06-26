package player

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

func fetchPlayer(ctx context.Context, id int) (*model.Player, error) {
	path := fmt.Sprintf("player/%d.yaml", id)
	cacheData, src, ok := cache.Read(ctx, path)
	if ok {
		cachePlayer := cacheData.(*model.Player)
		if cachePlayer != nil {
			if src != cache.SourceCacheMemory {
				err := cache.WriteMemoryCache(ctx, path, cachePlayer)
				if err != nil {
					return nil, fmt.Errorf("cache write: %w", err)
				}
			}

			return cachePlayer, nil
		}
	}

	query := "SELECT * FROM character_data WHERE id=:id LIMIT 1"

	rows, err := db.Query(ctx,
		query,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return nil, fmt.Errorf("query players: %w", err)
	}
	defer rows.Close()

	player := &model.Player{}

	if !rows.Next() {
		return nil, fmt.Errorf("not found")
	}

	err = rows.StructScan(player)
	if err != nil {
		return nil, fmt.Errorf("rows.StructScan: %w", err)
	}

	err = cache.Write(ctx, path, player)
	if err != nil {
		return nil, fmt.Errorf("cache write: %w", err)
	}

	return player, nil
}
