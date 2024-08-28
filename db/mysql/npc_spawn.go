package mysql

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/model"
)

func (b *Mysql) NpcSpawnByNpcID(ctx context.Context, npcID int64) (*model.NpcSpawn, error) {
	var err error
	npcSpawn := &model.NpcSpawn{}
	cNpcSpawns, err := b.query.NpcSpawnsByNpcID(ctx, int32(npcID))
	if err != nil {
		return nil, fmt.Errorf("npc spawn by npc id: %w", err)
	}

	for _, row := range cNpcSpawns {
		npcSpawnEntry := &model.NpcSpawnEntry{}
		npcSpawnEntry.DecodeNpcSpawns(row)
		npcSpawn.Entries = append(npcSpawn.Entries, npcSpawnEntry)
	}
	return npcSpawn, nil
}
