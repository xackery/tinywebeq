package store

import (
	"context"
	"database/sql"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

// NpcMerchantByNpcID fetches item merchant by item id, first by memory, then by cache, then by database
func NpcMerchantByNpcID(ctx context.Context, npcID int64) (*model.NpcMerchant, error) {
	npcMerchant, err := db.Mysql.NpcMerchantByNpcID(ctx, npcID)
	if err != nil {
		return nil, err
	}

	if npcMerchant == nil {
		return nil, sql.ErrNoRows
	}
	return npcMerchant, nil
}
