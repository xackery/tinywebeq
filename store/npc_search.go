package store

import (
	"context"
	"fmt"
	"strings"
	"sync"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/models"
)

var (
	npcSearchMux sync.RWMutex
	npcSearch    = map[string]*models.NpcSearch{}
)

func initNpcSearch(ctx context.Context) error {
	if !config.Get().Npc.Search.IsEnabled {
		return nil
	}

	if !config.Get().Npc.Search.IsMemorySearchEnabled {
		return nil
	}

	npcSearchMux.Lock()
	defer npcSearchMux.Unlock()

	npcSearch = make(map[string]*models.NpcSearch)

	rows, err := db.Mysql.NpcsAll(ctx)
	if err != nil {
		return fmt.Errorf("npcsAll: %w", err)
	}

	for _, row := range rows {
		npcSearch[row.Name] = &models.NpcSearch{
			ID:    int64(row.ID),
			Name:  row.Name,
			Level: int64(row.Level),
		}
	}
	return nil
}

func NpcSearchByName(ctx context.Context, name string) ([]*models.NpcSearch, error) {
	if !config.Get().Npc.Search.IsEnabled {
		return nil, fmt.Errorf("npc search is disabled")
	}

	if !config.Get().Npc.Search.IsMemorySearchEnabled {
		results := []*models.NpcSearch{}

		rows, err := db.Mysql.NpcSearchByName(ctx, name)
		if err != nil {
			return nil, fmt.Errorf("npc search by name: %w", err)
		}
		for _, row := range rows {
			npc := &models.NpcSearch{
				ID:    int64(row.ID),
				Name:  row.Name,
				Level: int64(row.Level),
			}
			results = append(results, npc)
		}
		return results, nil
	}

	npcSearchMux.RLock()
	defer npcSearchMux.RUnlock()

	var npcs []*models.NpcSearch

	npc, ok := npcSearch[name]
	if ok {
		npcs = append(npcs, npc)
		return npcs, nil
	}

	names := strings.Split(name, " ")
	for _, npc := range npcSearch {
		for _, n := range names {
			if strings.Contains(strings.ToLower(npc.Name), strings.ToLower(n)) {
				npcs = append(npcs, npc)
				break
			}
		}
	}

	return npcs, nil
}
