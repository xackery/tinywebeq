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
	spellSearchMux sync.RWMutex
	spellSearch    = map[string]*models.SpellSearch{}
)

func initSpellSearch(ctx context.Context) error {
	if !config.Get().Spell.Search.IsEnabled {
		return nil
	}

	if !config.Get().Spell.Search.IsMemorySearchEnabled {
		return nil
	}

	spellSearchMux.Lock()
	defer spellSearchMux.Unlock()

	spellSearch = make(map[string]*models.SpellSearch)

	rows, err := db.Mysql.SpellsAll(ctx)
	if err != nil {
		return fmt.Errorf("spellsAll: %w", err)
	}

	for _, sp := range rows {
		level := int32(255)

		for i := 0; i < 16; i++ {
			if sp.Classes[i] > 0 && sp.Classes[i] < 255 {
				newLevel := sp.Classes[i]
				if newLevel >= level {
					continue
				}
				level = newLevel
			}
		}

		spellSearch[sp.Name] = &models.SpellSearch{
			ID:    int64(sp.ID),
			Name:  sp.Name,
			Level: int64(level),
		}
	}
	return nil
}

func SpellSearchByName(ctx context.Context, name string) ([]*models.SpellSearch, error) {
	if !config.Get().Spell.Search.IsEnabled {
		return nil, fmt.Errorf("spell search is disabled")
	}

	names := strings.Split(name, " ")
	if !config.Get().Spell.Search.IsMemorySearchEnabled {
		results := []*models.SpellSearch{}

		rows, err := db.Mysql.SpellSearchByName(ctx, "%"+strings.Join(names, "%")+"%")
		if err != nil {
			return nil, fmt.Errorf("spell search by name: %w", err)
		}
		for _, sp := range rows {
			level := int32(255)

			for i := 0; i < 16; i++ {
				if sp.Classes[i] > 0 && sp.Classes[i] < 255 {
					newLevel := sp.Classes[i]
					if newLevel >= level {
						continue
					}
					level = newLevel
				}
			}
			spell := &models.SpellSearch{
				ID:    int64(sp.ID),
				Name:  sp.Name,
				Level: int64(level),
			}
			results = append(results, spell)
		}
		return results, nil
	}

	spellSearchMux.RLock()
	defer spellSearchMux.RUnlock()

	var spells []*models.SpellSearch

	spell, ok := spellSearch[name]
	if ok {
		spells = append(spells, spell)
		return spells, nil
	}

	for _, spell := range spellSearch {
		for _, n := range names {
			if strings.Contains(strings.ToLower(spell.Name), strings.ToLower(n)) {
				spells = append(spells, spell)
				break
			}
		}
	}

	return spells, nil
}
