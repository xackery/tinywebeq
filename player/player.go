package player

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/db"
)

func fetchPlayer(ctx context.Context, id int) (*db.Player, error) {
	path := fmt.Sprintf("player/%d.yaml", id)
	ok, cacheData := cache.Read(path)
	if ok {
		cachePlayer := cacheData.(*db.Player)
		if cachePlayer != nil {
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

	player := &db.Player{}

	if !rows.Next() {
		return nil, fmt.Errorf("not found")
	}

	err = rows.StructScan(player)
	if err != nil {
		return nil, fmt.Errorf("rows.StructScan: %w", err)
	}

	err = cache.Write(path, player)
	if err != nil {
		return nil, fmt.Errorf("cache write: %w", err)
	}

	return player, nil
}
