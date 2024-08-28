package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type Spell struct {
	ID                int32
	Name              string
	Player1           string
	YouCast           string
	OtherCasts        string
	CastOnYou         string
	CastOnOther       string
	SpellFades        string
	Attribs           []int32 // effectid
	Bases             []int32 // effect_base_value
	Calcs             []int32 // formula
	Limits            []int32 // effect_limit_value
	Maxes             []int32 // max
	Classes           []int32 // classes
	Components        []int32
	ComponentCounts   []int32
	NoExpendReagents  []int32
	Deities           []int32
	Basediff          int32
	Zonetype          int32
	Environmenttype   int32
	Timeofday         int32
	Castinganim       int32
	Targetanim        int32
	Traveltype        int32
	Spellaffectindex  int32
	DisallowSit       int32
	Lighttype         int32
	Goodeffect        int32
	Activated         int32
	Resisttype        int32
	Icon              int32
	Memicon           int32
	Range             int32
	Aoerange          int32
	DurationCap       int32
	DurationCalc      int32
	Aeduration        int32
	MaxTargets        int32
	TargetType        int32
	Skill             int32
	RecoveryTime      int32
	RecastTime        int32
	Pushback          int32
	Pushup            int32
	CastTime          int32
	TeleportZone      string
	Mana              int32
	SpellGroup        int32
	SpellIcon         int32
	NewIcon           int32
	Spellanim         int32
	Uninterruptable   int32
	Resistdiff        int32
	DotStackingExempt int32
	Field142          int32
	Field143          int32
	Deleteable        int32
	Recourselink      int32
	NoPartialResist   int32
	Field152          int32
	Field153          int32
	ShortBuffBox      int32
	Descnum           int32
	Typedescnum       int32
	Effectdescnum     int32
	Effectdescnum2    int32
	NpcNoLos          int32
	Field160          int32
	Reflectable       int32
	Bonushate         int32
	Field163          int32
}

func (t *Spell) Serialize() string {
	return serialize(t)
}

func (t *Spell) Deserialize(data string) error {
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
