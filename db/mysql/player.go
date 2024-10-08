package mysql

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/model"
)

func (b *Mysql) PlayerByCharacterID(ctx context.Context, characterID int64) (*model.Player, error) {
	player := &model.Player{}

	row, err := b.query.PlayerByCharacterID(ctx, uint32(characterID))
	if err != nil {
		return nil, fmt.Errorf("player by character id: %w", err)
	}
	player.ID = row.ID
	player.Name = row.Name
	player.Level = row.Level
	return player, nil
}
