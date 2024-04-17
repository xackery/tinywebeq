package util

import (
	"fmt"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	isInitialized bool
	spellNames    = map[int]string{}
)

func Init() error {
	if isInitialized {
		return nil
	}
	isInitialized = true

	query := "SELECT id, name FROM spells_new"
	rows, err := db.Instance.Query(query)
	if err != nil {
		return fmt.Errorf("query spells: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			return fmt.Errorf("rows.Scan: %w", err)
		}
		spellNames[id] = name
	}
	tlog.Debugf("Loaded %d spells", len(spellNames))
	return nil
}

type Util struct {
}
