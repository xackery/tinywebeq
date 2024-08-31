package model

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/xackery/tinywebeq/library"
)

type Item struct {
	CacheKey            string         `json:"-"`
	CacheExpiration     int64          `json:"-"`
	ID                  int32          `json:"id"`
	Minstatus           int16          `json:"-"`
	Name                string         `json:"name"`
	Aagi                int32          `json:"agi,omitempty"`
	Ac                  int32          `json:"ac,omitempty"`
	Accuracy            int32          `json:"accuracy,omitempty"`
	Acha                int32          `json:"cha,omitempty"`
	Adex                int32          `json:"dex,omitempty"`
	Aint                int32          `json:"int,omitempty"`
	Artifactflag        uint8          `json:"artifact_flag,omitempty"`
	Asta                int32          `json:"sta,omitempty"`
	Astr                int32          `json:"str,omitempty"`
	Attack              int32          `json:"attack,omitempty"`
	Augrestrict         int32          `json:"augment_restrictions,omitempty"`
	Augslot1type        int8           `json:"augment_slot_1_type,omitempty"`
	Augslot1visible     int8           `json:"-"`
	Augslot2type        int8           `json:"augment_slot_2_type,omitempty"`
	Augslot2visible     int8           `json:"-"`
	Augslot3type        int8           `json:"augment_slot_3_type,omitempty"`
	Augslot3visible     int8           `json:"-"`
	Augslot4type        int8           `json:"augment_slot_4_type,omitempty"`
	Augslot4visible     int8           `json:"-"`
	Augslot5type        int8           `json:"augment_slot_5_type,omitempty"`
	Augslot5visible     int8           `json:"-"`
	Augslot6type        int8           `json:"augment_slot_6_type,omitempty"`
	Augslot6visible     int8           `json:"-"`
	Augtype             int32          `json:"augment_type,omitempty"`
	Avoidance           int32          `json:"avoidance,omitempty"`
	Awis                int32          `json:"wis,omitempty"`
	Bagsize             int32          `json:"bag_size,omitempty"`
	Bagslots            int32          `json:"bag_slots,omitempty"`
	Bagtype             int32          `json:"bag_type,omitempty"`
	Bagwr               int32          `json:"bag_weight_reduction,omitempty"`
	Banedmgamt          int32          `json:"bane_damage_amount,omitempty"`
	Banedmgraceamt      int32          `json:"bane_damage_race_amount,omitempty"`
	Banedmgbody         int32          `json:"bane_damage_body_amount,omitempty"`
	Banedmgrace         int32          `json:"bane_damage_race,omitempty"`
	Bardtype            int32          `json:"bard_type,omitempty"`
	Bardvalue           int32          `json:"bard_value,omitempty"`
	Book                int32          `json:"book,omitempty"`
	Casttime            int32          `json:"cast_time,omitempty"`
	Charmfile           string         `json:"-"`
	Charmfileid         string         `json:"-"`
	Classes             int32          `json:"classes,omitempty"`
	Color               uint32         `json:"-"`
	Combateffects       string         `json:"combat_effects,omitempty"`
	Extradmgskill       int32          `json:"extra_damage_skill,omitempty"`
	Extradmgamt         int32          `json:"extra_damage_race,omitempty"`
	Price               int32          `json:"-"`
	Cr                  int32          `json:"cold_resist,omitempty"`
	Damage              int32          `json:"damage,omitempty"`
	Damageshield        int32          `json:"damage_shield,omitempty"`
	Deity               int32          `json:"deity,omitempty"`
	Delay               int32          `json:"delay,omitempty"`
	Augdistiller        uint32         `json:"augment_distiller,omitempty"`
	Dotshielding        int32          `json:"dot_shielding,omitempty"`
	Dr                  int32          `json:"disease_resist,omitempty"`
	Clicktype           int32          `json:"click_type,omitempty"`
	Clicklevel2         int32          `json:"click_level_2,omitempty"`
	Elemdmgtype         int32          `json:"elemental_damage_type,omitempty"`
	Elemdmgamt          int32          `json:"elemental_damage_race,omitempty"`
	Endur               int32          `json:"endurance,omitempty"`
	Factionamt1         int32          `json:"faction_amount_1,omitempty"`
	Factionamt2         int32          `json:"faction_amount_2,omitempty"`
	Factionamt3         int32          `json:"faction_amount_3,omitempty"`
	Factionamt4         int32          `json:"faction_amount_4,omitempty"`
	Factionmod1         int32          `json:"faction_mod_1,omitempty"`
	Factionmod2         int32          `json:"faction_mod_2,omitempty"`
	Factionmod3         int32          `json:"faction_mod_3,omitempty"`
	Factionmod4         int32          `json:"faction_mod_4,omitempty"`
	Filename            string         `json:"-"`
	Focuseffect         int32          `json:"focus_effect,omitempty"`
	Fr                  int32          `json:"fire_resist,omitempty"`
	Fvnodrop            int32          `json:"fv_no_drop,omitempty"`
	Haste               int32          `json:"haste,omitempty"`
	Clicklevel          int32          `json:"click_level,omitempty"`
	Hp                  int32          `json:"hp,omitempty"`
	Regen               int32          `json:"regen,omitempty"`
	Icon                int32          `json:"icon,omitempty"`
	Idfile              string         `json:"-"`
	Itemclass           int32          `json:"item_class,omitempty"`
	Itemtype            int32          `json:"item_type,omitempty"`
	Ldonprice           int32          `json:"-"`
	Ldontheme           int32          `json:"-"`
	Ldonsold            int32          `json:"-"`
	Light               int32          `json:"light,omitempty"`
	Lore                string         `json:"lore,omitempty"`
	Loregroup           int32          `json:"-"`
	Magic               int32          `json:"magic,omitempty"`
	Mana                int32          `json:"mana,omitempty"`
	Manaregen           int32          `json:"mana_regen,omitempty"`
	Enduranceregen      int32          `json:"endurance_regen,omitempty"`
	Material            int32          `json:"-"`
	Herosforgemodel     int32          `json:"-"`
	Maxcharges          int32          `json:"maxcharges,omitempty"`
	Mr                  int32          `json:"magic_resist,omitempty"`
	Nodrop              int32          `json:"no_drop,omitempty"`
	Norent              int32          `json:"no_rent,omitempty"`
	Pendingloreflag     uint8          `json:"pending_lore_flag,omitempty"`
	Pr                  int32          `json:"poison_resist,omitempty"`
	Procrate            int32          `json:"-"`
	Races               int32          `json:"races,omitempty"`
	Range               int32          `json:"range,omitempty"`
	Reclevel            int32          `json:"rec_level,omitempty"`
	Recskill            int32          `json:"rec_skill,omitempty"`
	Reqlevel            int32          `json:"req_level,omitempty"`
	Sellrate            float64        `json:"-"`
	Shielding           int32          `json:"shielding,omitempty"`
	Size                int32          `json:"size"`
	Skillmodtype        int32          `json:"skill_mod_type,omitempty"`
	Skillmodvalue       int32          `json:"skill_mod_value,omitempty"`
	Slots               int32          `json:"slots,omitempty"`
	Clickeffect         int32          `json:"click_effect,omitempty"`
	Spellshield         int32          `json:"spellshield,omitempty"`
	Strikethrough       int32          `json:"strikethrough,omitempty"`
	Stunresist          int32          `json:"stunresist,omitempty"`
	Summonedflag        uint8          `json:"summoned_flag,omitempty"`
	Tradeskills         int32          `json:"tradeskills,omitempty"`
	Favor               int32          `json:"-"`
	Weight              int32          `json:"weight,omitempty"`
	Unk012              int32          `json:"-"`
	Unk013              int32          `json:"-"`
	Benefitflag         int32          `json:"benefit_flag,omitempty"`
	Unk054              int32          `json:"-"`
	Unk059              int32          `json:"-"`
	Booktype            int32          `json:"book_type,omitempty"`
	Recastdelay         int32          `json:"recast_delay,omitempty"`
	Recasttype          int32          `json:"recast_type,omitempty"`
	Guildfavor          int32          `json:"-"`
	Unk123              int32          `json:"-"`
	Unk124              int32          `json:"-"`
	Attuneable          int32          `json:"attuneable,omitempty"`
	Nopet               int32          `json:"-"`
	Updated             sql.NullTime   `json:"-"`
	Comment             string         `json:"-"`
	Unk127              int32          `json:"-"`
	Pointtype           int32          `json:"-"`
	Potionbelt          int32          `json:"-"`
	Potionbeltslots     int32          `json:"-"`
	Stacksize           int32          `json:"-"`
	Notransfer          int32          `json:"no_transfer,omitempty"`
	Stackable           int32          `json:"-"`
	Unk134              string         `json:"-"`
	Unk137              int32          `json:"-"`
	Proceffect          int32          `json:"proc_effect,omitempty"`
	Proctype            int32          `json:"proc_type,omitempty"`
	Proclevel2          int32          `json:"proc_level_2,omitempty"`
	Proclevel           int32          `json:"proc_level,omitempty"`
	Unk142              int32          `json:"-"`
	Worneffect          int32          `json:"worn_effect,omitempty"`
	Worntype            int32          `json:"worn_type,omitempty"`
	Wornlevel2          int32          `json:"worn_level_2,omitempty"`
	Wornlevel           int32          `json:"worn_level,omitempty"`
	Unk147              int32          `json:"-"`
	Focustype           int32          `json:"focus_type,omitempty"`
	Focuslevel2         int32          `json:"focus_level_2,omitempty"`
	Focuslevel          int32          `json:"focus_level,omitempty"`
	Unk152              int32          `json:"-"`
	Scrolleffect        int32          `json:"-"`
	Scrolltype          int32          `json:"-"`
	Scrolllevel2        int32          `json:"-"`
	Scrolllevel         int32          `json:"-"`
	Unk157              int32          `json:"-"`
	Serialized          sql.NullTime   `json:"-"`
	Verified            sql.NullTime   `json:"-"`
	Serialization       sql.NullString `json:"-"`
	Source              string         `json:"-"`
	Unk033              int32          `json:"-"`
	Lorefile            string         `json:"-"`
	Unk014              int32          `json:"-"`
	Svcorruption        int32          `json:"resist_corruption,omitempty"`
	Skillmodmax         int32          `json:"skill_mod_max,omitempty"`
	Unk060              int32          `json:"-"`
	Augslot1unk2        int32          `json:"-"`
	Augslot2unk2        int32          `json:"-"`
	Augslot3unk2        int32          `json:"-"`
	Augslot4unk2        int32          `json:"-"`
	Augslot5unk2        int32          `json:"-"`
	Augslot6unk2        int32          `json:"-"`
	Unk120              int32          `json:"-"`
	Unk121              int32          `json:"-"`
	Questitemflag       int32          `json:"quest_item_flag,omitempty"`
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
	Purity              int32          `json:"purity,omitempty"`
	Evoitem             int32          `json:"evolving_item,omitempty"`
	Evoid               int32          `json:"evolving_id,omitempty"`
	Evolvinglevel       int32          `json:"evolving_level,omitempty"`
	Evomax              int32          `json:"evolving_max,omitempty"`
	Clickname           string         `json:"click_name,omitempty"`
	Procname            string         `json:"proc_name,omitempty"`
	Wornname            string         `json:"worn_name,omitempty"`
	Focusname           string         `json:"focus_name,omitempty"`
	Scrollname          string         `json:"scroll_name,omitempty"`
	Dsmitigation        int16          `json:"damage_shield_mitigation,omitempty"`
	HeroicStr           int16          `json:"heroic_str,omitempty"`
	HeroicInt           int16          `json:"heroic_int,omitempty"`
	HeroicWis           int16          `json:"heroic_wis,omitempty"`
	HeroicAgi           int16          `json:"heroic_agi,omitempty"`
	HeroicDex           int16          `json:"heroic_dex,omitempty"`
	HeroicSta           int16          `json:"heroic_sta,omitempty"`
	HeroicCha           int16          `json:"heroic_cha,omitempty"`
	HeroicPr            int16          `json:"heroic_poison_resist,omitempty"`
	HeroicDr            int16          `json:"heroic_disease_resist,omitempty"`
	HeroicFr            int16          `json:"heroic_fire_resist,omitempty"`
	HeroicCr            int16          `json:"heroic_cold_resist,omitempty"`
	HeroicMr            int16          `json:"heroic_magic_resist,omitempty"`
	HeroicSvcorrup      int16          `json:"heroic_corruption_resist,omitempty"`
	Healamt             int16          `json:"heal_amount,omitempty"`
	Spelldmg            int16          `json:"spell_damage,omitempty"`
	Clairvoyance        int16          `json:"clairvoyance,omitempty"`
	Backstabdmg         int16          `json:"backstab_damage,omitempty"`
	Created             string         `json:"-"`
	Elitematerial       int16          `json:"-"`
	Ldonsellbackrate    int16          `json:"-"`
	Scriptfileid        int32          `json:"-"`
	Expendablearrow     int16          `json:"-"`
	Powersourcecapacity int32          `json:"-"`
	Bardeffect          int32          `json:"-"`
	Bardeffecttype      int16          `json:"-"`
	Bardlevel2          int16          `json:"-"`
	Bardlevel           int16          `json:"-"`
	Bardunk1            int16          `json:"-"`
	Bardunk2            int16          `json:"-"`
	Bardunk3            int16          `json:"-"`
	Bardunk4            int16          `json:"-"`
	Bardunk5            int16          `json:"-"`
	Bardname            string         `json:"-"`
	Bardunk7            int16          `json:"-"`
	Unk214              int16          `json:"-"`
	Subtype             int32          `json:"-"`
	Unk220              int32          `json:"-"`
	Unk221              int32          `json:"-"`
	Heirloom            int32          `json:"heirloom,omitempty"`
	Unk223              int32          `json:"-"`
	Unk224              int32          `json:"-"`
	Unk225              int32          `json:"-"`
	Unk226              int32          `json:"-"`
	Unk227              int32          `json:"-"`
	Unk228              int32          `json:"-"`
	Unk229              int32          `json:"-"`
	Unk230              int32          `json:"-"`
	Unk231              int32          `json:"-"`
	Unk232              int32          `json:"-"`
	Unk233              int32          `json:"-"`
	Unk234              int32          `json:"-"`
	Placeable           int32          `json:"placeable,omitempty"`
	Unk236              int32          `json:"-"`
	Unk237              int32          `json:"-"`
	Unk238              int32          `json:"-"`
	Unk239              int32          `json:"-"`
	Unk240              int32          `json:"-"`
	Unk241              int32          `json:"-"`
	Epicitem            int32          `json:"-"`
	ItemID              uint32         `json:"item_id,omitempty"`
	CharName            string         `json:"-"`
	DiscoveredDate      uint32         `json:"-"`
	AccountStatus       int32          `json:"-"`
}

// MarshalJSON implements a custom item marshaler for displaying the item with JSON encoding.
func (t *Item) MarshalJSON() ([]byte, error) {
	type Alias Item
	i := struct {
		Alias
		Weight  float64         `json:"weight,omitempty"`
		Classes library.Classes `json:"classes,omitempty"`
		Races   library.Races   `json:"races,omitempty"`
		Slots   library.Slots   `json:"slots"`
	}{
		Alias:   Alias(*t),
		Weight:  float64(t.Weight) / 10,
		Classes: library.ClassesFromBitmask(t.Classes),
		Races:   library.RacesFromBitmask(t.Races),
		Slots:   library.SlotsFromBitmask(t.Slots),
	}

	return json.Marshal(i)
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
	return library.ClassesFromBitmask(t.Classes).String()
}

func (t *Item) RaceStr() string {
	return library.RacesFromBitmask(t.Races).String()
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
