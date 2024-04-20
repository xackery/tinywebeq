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

	tlog.Debugf("Loading spells")
	err := initSpells()
	if err != nil {
		return fmt.Errorf("initSpells: %w", err)
	}

	tlog.Debugf("Loading zones")
	err = initZones()
	if err != nil {
		return fmt.Errorf("initZones: %w", err)
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

func (l *Library) SpellInfo(id int) []string {
	return SpellInfo(id)
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
