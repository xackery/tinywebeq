package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type Zone struct {
	Icon         int32
	ID           int32
	Zoneidnumber int32
	Version      uint8
	ShortName    string
	LongName     string
	// MinStatus               uint8
	// MapFileName             sql.NullString
	// Note                    sql.NullString
	MinExpansion int8
	MaxExpansion int8
	// ContentFlags            sql.NullString
	// ContentFlagsDisabled    sql.NullString
	Expansion int8
	// FileName                sql.NullString
	// SafeX                   float64
	// SafeY                   float64
	// SafeZ                   float64
	// SafeHeading             float64
	// GraveyardID             float64
	MinLevel uint8
	MaxLevel uint8
	// Timezone                int32
	// Maxclients              int32
	// Ruleset                 uint32
	// Underworld              float64
	// Minclip                 float64
	// Maxclip                 float64
	// FogMinclip              float64
	// FogMaxclip              float64
	// FogBlue                 uint8
	// FogRed                  uint8
	// FogGreen                uint8
	// Sky                     uint8
	// Ztype                   uint8
	// ZoneExpMultiplier       string
	// Walkspeed               float64
	// TimeType                uint8
	// FogRed1                 uint8
	// FogGreen1               uint8
	// FogBlue1                uint8
	// FogMinclip1             float64
	// FogMaxclip1             float64
	// FogRed2                 uint8
	// FogGreen2               uint8
	// FogBlue2                uint8
	// FogMinclip2             float64
	// FogMaxclip2             float64
	// FogRed3                 uint8
	// FogGreen3               uint8
	// FogBlue3                uint8
	// FogMinclip3             float64
	// FogMaxclip3             float64
	// FogRed4                 uint8
	// FogGreen4               uint8
	// FogBlue4                uint8
	// FogMinclip4             float64
	// FogMaxclip4             float64
	// FogDensity              float64
	// FlagNeeded              string
	// Canbind                 int8
	// Cancombat               int8
	// Canlevitate             int8
	// Castoutdoor             int8
	// Hotzone                 uint8
	// Insttype                bool
	// Shutdowndelay           uint64
	// Peqzone                 int8
	// BypassExpansionCheck    int8
	// Suspendbuffs            bool
	// RainChance1             int32
	// RainChance2             int32
	// RainChance3             int32
	// RainChance4             int32
	// RainDuration1           int32
	// RainDuration2           int32
	// RainDuration3           int32
	// RainDuration4           int32
	// SnowChance1             int32
	// SnowChance2             int32
	// SnowChance3             int32
	// SnowChance4             int32
	// SnowDuration1           int32
	// SnowDuration2           int32
	// SnowDuration3           int32
	// SnowDuration4           int32
	// Gravity                 float64
	// Type                    int32
	// Skylock                 int8
	// FastRegenHp             int32
	// FastRegenMana           int32
	// FastRegenEndurance      int32
	// NpcMaxAggroDist         int32
	// MaxMovementUpdateRange  uint32
	// UnderworldTeleportIndex int32
	// LavaDamage              sql.NullInt32
	// MinLavaDamage           int32
	// IdleWhenEmpty           bool
	// SecondsBeforeIdle       uint32
}

func (t *Zone) Serialize() string {
	return serialize(t)
}

func (t *Zone) Deserialize(data string) error {
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
