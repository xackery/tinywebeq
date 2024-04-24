package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"

	"github.com/xackery/tinywebeq/library"
)

type Npc struct {
	key             string
	expiration      int64
	ID              int    `db:"id"`
	Name            string `db:"name"`
	Attackspeed     int    `db:"attack_speed"`
	Class           int    `db:"class"`
	Hp              int    `db:"hp"`
	Lastname        string `db:"lastname"`
	Level           int    `db:"level"`
	Maxdmg          int    `db:"maxdmg"`
	Mindmg          int    `db:"mindmg"`
	Npcspellsid     int    `db:"npc_spells_id"`
	Npcspecialattks string `db:"npcspecialattks"`
	Race            int    `db:"race"`
	Trackable       int    `db:"trackable"`
	Loottableid     int    `db:"loottable_id"`
	Merchantid      int    `db:"merchant_id"`
	Npcfactionid    int    `db:"npc_faction_id"`
	Rarespawn       int    `db:"rare_spawn"`
}

func (t *Npc) Identifier() string {
	return "npc"
}

func (t *Npc) Key() string {
	return t.key
}

func (t *Npc) SetKey(key string) {
	t.key = key
}

func (t *Npc) SetExpiration(expiration int64) {
	t.expiration = expiration
}

func (t *Npc) Expiration() int64 {
	return t.expiration
}

func (t *Npc) Serialize() string {
	return serialize(t)
}

func (t *Npc) Deserialize(data string) error {
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

func (t *Npc) RaceStr() string {
	return library.RaceStr(t.Race)
}

func (t *Npc) ClassStr() string {
	return library.ClassStr(t.Class)
}

func (t *Npc) NpcSpecialAttacksStr() string {

	out := ""
	for _, v := range t.Npcspecialattks {
		switch v {
		case 'A':
			out += "Immune to melee, "
		case 'B':
			out += "Immune to magic, "
		case 'C':
			out += "Uncharmable, "
		case 'D':
			out += "Unfearable, "
		case 'E':
			out += "Enrage, "
		case 'F':
			out += "Flurry, "
		case 'f':
			out += "Immune to fleeing, "
		case 'I':
			out += "Unsnarable, "
		case 'M':
			out += "Unmezzable, "
		case 'N':
			out += "Unstunable, "
		case 'O':
			out += "Immune to melee except bane, "
		case 'Q':
			out += "Quadruple Attack, "
		case 'R':
			out += "Rampage, "
		case 'S':
			out += "Summon, "
		case 'T':
			out += "Triple Attack, "
		case 'U':
			out += "Unslowable, "
		case 'W':
			out += "Immune to melee except magical, "
		default:
			out += fmt.Sprintf("Unknown %s, ", string(v))
		}
	}
	if len(out) > 0 {
		out = out[:len(out)-2]
	}
	return out
}
