package mysql

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/model"
)

func (b *Mysql) NpcFactionsByFactionID(ctx context.Context, factionID int64) (*model.NpcFaction, error) {
	var err error
	npcFaction := &model.NpcFaction{}
	cNpcFactions, err := b.query.NpcFactionsByFactionID(ctx, uint32(factionID))
	if err != nil {
		return nil, fmt.Errorf("npc faction by faction id: %w", err)
	}

	for _, row := range cNpcFactions {
		npcFactionEntry := &model.NpcFactionEntry{}
		npcFactionEntry.DecodeNpcFactionsByFactionIDRow(row)
		npcFaction.Entries = append(npcFaction.Entries, npcFactionEntry)
	}
	return npcFaction, nil
}
