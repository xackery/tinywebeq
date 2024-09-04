package repo

import (
	"context"
	"database/sql"
	"errors"

	"go.uber.org/zap"

	"github.com/xackery/tinywebeq/models"
)

const itemByID = `-- name: ItemByID :one
SELECT * 
FROM items 
WHERE id = ? LIMIT 1
`

// ItemByID returns a full item representation from the database.
// Unsafe calls are used to exclude unknown fields from the scanned data.
func (r *Repo) ItemByID(ctx context.Context, id int64) (*models.Item, error) {
	var err error

	r.logger.Debug("Executing ItemByID", zap.Int64("id", id))

	r.logger.Debug("Switching DB to unsafe")
	db := r.db.Unsafe()

	item := models.Item{}
	if err = db.QueryRowxContext(ctx, itemByID, id).StructScan(&item); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}

		r.logger.Error("error scanning item", err)
		return nil, err
	}

	return &item, nil
}

const itemDiscoveredOnlyByID = `
-- name: ItemDiscoveredOnlyByID :one
SELECT *
FROM items, discovered_items 
WHERE items.id=discovered_items.item_id 
AND discovered_items.item_id=? 
LIMIT 1
`

// DiscoveredItemByID returns an item if it has been discovered by a player.
// Unsafe calls are used to exclude unknown fields from the scanned data.
func (r *Repo) DiscoveredItemByID(ctx context.Context, id int64) (*models.DiscoveredItem, error) {
	var err error

	r.logger.Debug("Executing DiscoveredItemByID ", "id ", id)

	r.logger.Debug("Switching DB to unsafe")
	db := r.db.Unsafe()

	item := models.DiscoveredItem{}
	if err = db.QueryRowxContext(ctx, itemDiscoveredOnlyByID, id).StructScan(&item); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNotFound
		}
		return nil, err
	}

	if item.Item == nil {
		return nil, ErrNotFound
	}

	return &item, nil
}
