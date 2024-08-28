package model

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/xackery/tinywebeq/library"
)

type Player struct {
	CacheKey              string
	CacheExpiration       int64
	ID                    uint32
	AccountID             int32
	Name                  string
	LastName              string
	Title                 string
	Suffix                string
	ZoneID                uint32
	ZoneInstance          uint32
	Y                     float64
	X                     float64
	Z                     float64
	Heading               float64
	Gender                uint8
	Race                  uint16
	Class                 uint8
	Level                 uint32
	Deity                 uint32
	Birthday              uint32
	LastLogin             uint32
	TimePlayed            uint32
	Level2                uint8
	Anon                  uint8
	Gm                    uint8
	Face                  uint32
	HairColor             uint8
	HairStyle             uint8
	Beard                 uint8
	BeardColor            uint8
	EyeColor1             uint8
	EyeColor2             uint8
	DrakkinHeritage       uint32
	DrakkinTattoo         uint32
	DrakkinDetails        uint32
	AbilityTimeSeconds    uint8
	AbilityNumber         uint8
	AbilityTimeMinutes    uint8
	AbilityTimeHours      uint8
	Exp                   uint32
	ExpEnabled            bool
	AaPointsSpent         uint32
	AaExp                 uint32
	AaPoints              uint32
	GroupLeadershipExp    uint32
	RaidLeadershipExp     uint32
	GroupLeadershipPoints uint32
	RaidLeadershipPoints  uint32
	Points                uint32
	CurHp                 uint32
	Mana                  uint32
	Endurance             uint32
	Intoxication          uint32
	Str                   uint32
	Sta                   uint32
	Cha                   uint32
	Dex                   uint32
	Int                   uint32
	Agi                   uint32
	Wis                   uint32
	ZoneChangeCount       uint32
	Toxicity              uint32
	HungerLevel           uint32
	ThirstLevel           uint32
	AbilityUp             uint32
	LdonPointsGuk         uint32
	LdonPointsMir         uint32
	LdonPointsMmc         uint32
	LdonPointsRuj         uint32
	LdonPointsTak         uint32
	LdonPointsAvailable   uint32
	TributeTimeRemaining  uint32
	CareerTributePoints   uint32
	TributePoints         uint32
	TributeActive         uint32
	PvpStatus             uint8
	PvpKills              uint32
	PvpDeaths             uint32
	PvpCurrentPoints      uint32
	PvpCareerPoints       uint32
	PvpBestKillStreak     uint32
	PvpWorstDeathStreak   uint32
	PvpCurrentKillStreak  uint32
	Pvp2                  uint32
	PvpType               uint32
	ShowHelm              uint32
	GroupAutoConsent      uint8
	RaidAutoConsent       uint8
	GuildAutoConsent      uint8
	LeadershipExpOn       uint8
	Resttimer             uint32
	AirRemaining          uint32
	AutosplitEnabled      uint32
	Lfp                   bool
	Lfg                   bool
	Mailkey               string
	Xtargets              uint8
	Firstlogon            int8
	EAaEffects            uint32
	EPercentToAa          uint32
	EExpendedAaSpent      uint32
	AaPointsSpentOld      uint32
	AaPointsOld           uint32
	ELastInvsnapshot      uint32
	DeletedAt             sql.NullTime
}

func (t *Player) Identifier() string {
	return "player"
}

func (t *Player) Key() string {
	return t.CacheKey
}

func (t *Player) SetKey(key string) {
	t.CacheKey = key
}

func (t *Player) SetExpiration(expiration int64) {
	t.CacheExpiration = expiration
}

func (t *Player) Expiration() int64 {
	return t.CacheExpiration
}

func (t *Player) ClassStr() string {
	return library.ClassStr(t.Class)
}

func (t *Player) RaceStr() string {
	out := ""

	out = strings.TrimSuffix(out, " ")
	return out
}

func (t *Player) IconUrl() string {
	return "https://www.eqitems.com/item_images/"
}

func (t *Player) Serialize() string {
	return serialize(t)
}

func (t *Player) Deserialize(data string) error {
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
