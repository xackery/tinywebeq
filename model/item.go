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

type Item struct {
	CacheKey            string `json:"-"`
	CacheExpiration     int64  `json:"-"`
	ID                  int32  `json:"id"`
	Minstatus           int16  `json:"-"`
	Name                string `json:"name"`
	Aagi                int32  `json:"agi,omitempty"`
	Ac                  int32  `json:"ac,omitempty"`
	Accuracy            int32  `json:"accuracy,omitempty"`
	Acha                int32  `json:"cha,omitempty"`
	Adex                int32  `json:"dex,omitempty"`
	Aint                int32  `json:"int,omitempty"`
	Artifactflag        uint8  `json:"artifact_flag,omitempty"`
	Asta                int32  `json:"sta,omitempty"`
	Astr                int32  `json:"str,omitempty"`
	Attack              int32  `json:"attack,omitempty"`
	Augrestrict         int32  `json:"augment_restrictions,omitempty"`
	Augslot1type        int8   `json:"augment_slot_1_type,omitempty"`
	Augslot1visible     int8   `json:"augment_slot_1_visible,omitempty"`
	Augslot2type        int8   `json:"augment_slot_2_type,omitempty"`
	Augslot2visible     int8   `json:"augment_slot_2_visible,omitempty"`
	Augslot3type        int8   `json:"augment_slot_3_type,omitempty"`
	Augslot3visible     int8   `json:"augment_slot_3_visible,omitempty"`
	Augslot4type        int8   `json:"augment_slot_4_type,omitempty"`
	Augslot4visible     int8   `json:"augment_slot_4_visible,omitempty"`
	Augslot5type        int8   `json:"augment_slot_5_type,omitempty"`
	Augslot5visible     int8   `json:"augment_slot_5_visible,omitempty"`
	Augslot6type        int8   `json:"augment_slot_6_type,omitempty"`
	Augslot6visible     int8   `json:"augment_slot_6_visible,omitempty"`
	Augtype             int32  `json:"augment_type,omitempty"`
	Avoidance           int32  `json:"avoidance,omitempty"`
	Awis                int32  `json:"wis,omitempty"`
	Bagsize             int32  `json:"bag_size,omitempty"`
	Bagslots            int32  `json:"bag_slots,omitempty"`
	Bagtype             int32  `json:"bag_type,omitempty"`
	Bagwr               int32  `json:"bag_weight_reduction,omitempty"`
	Banedmgamt          int32  `json:"bane_damage_amount,omitempty"`
	Banedmgraceamt      int32  `json:"bane_damage_race_amount,omitempty"`
	Banedmgbody         int32  `json:"bane_damage_body_amount,omitempty"`
	Banedmgrace         int32  `json:"bane_damage_race,omitempty"`
	Bardtype            int32  `json:"bard_type,omitempty"`
	Bardvalue           int32  `json:"bard_value,omitempty"`
	Book                int32  `json:"book,omitempty"`
	Casttime            int32  `json:"cast_time,omitempty"`
	Charmfile           string `json:"charm_file,omitempty"`
	Charmfileid         string `json:"charm_file_id,omitempty"`
	Classes             int32  `json:"classes,omitempty"`
	Color               uint32 `json:"color,omitempty"`
	Combateffects       string `json:"combat_effects,omitempty"`
	Extradmgskill       int32  `json:"extra_damage_skill,omitempty"`
	Extradmgamt         int32  `json:"extra_damage_race,omitempty"`
	Price               int32  `json:"price,omitempty"`
	Cr                  int32  `json:"cold_resist,omitempty"`
	Damage              int32  `json:"damage,omitempty"`
	Damageshield        int32  `json:"damage_shield,omitempty"`
	Deity               int32  `json:"deity,omitempty"`
	Delay               int32  `json:"delay,omitempty"`
	Augdistiller        uint32 `json:"augment_distiller,omitempty"`
	Dotshielding        int32  `json:"dot_shielding,omitempty"`
	Dr                  int32  `json:"disease_resist,omitempty"`
	Clicktype           int32  `json:"click_type,omitempty"`
	Clicklevel2         int32  `json:"click_level_2,omitempty"`
	Elemdmgtype         int32  `json:"elemental_damage_type,omitempty"`
	Elemdmgamt          int32  `json:"elemental_damage_race,omitempty"`
	Endur               int32  `json:"endurance,omitempty"`
	Factionamt1         int32  `json:"faction_amount_1,omitempty"`
	Factionamt2         int32  `json:"faction_amount_2,omitempty"`
	Factionamt3         int32  `json:"faction_amount_3,omitempty"`
	Factionamt4         int32  `json:"faction_amount_4,omitempty"`
	Factionmod1         int32  `json:"faction_mod_1,omitempty"`
	Factionmod2         int32  `json:"faction_mod_2,omitempty"`
	Factionmod3         int32  `json:"faction_mod_3,omitempty"`
	Factionmod4         int32  `json:"faction_mod_4,omitempty"`
	Filename            string `json:"filename,omitempty"`
	Focuseffect         int32  `json:"focus_effect,omitempty"`
	Fr                  int32  `json:"fire_resist,omitempty"`
	Fvnodrop            int32  `json:"fv_no_drop,omitempty"`
	Haste               int32  `json:"haste,omitempty"`
	Clicklevel          int32  `json:"click_level,omitempty"`
	Hp                  int32  `json:"hp,omitempty"`
	Regen               int32  `json:"regen,omitempty"`
	Icon                int32  `json:"icon,omitempty"`
	Idfile              string `json:"id_file,omitempty"`
	Itemclass           int32
	Itemtype            int32
	Ldonprice           int32
	Ldontheme           int32
	Ldonsold            int32
	Light               int32
	Lore                string
	Loregroup           int32
	Magic               int32
	Mana                int32
	Manaregen           int32
	Enduranceregen      int32
	Material            int32
	Herosforgemodel     int32
	Maxcharges          int32
	Mr                  int32
	Nodrop              int32
	Norent              int32
	Pendingloreflag     uint8
	Pr                  int32
	Procrate            int32
	Races               int32
	Range               int32
	Reclevel            int32
	Recskill            int32
	Reqlevel            int32
	Sellrate            float64
	Shielding           int32
	Size                int32
	Skillmodtype        int32
	Skillmodvalue       int32
	Slots               int32
	Clickeffect         int32
	Spellshield         int32
	Strikethrough       int32
	Stunresist          int32
	Summonedflag        uint8
	Tradeskills         int32
	Favor               int32
	Weight              int32
	Unk012              int32 `json:"-"`
	Unk013              int32 `json:"-"`
	Benefitflag         int32
	Unk054              int32 `json:"-"`
	Unk059              int32 `json:"-"`
	Booktype            int32
	Recastdelay         int32
	Recasttype          int32
	Guildfavor          int32
	Unk123              int32 `json:"-"`
	Unk124              int32 `json:"-"`
	Attuneable          int32
	Nopet               int32
	Updated             sql.NullTime
	Comment             string
	Unk127              int32 `json:"-"`
	Pointtype           int32
	Potionbelt          int32
	Potionbeltslots     int32
	Stacksize           int32
	Notransfer          int32
	Stackable           int32
	Unk134              string `json:"-"`
	Unk137              int32  `json:"-"`
	Proceffect          int32
	Proctype            int32
	Proclevel2          int32
	Proclevel           int32
	Unk142              int32 `json:"-"`
	Worneffect          int32
	Worntype            int32
	Wornlevel2          int32
	Wornlevel           int32
	Unk147              int32 `json:"-"`
	Focustype           int32
	Focuslevel2         int32
	Focuslevel          int32
	Unk152              int32 `json:"-"`
	Scrolleffect        int32
	Scrolltype          int32
	Scrolllevel2        int32
	Scrolllevel         int32
	Unk157              int32 `json:"-"`
	Serialized          sql.NullTime
	Verified            sql.NullTime
	Serialization       sql.NullString
	Source              string
	Unk033              int32 `json:"-"`
	Lorefile            string
	Unk014              int32 `json:"-"`
	Svcorruption        int32
	Skillmodmax         int32
	Unk060              int32 `json:"-"`
	Augslot1unk2        int32 `json:"-"`
	Augslot2unk2        int32 `json:"-"`
	Augslot3unk2        int32 `json:"-"`
	Augslot4unk2        int32 `json:"-"`
	Augslot5unk2        int32 `json:"-"`
	Augslot6unk2        int32 `json:"-"`
	Unk120              int32 `json:"-"`
	Unk121              int32 `json:"-"`
	Questitemflag       int32
	Unk132              sql.NullString `json:"-"`
	Clickunk5           int32          `json:"-"`
	Clickunk6           string         `json:"-"`
	Clickunk7           int32          `json:"-"`
	Procunk1            int32          `json:"-"`
	Procunk2            int32          `json:"-"`
	Procunk3            int32          `json:"-"`
	Procunk4            int32          `json:"-"`
	Procunk6            string         `json:"-"`
	Procunk7            int32          `json:"-"`
	Wornunk1            int32          `json:"-"`
	Wornunk2            int32          `json:"-"`
	Wornunk3            int32          `json:"-"`
	Wornunk4            int32          `json:"-"`
	Wornunk5            int32          `json:"-"`
	Wornunk6            string         `json:"-"`
	Wornunk7            int32          `json:"-"`
	Focusunk1           int32          `json:"-"`
	Focusunk2           int32          `json:"-"`
	Focusunk3           int32          `json:"-"`
	Focusunk4           int32          `json:"-"`
	Focusunk5           int32          `json:"-"`
	Focusunk6           string         `json:"-"`
	Focusunk7           int32          `json:"-"`
	Scrollunk1          uint32         `json:"-"`
	Scrollunk2          int32          `json:"-"`
	Scrollunk3          int32          `json:"-"`
	Scrollunk4          int32          `json:"-"`
	Scrollunk5          int32          `json:"-"`
	Scrollunk6          string         `json:"-"`
	Scrollunk7          int32          `json:"-"`
	Unk193              int32          `json:"-"`
	Purity              int32
	Evoitem             int32
	Evoid               int32
	Evolvinglevel       int32
	Evomax              int32
	Clickname           string
	Procname            string
	Wornname            string
	Focusname           string
	Scrollname          string
	Dsmitigation        int16
	HeroicStr           int16
	HeroicInt           int16
	HeroicWis           int16
	HeroicAgi           int16
	HeroicDex           int16
	HeroicSta           int16
	HeroicCha           int16
	HeroicPr            int16
	HeroicDr            int16
	HeroicFr            int16
	HeroicCr            int16
	HeroicMr            int16
	HeroicSvcorrup      int16
	Healamt             int16
	Spelldmg            int16
	Clairvoyance        int16
	Backstabdmg         int16
	Created             string
	Elitematerial       int16
	Ldonsellbackrate    int16
	Scriptfileid        int32
	Expendablearrow     int16
	Powersourcecapacity int32
	Bardeffect          int32
	Bardeffecttype      int16
	Bardlevel2          int16
	Bardlevel           int16
	Bardunk1            int16 `json:"-"`
	Bardunk2            int16 `json:"-"`
	Bardunk3            int16 `json:"-"`
	Bardunk4            int16 `json:"-"`
	Bardunk5            int16 `json:"-"`
	Bardname            string
	Bardunk7            int16 `json:"-"`
	Unk214              int16 `json:"-"`
	Subtype             int32
	Unk220              int32 `json:"-"`
	Unk221              int32 `json:"-"`
	Heirloom            int32
	Unk223              int32 `json:"-"`
	Unk224              int32 `json:"-"`
	Unk225              int32 `json:"-"`
	Unk226              int32 `json:"-"`
	Unk227              int32 `json:"-"`
	Unk228              int32 `json:"-"`
	Unk229              int32 `json:"-"`
	Unk230              int32 `json:"-"`
	Unk231              int32 `json:"-"`
	Unk232              int32 `json:"-"`
	Unk233              int32 `json:"-"`
	Unk234              int32 `json:"-"`
	Placeable           int32
	Unk236              int32 `json:"-"`
	Unk237              int32 `json:"-"`
	Unk238              int32 `json:"-"`
	Unk239              int32 `json:"-"`
	Unk240              int32 `json:"-"`
	Unk241              int32 `json:"-"`
	Epicitem            int32
	ItemID              uint32
	CharName            string
	DiscoveredDate      uint32
	AccountStatus       int32
}

