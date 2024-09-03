package mysql

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/models"
)

func (b *Mysql) NpcMerchantByNpcID(ctx context.Context, npcID int64) (*models.NpcMerchant, error) {
	npcMerchant := &models.NpcMerchant{}
	rows, err := b.query.NpcMerchantsByMerchantID(ctx, int32(npcID))
	if err != nil {
		return nil, fmt.Errorf("npc merchant by npc id: %w", err)
	}
	for _, row := range rows {
		npcMerchantEntry := &models.NpcMerchantEntry{}
		npcMerchantEntry.DecodeNpcMerchantsByMerchantIDRow(row)
		npcMerchant.Entries = append(npcMerchant.Entries, npcMerchantEntry)
	}
	return npcMerchant, nil
}
