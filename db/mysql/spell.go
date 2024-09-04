package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/xackery/tinywebeq/models"
)

func (b *Mysql) SpellsAll(ctx context.Context) ([]*models.Spell, error) {
	spells := []*models.Spell{}
	rows, err := b.query.SpellsAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("spells all: %w", err)
	}
	for _, row := range rows {
		spell := &models.Spell{}
		spell.DecodeSpellsNew(row)
		spells = append(spells, spell)
	}
	return spells, nil
}

func (b *Mysql) SpellSearchByName(ctx context.Context, name string) ([]*models.Spell, error) {
	spells := []*models.Spell{}

	rows, err := b.query.SpellSearchByName(ctx, sql.NullString{String: name, Valid: true})
	if err != nil {
		return nil, fmt.Errorf("spell search by name: %w", err)
	}
	for _, row := range rows {
		spell := &models.Spell{}
		spell.DecodeSpellsNew(row)
		spells = append(spells, spell)
	}
	return spells, nil
}

func (b *Mysql) SpellByID(ctx context.Context, spellID int64) (*models.Spell, error) {
	spell := &models.Spell{}

	row, err := b.query.SpellByID(ctx, int32(spellID))
	if err != nil {
		return nil, fmt.Errorf("spell by spell id: %w", err)
	}
	spell.DecodeSpellsNew(row)
	return spell, nil
}
