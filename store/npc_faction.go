package store

import (
	"context"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

// NpcFactionByFactionID fetches item faction by item id, first by memory, then by cache, then by database
func NpcFactionByFactionID(ctx context.Context, factionID int64) (*model.NpcFaction, error) {
	return db.Mysql.NpcFactionsByFactionID(ctx, factionID)
}
