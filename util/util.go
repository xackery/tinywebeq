package util

import (
	"fmt"
	"sync"
)

var (
	isInitialized bool
	mu            = sync.RWMutex{}
)

func Init() error {
	if isInitialized {
		return nil
	}
	isInitialized = true

	err := initSpells()
	if err != nil {
		return fmt.Errorf("initSpells: %w", err)
	}
	return nil
}

type Util struct {
}