func (t *Item) Identifier() string {
	return "item"
}

func (t *Item) Key() string {
	return t.CacheKey
}

func (t *Item) SetKey(key string) {
	t.CacheKey = key
}

func (t *Item) SetExpiration(expiration int64) {
	t.CacheExpiration = expiration
}

func (t *Item) Expiration() int64 {
	return t.CacheExpiration
}

func (t *Item) Serialize() string {
	return serialize(t)
}

func (t *Item) Deserialize(data string) error {
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

func (t *Item) ClassesStr() string {
	return library.ClassesFromMask(t.Classes)
}

func (t *Item) RaceStr() string {
	out := ""
	if t.Races == 65535 {
		return "ALL"
	}
	if t.Races&1 != 0 {
		out += "HUM "
	}
	if t.Races&2 != 0 {
		out += "BAR "
	}
	if t.Races&4 != 0 {
		out += "ERU "
	}
	if t.Races&8 != 0 {
		out += "WLF "
	}
	if t.Races&16 != 0 {
		out += "HEF "
	}
	if t.Races&32 != 0 {
		out += "DKE "
	}
	if t.Races&64 != 0 {
		out += "HLF "
	}
	if t.Races&128 != 0 {
		out += "DWF "
	}
	if t.Races&256 != 0 {
		out += "TRL "
	}
	if t.Races&512 != 0 {
		out += "OGR "
	}
	if t.Races&1024 != 0 {
		out += "HFL "
	}
	if t.Races&2048 != 0 {
		out += "GNM "
	}
	if t.Races&4096 != 0 {
		out += "IKS "
	}
	if t.Races&8192 != 0 {
		out += "VAH "
	}
	if t.Races&16384 != 0 {
		out += "FRG "
	}
	if t.Races&32768 != 0 {
		out += "DRK "
	}
	out = strings.TrimSuffix(out, " ")
	return out
}

func (t *Item) DeityStr() string {
	out := ""
	if t.Deity == 65535 {
		return "ALL"
	}
	if t.Deity&32768 != 0 {
		out += "Tunare, "
	}
	if t.Deity&16384 != 0 {
		out += "The Tribunal, "
	}
	if t.Deity&8192 != 0 {
		out += "Solusek Ro, "
	}
	if t.Deity&4096 != 0 {
		out += "Rodcet Nife, "
	}
	if t.Deity&2048 != 0 {
		out += "Rallos Zek, "
	}
	if t.Deity&1024 != 0 {
		out += "Quellious, "
	}
	if t.Deity&512 != 0 {
		out += "Prexus, "
	}
	if t.Deity&256 != 0 {
		out += "Mithaniel Marr, "
	}
	if t.Deity&128 != 0 {
		out += "Karana, "
	}
	if t.Deity&64 != 0 {
		out += "Innoruuk, "
	}
	if t.Deity&32 != 0 {
		out += "Bristlebane, "
	}
	if t.Deity&16 != 0 {
		out += "Erollisi Marr, "
	}
	if t.Deity&8 != 0 {
		out += "Cazic Thule, "
	}
	if t.Deity&4 != 0 {
		out += "Brell Serilis, "
	}
	if t.Deity&2 != 0 {
		out += "Bertoxxulous, "
	}
	if t.Deity&1 != 0 {
		out += "Agnostic, "
	}
	out = strings.TrimSuffix(out, ", ")
	return out

}

func (t *Item) SlotStr() string {
	out := ""
	if t.Slots&4194304 != 0 {
		out += "Power Source, "
	}
	if t.Slots&2097152 != 0 {
		out += "Ammo, "
	}
	if t.Slots&1048576 != 0 {
		out += "Waist, "
	}
	if t.Slots&524288 != 0 {
		out += "Feet, "
	}
	if t.Slots&262144 != 0 {
		out += "Legs, "
	}
	if t.Slots&131072 != 0 {
		out += "Chest, "
	}
	if t.Slots&98304 != 0 {
		out += "Fingers, "
	}
	if t.Slots&65536 != 0 {
		out += "Finger, "
	}
	if t.Slots&32768 != 0 {
		out += "Finger, "
	}
	if t.Slots&8192 != 0 {
		out += "Primary, "
	}
	if t.Slots&16384 != 0 {
		out += "Secondary, "
	}
	if t.Slots&4096 != 0 {
		out += "Hands, "
	}
	if t.Slots&2048 != 0 {
		out += "Range, "
	}
	if t.Slots&1536 != 0 {
		out += "Wrists, "
	}
	if t.Slots&1024 != 0 {
		out += "Wrist, "
	}
	if t.Slots&512 != 0 {
		out += "Wrist, "
	}
	if t.Slots&256 != 0 {
		out += "Back, "
	}
	if t.Slots&128 != 0 {
		out += "Arms, "
	}
	if t.Slots&64 != 0 {
		out += "Shoulders, "
	}
	if t.Slots&32 != 0 {
		out += "Neck, "
	}
	if t.Slots&18 != 0 {
		out += "Ears, "
	}
	if t.Slots&16 != 0 {
		out += "Ear, "
	}
	if t.Slots&8 != 0 {
		out += "Face, "
	}
	if t.Slots&4 != 0 {
		out += "Head, "
	}
	if t.Slots&2 != 0 {
		out += "Ear, "
	}
	if t.Slots&1 != 0 {
		out += "Charm, "
	}
	return strings.TrimSuffix(out, ", ")
}

func (t *Item) SizeStr() string {
	switch t.Size {
	case 0:
		return "TINY"
	case 1:
		return "SMALL"
	case 2:
		return "MEDIUM"
	case 3:
		return "LARGE"
	case 4:
		return "GIANT"
	default:
		return fmt.Sprintf("Unknown Size %d", t.Size)
	}
}

func (t *Item) BagsizeStr() string {
	switch t.Bagsize {
	case 0:
		return "TINY"
	case 1:
		return "SMALL"
	case 2:
		return "MEDIUM"
	case 3:
		return "LARGE"
	case 4:
		return "GIANT"
	default:
		return fmt.Sprintf("Unknown Size %d", t.Bagsize)
	}
}

func (t *Item) IconUrl() string {
	return "https://www.eqitems.com/item_images/"
}

func (t *Item) TagStr() string {
	out := ""
	if t.Itemtype == 54 {
		out += "AUGMENT, "
	}
	if t.Magic == 1 {
		out += "MAGIC, "
	}
	if t.Loregroup > 0 {
		out += "LORE, "
	}
	if t.Nodrop == 0 {
		out += "NO TRADE, "
	}
	if t.Norent == 0 {
		out += "NO RENT, "
	}
	return strings.TrimSuffix(out, ", ")
}

func (t *Item) TypeStr() string {
	switch t.Itemtype {
	case 0, 2, 3, 42, 1, 4, 35:
		return "Skill"
	default:
		return "Item Type"
	}
}

func (t *Item) BagTypeStr() string {
	switch t.Bagtype {
	case 9:
		return "Alchemy"
	case 10:
		return "Tinkering"
	case 12:
		return "Poison Making"
	case 13:
		return "Special Quests"
	case 14:
		return "Baking"
	case 15:
		return "Baking"
	case 16:
		return "Tailoring"
	case 18:
		return "Fletching"
	case 20:
		return "Jewelry"
	case 30:
		return "Pottery"
	case 24:
		return "Research"
	case 25:
		return "Research"
	case 26:
		return "Research"
	case 27:
		return "Research"
	case 46:
		return "Fishing"
	}
	return "Unknown Bagtype " + fmt.Sprintf("%d", t.Bagtype)
}

func (t *Item) EleDamageTypeStr() string {
	switch t.Elemdmgtype {
	case 0:
		return "Unknown"
	case 1:
		return "Magic"
	case 2:
		return "Fire"
	case 3:
		return "Cold"
	case 4:
		return "Poison"
	case 5:
		return "Disease"
	case 6:
		return "Corruption"
	}
	return "Unknown"
}

func (t *Item) BaneDamageTypeStr() string {
	return library.RaceStr(t.Banedmgrace)
}

var (
	dmg2h = []int{0,
		14,
		14,
		14,
		14,
		14,
		14,
		14,
		14,
		14, // 0->9
		14,
		14,
		14,
		14,
		14,
		14,
		14,
		14,
		14,
		14, // 10->19
		14,
		14,
		14,
		14,
		14,
		14,
		14,
		14,
		35,
		35, // 20->29
		36,
		36,
		37,
		37,
		38,
		38,
		39,
		39,
		40,
		40, // 30->39
		42,
		42,
		42,
		45,
		45,
		47,
		48,
		49,
		49,
		51, // 40->49
		51,
		52,
		53,
		54,
		54,
		56,
		56,
		57,
		58,
		59, // 50->59
		59,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0, // 60->69
		68,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0, // 70->79
		0,
		0,
		0,
		0,
		0,
		80,
		0,
		0,
		0,
		0, // 80->89
		0,
		0,
		0,
		0,
		0,
		88,
		0,
		0,
		0,
		0, // 90->99
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0, // 100->109
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0, // 110->119
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0, // 120->129
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0, // 130->139
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0,
		0, // 140->149
		132}
)

func (t *Item) DamageBonus() int {
	switch t.Itemtype {
	case 0, 2, 3, 42: // 1hs
		return 13
	case 1, 4, 35: // 2hs
		return dmg2h[t.Delay]

	}
	return 0
}

func (t *Item) ItemTypeStr() string {
	return library.ItemTypeStr(t.Itemtype)
}

func (t *Item) BardTypeStr() string {
	switch t.Bardtype {
	case 23:
		return "Wind"
	case 24:
		return "Strings"
	case 25:
		return "Brass"
	case 26:
		return "Percussion"
	case 51:
		return "All Instruments"
	}
	return fmt.Sprintf("Unknown %d", t.Bardtype)
}

func (t *Item) AugSlotStr() string {
	if t.Itemtype != 54 {
		return ""
	}
	if t.Augtype > 0 {
		Comma := ""
		AugSlots := ""
		AugType := t.Augtype
		Bit := int32(1)
		for i := 1; i < 25; i++ {
			if Bit <= AugType && Bit&AugType != 0 {
				AugSlots += Comma + fmt.Sprintf("%d", i)
				Comma = ", "
			}
			Bit *= 2
		}
		return fmt.Sprintf("Augmentation Slot Type: %s\n", AugSlots)
	}
	return "Augmentation Slot Type: All Slots\n"
}

func (t *Item) AugRestrictStr() string {
	if t.Itemtype != 54 {
		return ""
	}
	if t.Augrestrict < 1 {
		return ""
	}

	if t.Augrestrict > 12 {
		return fmt.Sprintf("Augmentation Restriction: Unknown Type %d\n", t.Augrestrict)
	}

	return fmt.Sprintf("Augmentation Restriction: %s", t.AugRestrictType(t.Augrestrict))
}

func (t *Item) AugRestrictType(val int32) string {
	switch val {
	case 1:
		return "Armor Only"
	case 2:
		return "Weapons Only"
	case 3:
		return "1h Weapons Only"
	case 4:
		return "2h Weapons Only"
	case 5:
		return "1h Slash Only"
	case 6:
		return "1h Blunt Only"
	case 7:
		return "Piercing Only"
	case 8:
		return "Hand To Hand Only"
	case 9:
		return "2h Slash Only"
	case 10:
		return "2h Blunt Only"
	case 11:
		return "2h Pierce Only"
	case 12:
		return "Bows Only"
	}
	return fmt.Sprintf("Unknown %d", val)

}

func (t *Item) ExtraDamageSkillStr() string {
	return library.SkillName(t.Extradmgskill)
}

func (t *Item) SkillModTypeStr() string {
	return library.SkillName(t.Skillmodtype)
}

func (t *Item) ProcRateTotal() int32 {
	return t.Procrate + 100
}
