package model

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type NpcSpell struct {
	Entries         []*NpcSpellEntry
	CacheKey        string `db:"key"`
	CacheExpiration int64
}

type NpcSpellEntry struct {
	ID            uint32
	Name          sql.NullString
	ProcChance    int8
	AttackProc    int16
	RangeProc     int16
	RprocChance   int16
	DefensiveProc int16
	DprocChance   int16
	Npcspellid    int `db:"npc_spells_id"`
	Spellid       int `db:"spellid"`
}

func (t *NpcSpell) Identifier() string {
	return "npc_spell"
}

func (t *NpcSpell) Key() string {
	return t.CacheKey
}

func (t *NpcSpell) SetKey(key string) {
	t.CacheKey = key
}

func (t *NpcSpell) SetExpiration(expiration int64) {
	t.CacheExpiration = expiration
}

func (t *NpcSpell) Expiration() int64 {
	return t.CacheExpiration
}

func (t *NpcSpell) Serialize() string {
	return serialize(t)
}

func (t *NpcSpell) Deserialize(data string) error {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return fmt.Errorf("base64 decode: %w", err)
	}
	buf := bytes.NewBuffer(decoded)
	d := gob.NewDecoder(buf)

	err = d.Decode(&t)
	if err != nil {
		return fmt.Errorf("gob decode: %w", err)
	}
	return nil
}
