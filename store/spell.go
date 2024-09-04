package store

import (
	"context"
	"fmt"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/models"
	"github.com/xackery/tinywebeq/tlog"
)

func SpellName(id int32) string {
	spell := SpellByID(id)
	if spell == nil {
		return fmt.Sprintf("Unknown Spell (%d)", id)
	}
	return spell.Name
}

func SpellByID(id int32) *models.Spell {
	spell, err := db.Mysql.SpellByID(context.Background(), int64(id))
	if err != nil {
		tlog.Debugf("SpellByID %d: %v", id, err)
		return nil
	}
	return spell
}
