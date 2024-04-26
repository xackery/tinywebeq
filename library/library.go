package library

import (
	"fmt"
	"sync"

	"github.com/xackery/tinywebeq/tlog"
)

var (
	isInitialized bool
	mu            = sync.RWMutex{}
	instance      *Library
)

func Init() error {
	if isInitialized {
		return nil
	}
	isInitialized = true

	tlog.Debugf("Loading library spells")
	err := initSpells()
	if err != nil {
		return fmt.Errorf("initSpells: %w", err)
	}

	tlog.Debugf("Loading library zones")
	err = initZones()
	if err != nil {
		return fmt.Errorf("initZones: %w", err)
	}

	tlog.Debugf("Loading library items")
	err = initItems()
	if err != nil {
		return fmt.Errorf("initItems: %w", err)
	}

	tlog.Debugf("Loading library npcs")
	err = initNpcs()
	if err != nil {
		return fmt.Errorf("initNpcs: %w", err)
	}

	err = initQuests()
	if err != nil {
		return fmt.Errorf("initQuests: %w", err)
	}
	return nil
}

type Library struct {
}

func Instance() *Library {
	if instance == nil {
		instance = &Library{}
	}
	return instance
}

func (l *Library) SpellInfo(id int, level int) (int, []string) {
	return SpellInfo(id, level)
}

func (l *Library) SpellName(id int) string {
	return SpellName(id)
}

func (l *Library) ZoneLongNameByID(id int) string {
	return ZoneLongNameByID(id)
}

func (l *Library) ZoneLongNameByShortName(shortName string) string {
	return ZoneLongNameByShortName(shortName)
}
