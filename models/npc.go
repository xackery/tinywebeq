package models

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/xackery/tinywebeq/library"
)

type Npc struct {
	CacheKey             string `db:"key"`
	CacheExpiration      int64  `db:"expiration"`
	ID                   int32
	Name                 string
	Lastname             sql.NullString
	Level                uint8
	Race                 uint16
	Class                uint8
	Bodytype             int32
	Hp                   int64
	Mana                 int64
	Gender               uint8
	Texture              uint8
	Helmtexture          uint8
	Herosforgemodel      int32
	Size                 float64
	HpRegenRate          int64
	HpRegenPerSecond     int64
	ManaRegenRate        int64
	LoottableID          uint32
	MerchantID           uint32
	AltCurrencyID        uint32
	NpcSpellsID          uint32
	NpcSpellsEffectsID   uint32
	NpcFactionID         int32
	AdventureTemplateID  uint32
	TrapTemplate         sql.NullInt32
	Mindmg               uint32
	Maxdmg               uint32
	AttackCount          int16
	Npcspecialattks      string
	SpecialAbilities     sql.NullString
	Aggroradius          uint32
	Assistradius         uint32
	Face                 uint32
	LuclinHairstyle      uint32
	LuclinHaircolor      uint32
	LuclinEyecolor       uint32
	LuclinEyecolor2      uint32
	LuclinBeardcolor     uint32
	LuclinBeard          uint32
	DrakkinHeritage      int32
	DrakkinTattoo        int32
	DrakkinDetails       int32
	ArmortintID          uint32
	ArmortintRed         uint8
	ArmortintGreen       uint8
	ArmortintBlue        uint8
	DMeleeTexture1       uint32
	DMeleeTexture2       uint32
	AmmoIdfile           string
	PrimMeleeType        uint8
	SecMeleeType         uint8
	RangedType           uint8
	Runspeed             float64
	Mr                   int16
	Cr                   int16
	Dr                   int16
	Fr                   int16
	Pr                   int16
	Corrup               int16
	Phr                  uint16
	SeeInvis             int16
	SeeInvisUndead       int16
	Qglobal              uint32
	Ac                   int16
	NpcAggro             int8
	SpawnLimit           int8
	AttackSpeed          float64
	AttackDelay          uint8
	Findable             int8
	Str                  uint32
	Sta                  uint32
	Dex                  uint32
	Agi                  uint32
	Int                  uint32
	Wis                  uint32
	Cha                  uint32
	SeeHide              int8
	SeeImprovedHide      int8
	Trackable            int8
	Isbot                int8
	Exclude              int8
	Atk                  int32
	Accuracy             int32
	Avoidance            uint32
	SlowMitigation       int16
	Version              uint16
	Maxlevel             int8
	Scalerate            int32
	PrivateCorpse        uint8
	UniqueSpawnByName    uint8
	Underwater           uint8
	Isquest              int8
	Emoteid              uint32
	Spellscale           float64
	Healscale            float64
	NoTargetHotkey       bool
	RaidTarget           bool
	Armtexture           int8
	Bracertexture        int8
	Handtexture          int8
	Legtexture           int8
	Feettexture          int8
	Light                int8
	Walkspeed            int8
	Peqid                int32
	Unique               int8
	Fixed                int8
	IgnoreDespawn        int8
	ShowName             int8
	Untargetable         int8
	CharmAc              sql.NullInt16
	CharmMinDmg          sql.NullInt32
	CharmMaxDmg          sql.NullInt32
	CharmAttackDelay     sql.NullInt16
	CharmAccuracyRating  sql.NullInt32
	CharmAvoidanceRating sql.NullInt32
	CharmAtk             sql.NullInt32
	SkipGlobalLoot       sql.NullInt16
	RareSpawn            int16
	StuckBehavior        int8
	Model                int16
	Flymode              int8
	AlwaysAggro          bool
	ExpMod               int32
	HeroicStrikethrough  int32
	FactionAmount        int32
	KeepsSoldItems       bool
}

func (t *Npc) Identifier() string {
	return "npc"
}

func (t *Npc) Key() string {
	return t.CacheKey
}

func (t *Npc) SetKey(key string) {
	t.CacheKey = key
}

func (t *Npc) SetExpiration(expiration int64) {
	t.CacheExpiration = expiration
}

func (t *Npc) Expiration() int64 {
	return t.CacheExpiration
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

func (t *Npc) CleanName() string {
	out := t.Name
	out = strings.ReplaceAll(out, "_", " ")
	out = strings.ReplaceAll(out, "-", "`")
	out = strings.ReplaceAll(out, "#", "")
	out = strings.ReplaceAll(out, "!", "")
	out = strings.ReplaceAll(out, "~", "")
	return out
}

func (t *Npc) RaceStr() string {
	return library.RaceStr(int32(t.Race))
}

func (t *Npc) ClassStr() string {
	return library.Class(t.Class).String()
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
			out += "Unstunnable, "
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

func (t *Npc) ZoneID() int32 {
	return t.ID / 1000
}
