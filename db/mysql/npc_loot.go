package mysql

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/models"
)

func (b *Mysql) NpcLootByLootTableID(ctx context.Context, lootTableID int64) (*models.NpcLoot, error) {
	var err error
	npcLoot := &models.NpcLoot{}
	cNpcLoots, err := b.query.NpcLootsByLootTableID(ctx, uint32(lootTableID))
	if err != nil {
		return nil, fmt.Errorf("npc loot by loot table id: %w", err)
	}

	for _, row := range cNpcLoots {
		nle := &models.NpcLootEntry{}
		nle.DecodeNpcLoots(row)

		npcLoot.Entries = append(npcLoot.Entries, nle)
	}
	return npcLoot, nil
}
