package mysql

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/model"
)

func (b *Mysql) NpcSpellsByNpcSpellsID(ctx context.Context, npcSpellID int64) (*model.NpcSpell, error) {
	var err error
	npcSpell := &model.NpcSpell{}
	cNpcSpells, err := b.query.NpcSpellsByNpcSpellID(ctx, uint32(npcSpellID))
	if err != nil {
		return nil, fmt.Errorf("npc spell by npc id: %w", err)
	}

	for _, row := range cNpcSpells {
		npcSpellEntry := &model.NpcSpellEntry{}
		npcSpellEntry.DecodeNpcSpells(row)
		npcSpell.Entries = append(npcSpell.Entries, npcSpellEntry)
	}
	return npcSpell, nil
}
