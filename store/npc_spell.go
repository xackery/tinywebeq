package store

import (
	"context"

	"github.com/xackery/tinywebeq/db"
	"github.com/xackery/tinywebeq/model"
)

// NpcSpellByNpcSpellsID fetches item spell by item id, first by memory, then by cache, then by database
func NpcSpellByNpcSpellsID(ctx context.Context, npcSpellsID int64) (*model.NpcSpell, error) {
	npcSpell, err := db.Mysql.NpcSpellsByNpcSpellsID(ctx, npcSpellsID)
	if err != nil {
		return nil, err
	}
	return npcSpell, nil

}
