package npc

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

func fetchNpcMerchant(ctx context.Context, id int) (*model.NpcMerchant, error) {
	path := fmt.Sprintf("npc_merchant/%d.yaml", id)
	cacheData, ok := cache.Read(ctx, path)
	if ok {
		cacheNpc, ok := cacheData.(*model.NpcMerchant)
		if !ok {
			return nil, fmt.Errorf("cache read: invalid type, wanted *model.NpcMerchant, got %T", cacheData)
		}
		if cacheNpc != nil {
			return cacheNpc, nil
		}
	}

	query := `SELECT i.id, i.Name name, i.price, i.ldonprice, i.icon
FROM items i, merchantlist ml
WHERE ml.merchantid = :id
AND ml.item = i.id
ORDER BY ml.slot`

	rows, err := db.Query(ctx,
		query,
		map[string]interface{}{
			"id": id,
		})
	if err != nil {
		return nil, fmt.Errorf("query npcs: %w", err)
	}
	defer rows.Close()

	npcMerchant := &model.NpcMerchant{}

	for rows.Next() {
		npcMerchantEntry := &model.NpcMerchantEntry{}
		err = rows.StructScan(npcMerchantEntry)
		if err != nil {
			return nil, fmt.Errorf("rows.StructScan: %w", err)
		}
		npcMerchant.Entries = append(npcMerchant.Entries, npcMerchantEntry)
	}

	err = cache.Write(ctx, path, npcMerchant)
	if err != nil {
		return nil, fmt.Errorf("cache write: %w", err)
	}

	return npcMerchant, nil
}
