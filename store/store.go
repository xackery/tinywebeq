// store on load creates in memory lookups of commonly accessed data
package store

import (
	"context"
	"fmt"
)

func Init(ctx context.Context) error {

	type createFunc struct {
		name string
		fn   func(context.Context) error
	}

	createFuncs := []createFunc{
		{name: "item_icon", fn: initItemIcon},
		{name: "npc_icon", fn: initNpcIcon},
		{name: "quest_icon", fn: initQuestIcon},
		{name: "spell_icon", fn: initSpellIcon},
		{name: "spell_search", fn: initSpellSearch},
		{name: "zone_icon", fn: initZoneIcon},
		{name: "zone_search", fn: initZoneSearch},
	}
	for _, create := range createFuncs {
		err := create.fn(ctx)
		if err != nil {
			return fmt.Errorf("init %s: %w", create.name, err)
		}
	}

	return nil
}

type Store struct {
}

func Instance() *Store {
	return &Store{}
}

func (s *Store) SpellName(id int32) string {
	return SpellName(id)
}
