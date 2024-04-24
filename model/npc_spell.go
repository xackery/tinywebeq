package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/xackery/tinywebeq/library"
)

type NpcSpell struct {
	Entries    []*NpcSpellEntry
	key        string
	expiration int64
}

type NpcSpellEntry struct {
	ID            int    `db:"id"`
	Name          string `db:"name"`
	Procchance    int    `db:"proc_chance"`
	Attackproc    int    `db:"attack_proc"`
	Rangeproc     int    `db:"range_proc"`
	Rprocchance   int    `db:"rproc_chance"`
	Defensiveproc int    `db:"defensive_proc"`
	Dprocchance   int    `db:"dproc_chance"`
	Npcspellid    int    `db:"npc_spells_id"`
	Spellid       int    `db:"spellid"`
}

func (t *NpcSpell) Identifier() string {
	return "npc_spell"
}

func (t *NpcSpell) Key() string {
	return t.key
}

func (t *NpcSpell) SetKey(key string) {
	t.key = key
}

func (t *NpcSpell) SetExpiration(expiration int64) {
	t.expiration = expiration
}

func (t *NpcSpell) Expiration() int64 {
	return t.expiration
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

func (t *NpcSpellEntry) Spell() *library.Spell {
	return library.SpellByID(t.Spellid)
}

func (t *NpcSpellEntry) SpellInfo(level int) []string {
	_, lines := library.SpellInfo(t.Spellid, level)
	newLines := []string{}
	isSlot := false
	for _, line := range lines {
		if strings.HasPrefix(line, "Spell Info") {
			continue
		}
		if strings.HasPrefix(line, "ID: ") {
			continue
		}
		if strings.HasPrefix(line, "Recovery Time: ") {
			continue
		}
		if strings.HasPrefix(line, "Mana: ") {
			continue
		}
		if strings.HasPrefix(line, "Slot") {
			isSlot = true
		}
		if isSlot && !strings.HasPrefix(line, "Slot") {
			break
		}

		newLines = append(newLines, line)
	}
	return newLines
}
