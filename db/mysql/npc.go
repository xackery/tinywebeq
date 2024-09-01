package mysql

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/model"
)

func (b *Mysql) NpcByNpcID(ctx context.Context, npcID int64) (*model.Npc, error) {
	npc := &model.Npc{}

	row, err := b.query.NpcByNpcID(ctx, int32(npcID))
	if err != nil {
		return nil, fmt.Errorf("npc by npc id: %w", err)
	}
	npc.ID = int32(row.ID)
	npc.Name = row.Name
	npc.Level = row.Level
	npc.DecodeNpcType(row)

	return npc, nil
}

func (b *Mysql) NpcsAll(ctx context.Context) ([]*model.Npc, error) {
	npcs := []*model.Npc{}
	rows, err := b.query.NpcsAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("npcs all: %w", err)
	}
	for _, row := range rows {
		npc := &model.Npc{
			ID:    int32(row.ID),
			Name:  row.Name,
			Level: row.Level,
		}
		npcs = append(npcs, npc)
	}
	return npcs, nil
}

func (b *Mysql) NpcSearchByName(ctx context.Context, name string) ([]*model.Npc, error) {
	npcs := []*model.Npc{}

	rows, err := b.query.NpcSearchByName(ctx, name)
	if err != nil {
		return nil, fmt.Errorf("npc search by name: %w", err)
	}
	for _, row := range rows {
		npc := &model.Npc{
			ID:    int32(row.ID),
			Name:  row.Name,
			Level: row.Level,
		}
		npcs = append(npcs, npc)
	}
	return npcs, nil
}
