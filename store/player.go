package store

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/models"
)

func PlayerByCharacterID(ctx context.Context, playerID int64) (*models.Player, error) {

	player, err := db.Mysql.PlayerByCharacterID(ctx, playerID)
	if err != nil {
		return nil, fmt.Errorf("query players: %w", err)
	}
	return player, nil
}
