package mysql

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/models"
)

func (b *Mysql) NpcSpawnByNpcID(ctx context.Context, npcID int64) (*models.NpcSpawn, error) {
	var err error
	npcSpawn := &models.NpcSpawn{}
	cNpcSpawns, err := b.query.NpcSpawnsByNpcID(ctx, int32(npcID))
	if err != nil {
		return nil, fmt.Errorf("npc spawn by npc id: %w", err)
	}

	for _, row := range cNpcSpawns {
		npcSpawnEntry := &models.NpcSpawnEntry{}
		npcSpawnEntry.DecodeNpcSpawns(row)
		npcSpawn.Entries = append(npcSpawn.Entries, npcSpawnEntry)
	}
	return npcSpawn, nil
}
