package mysql

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/model"
)

func (b *Mysql) NpcMerchantByNpcID(ctx context.Context, npcID int64) (*model.NpcMerchant, error) {
	npcMerchant := &model.NpcMerchant{}
	rows, err := b.query.NpcMerchantsByMerchantID(ctx, int32(npcID))
	if err != nil {
		return nil, fmt.Errorf("npc merchant by npc id: %w", err)
	}
	for _, row := range rows {
		npcMerchantEntry := &model.NpcMerchantEntry{}
		npcMerchantEntry.DecodeNpcMerchantsByMerchantIDRow(row)
		npcMerchant.Entries = append(npcMerchant.Entries, npcMerchantEntry)
	}
	return npcMerchant, nil
}
