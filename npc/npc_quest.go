package npc

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/cache"
	"github.com/xackery/tinywebeq/model"
)

func fetchNpcQuest(ctx context.Context, id int) (*model.NpcQuest, error) {
	path := fmt.Sprintf("npc_quest/%d.yaml", id)
	cacheData, src, ok := cache.Read(ctx, path)
	if ok {
		cacheNpc, ok := cacheData.(*model.NpcQuest)
		if !ok {
			return nil, fmt.Errorf("cache read: invalid type, wanted *model.NpcQuest, got %T", cacheData)
		}
		if cacheNpc != nil {
			if src != cache.SourceCacheMemory {
				err := cache.WriteMemoryCache(ctx, path, cacheNpc)
				if err != nil {
					return nil, fmt.Errorf("cache write: %w", err)
				}
			}

			return cacheNpc, nil
		}
	}
	return nil, fmt.Errorf("not found")
}
