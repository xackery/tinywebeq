package models

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/xackery/tinywebeq/eq/types"
	"github.com/xackery/tinywebeq/library"
)

type Item struct {
	ID                  int32              `json:"id" db:"items.id"`
	Minstatus           int16              `json:"-" db:"items.minstatus"`
	Name                string             `json:"name" db:"items.Name"`
	Aagi                int32              `json:"agi,omitempty" db:"items.aagi"`
	Ac                  int32              `json:"ac,omitempty" db:"items.ac"`
	Accuracy            int32              `json:"accuracy,omitempty" db:"items.accuracy"`
	Acha                int32              `json:"cha,omitempty" db:"items.acha"`
	Adex                int32              `json:"dex,omitempty" db:"items.adex"`
	Aint                int32              `json:"int,omitempty" db:"items.aint"`
	Artifactflag        uint8              `json:"artifact_flag,omitempty" db:"items.artifactflag"`
	Asta                int32              `json:"sta,omitempty" db:"items.asta"`
	Astr                int32              `json:"str,omitempty" db:"items.astr"`
	Attack              int32              `json:"attack,omitempty" db:"items.attack"`
	Augrestrict         int32              `json:"aug_restrictions,omitempty" db:"items.augrestrict"`
	Augslot1type        int8               `json:"aug_slot_1_type,omitempty" db:"items.augslot1type"`
	Augslot1visible     int8               `json:"aug_slot_1_visible,omitempty" db:"items.augslot1visible"`
	Augslot2type        int8               `json:"aug_slot_2_type,omitempty" db:"items.augslot2type"`
	Augslot2visible     int8               `json:"aug_slot_2_visible,omitempty" db:"items.augslot2visible"`
	Augslot3type        int8               `json:"aug_slot_3_type,omitempty" db:"items.augslot3type"`
	Augslot3visible     int8               `json:"aug_slot_3_visible,omitempty" db:"items.augslot3visible"`
	Augslot4type        int8               `json:"aug_slot_4_type,omitempty" db:"items.augslot4type"`
	Augslot4visible     int8               `json:"aug_slot_4_visible,omitempty" db:"items.augslot4visible"`
	Augslot5type        int8               `json:"aug_slot_5_type,omitempty" db:"items.augslot5type"`
	Augslot5visible     int8               `json:"aug_slot_5_visible,omitempty" db:"items.augslot5visible"`
	Augslot6type        int8               `json:"aug_slot_6_type,omitempty" db:"items.augslot6type"`
	Augslot6visible     int8               `json:"aug_slot_6_visible,omitempty" db:"items.augslot6visible"`
	Augtype             int32              `json:"augment_type,omitempty" db:"items.augtype"`
	Avoidance           int32              `json:"avoidance,omitempty" db:"items.avoidance"`
	Awis                int32              `json:"wis,omitempty" db:"items.awis"`
	Bagsize             int32              `json:"bag_size,omitempty" db:"items.bagsize"`
	Bagslots            int32              `json:"bag_slots,omitempty" db:"items.bagslots"`
	Bagtype             int32              `json:"bag_type,omitempty" db:"items.bagtype"`
	Bagwr               int32              `json:"bag_weight_reduction,omitempty" db:"items.bagwr"`
	Banedmgamt          int32              `json:"bane_damage_amount,omitempty" db:"items.banedmgamt"`
	Banedmgraceamt      int32              `json:"bane_damage_race_amount,omitempty" db:"items.banedmgraceamt"`
	Banedmgbody         int32              `json:"bane_damage_body_amount,omitempty" db:"items.banedmgbody"`
	Banedmgrace         types.NPCRace      `json:"bane_damage_race,omitempty" db:"items.banedmgrace"`
	Bardtype            int32              `json:"bard_type,omitempty" db:"items.bardtype"`
	Bardvalue           int32              `json:"bard_value,omitempty" db:"items.bardvalue"`
	Book                int32              `json:"book,omitempty" db:"items.book"`
	Casttime            int32              `json:"cast_time,omitempty" db:"items.casttime"`
	_                   int32              `db:"items.casttime_"`
	Charmfile           string             `json:"charm_file" db:"items.charmfile"`
	Charmfileid         string             `json:"charm_file_id" db:"items.charmfileid"`
	Classes             types.Class        `json:"classes,omitempty" db:"items.classes"`
	Color               uint32             `json:"color,omitempty" db:"items.color"`
	Combateffects       string             `json:"combat_effects,omitempty" db:"items.combateffects"`
	Extradmgskill       int32              `json:"extra_damage_skill,omitempty" db:"items.extradmgskill"`
	Extradmgamt         int32              `json:"extra_damage_amount,omitempty" db:"items.extradmgamt"`
	Price               int32              `json:"price,omitempty" db:"items.price"`
	Cr                  int32              `json:"cold_resist,omitempty" db:"items.cr"`
	Damage              int32              `json:"damage,omitempty" db:"items.damage"`
	Damageshield        int32              `json:"damage_shield,omitempty" db:"items.damageshield"`
	Deity               int32              `json:"deity,omitempty" db:"items.deity"`
	Delay               int32              `json:"delay,omitempty" db:"items.delay"`
	Augdistiller        uint32             `json:"augment_distiller,omitempty" db:"items.augdistiller"`
	Dotshielding        int32              `json:"dot_shielding,omitempty" db:"items.dotshielding"`
	Dr                  int32              `json:"disease_resist,omitempty" db:"items.dr"`
	Clicktype           int32              `json:"click_type,omitempty" db:"items.clicktype"`
	Clicklevel2         int32              `json:"click_level_2,omitempty" db:"items.clicklevel2"`
	Elemdmgtype         int32              `json:"elemental_damage_type,omitempty" db:"items.elemdmgtype"`
	Elemdmgamt          int32              `json:"elemental_damage_amount,omitempty" db:"items.elemdmgamt"`
	Endur               int32              `json:"endurance,omitempty" db:"items.endur"`
	Factionamt1         int32              `json:"faction_amount_1,omitempty" db:"items.factionamt1"`
	Factionamt2         int32              `json:"faction_amount_2,omitempty" db:"items.factionamt2"`
	Factionamt3         int32              `json:"faction_amount_3,omitempty" db:"items.factionamt3"`
	Factionamt4         int32              `json:"faction_amount_4,omitempty" db:"items.factionamt4"`
	Factionmod1         int32              `json:"faction_mod_1,omitempty" db:"items.factionmod1"`
	Factionmod2         int32              `json:"faction_mod_2,omitempty" db:"items.factionmod2"`
	Factionmod3         int32              `json:"faction_mod_3,omitempty" db:"items.factionmod3"`
	Factionmod4         int32              `json:"faction_mod_4,omitempty" db:"items.factionmod4"`
	Filename            string             `json:"filename,omitempty" db:"items.filename"`
	Focuseffect         int32              `json:"focus_effect,omitempty" db:"items.focuseffect"`
	Fr                  int32              `json:"fire_resist,omitempty" db:"items.fr"`
	Fvnodrop            int32              `json:"fv_no_drop,omitempty" db:"items.fvnodrop"`
	Haste               int32              `json:"haste,omitempty" db:"items.haste"`
	Clicklevel          int32              `json:"click_level,omitempty" db:"items.clicklevel"`
	Hp                  int32              `json:"hp,omitempty" db:"items.hp"`
	Regen               int32              `json:"regen,omitempty" db:"items.regen"`
	Icon                int32              `json:"icon,omitempty" db:"items.icon"`
	Idfile              string             `json:"id_file,omitempty" db:"items.idfile"`
	Itemclass           int32              `json:"item_class,omitempty" db:"items.itemclass"`
	Itemtype            types.ItemType     `json:"item_type,omitempty" db:"items.itemtype"`
	Ldonprice           int32              `json:"ldon_price,omitempty" db:"items.ldonprice"`
	Ldontheme           int32              `json:"ldon_theme,omitempty" db:"items.ldontheme"`
	Ldonsold            int32              `json:"ldon_sold," db:"items.ldonsold"`
	Light               int32              `json:"light,omitempty" db:"items.light"`
	Lore                string             `json:"lore,omitempty" db:"items.lore"`
	Loregroup           int32              `json:"lore_group,omitempty" db:"items.loregroup"`
	Magic               int32              `json:"magic,omitempty" db:"items.magic"`
	Mana                int32              `json:"mana,omitempty" db:"items.mana"`
	Manaregen           int32              `json:"mana_regen,omitempty" db:"items.manaregen"`
	Enduranceregen      int32              `json:"endurance_regen,omitempty" db:"items.enduranceregen"`
	Material            int32              `json:"material,omitempty" db:"items.material"`
	Herosforgemodel     int32              `json:"heroes_forge_model,omitempty" db:"items.herosforgemodel"`
	Maxcharges          int32              `json:"max_charges,omitempty" db:"items.maxcharges"`
	Mr                  int32              `json:"magic_resist,omitempty" db:"items.mr"`
	Nodrop              int32              `json:"no_drop,omitempty" db:"items.nodrop"`
	Norent              int32              `json:"no_rent,omitempty" db:"items.norent"`
	Pendingloreflag     uint8              `json:"pending_lore_flag,omitempty" db:"items.pendingloreflag"`
	Pr                  int32              `json:"poison_resist,omitempty" db:"items.pr"`
	Procrate            int32              `json:"proc_rate,omitempty" db:"items.procrate"`
	Races               types.RacesBitmask `json:"races,omitempty" db:"items.races"`
	Range               int32              `json:"range,omitempty" db:"items.range"`
	Reclevel            int32              `json:"rec_level,omitempty" db:"items.reclevel"`
	Recskill            int32              `json:"rec_skill,omitempty" db:"items.recskill"`
	Reqlevel            int32              `json:"req_level,omitempty" db:"items.reqlevel"`
	Sellrate            float64            `json:"sell_rate,omitempty" db:"items.sellrate"`
	Shielding           int32              `json:"shielding,omitempty" db:"items.shielding"`
	Size                int32              `json:"size" db:"items.size"`
	Skillmodtype        int32              `json:"skill_mod_type,omitempty" db:"items.skillmodtype"`
	Skillmodvalue       int32              `json:"skill_mod_value,omitempty" db:"items.skillmodvalue"`
	Slots               int32              `json:"slots,omitempty" db:"items.slots"`
	Clickeffect         int32              `json:"click_effect,omitempty" db:"items.clickeffect"`
	Spellshield         int32              `json:"spellshield,omitempty" db:"items.spellshield"`
	Strikethrough       int32              `json:"strikethrough,omitempty" db:"items.strikethrough"`
	Stunresist          int32              `json:"stunresist,omitempty" db:"items.stunresist"`
	Summonedflag        uint8              `json:"summoned_flag,omitempty" db:"items.summonedflag"`
	Tradeskills         int32              `json:"tradeskills,omitempty" db:"items.tradeskills"`
	Favor               int32              `json:"favor,omitempty" db:"items.favor"`
	Weight              types.Weight       `json:"weight,omitempty" db:"items.weight"`
	_                   int32              `db:"items.unk012"`
	_                   int32              `db:"items.unk013"`
	Benefitflag         int32              `json:"benefit_flag,omitempty" db:"items.benefitflag"`
	_                   int32              `db:"items.unk054"`
	_                   int32              `db:"items.unk059"`
	Booktype            int32              `json:"book_type,omitempty" db:"items.booktype"`
	Recastdelay         int32              `json:"recast_delay,omitempty" db:"items.recastdelay"`
	Recasttype          int32              `json:"recast_type,omitempty" db:"items.recasttype"`
	Guildfavor          int32              `json:"guild_favor,omitempty" db:"items.guildfavor"`
	_                   int32              `db:"items.unk123"`
	_                   int32              `db:"items.unk124"`
	Attuneable          int32              `json:"attuneable,omitempty" db:"items.attuneable"`
	Nopet               int32              `json:"no_pet,omitempty" db:"items.nopet"`
	Updated             sql.NullTime       `json:"updated,omitempty" db:"items.updated"`
	Comment             string             `json:"comment,omitempty" db:"items.comment"`
	_                   int32              `db:"items.unk127"`
	Pointtype           int32              `json:"potion_type,omitempty" db:"items.pointtype"`
	Potionbelt          int32              `json:"potion_belt,omitempty" db:"items.potionbelt"`
	Potionbeltslots     int32              `json:"potion_belt_slots,omitempty" db:"items.potionbeltslots"`
	Stacksize           int32              `json:"stack_size,omitempty" db:"items.stacksize"`
	Notransfer          int32              `json:"no_transfer,omitempty" db:"items.notransfer"`
	Stackable           int32              `json:"stackable,omitempty" db:"items.stackable"`
	_                   string             `db:"items.unk134"`
	_                   int32              `db:"items.unk137"`
	Proceffect          int32              `json:"proc_effect,omitempty" db:"items.proceffect"`
	Proctype            int32              `json:"proc_type,omitempty" db:"items.proctype"`
	Proclevel2          int32              `json:"proc_level_2,omitempty" db:"items.proclevel2"`
	Proclevel           int32              `json:"proc_level,omitempty" db:"items.proclevel"`
	_                   int32              `db:"items.unk142"`
	Worneffect          int32              `json:"worn_effect,omitempty" db:"items.worneffect"`
	Worntype            int32              `json:"worn_type,omitempty" db:"items.worntype"`
	Wornlevel2          int32              `json:"worn_level_2,omitempty" db:"items.wornlevel2"`
	Wornlevel           int32              `json:"worn_level,omitempty" db:"items.wornlevel"`
	_                   int32              `db:"items.unk147"`
	Focustype           int32              `json:"focus_type,omitempty" db:"items.focustype"`
	Focuslevel2         int32              `json:"focus_level_2,omitempty" db:"items.focuslevel2"`
	Focuslevel          int32              `json:"focus_level,omitempty" db:"items.focuslevel"`
	_                   int32              `db:"items.unk152"`
	Scrolleffect        int32              `json:"scroll_effect,omitempty" db:"items.scrolleffect"`
	Scrolltype          int32              `json:"scroll_type,omitempty" db:"items.scrolltype"`
	Scrolllevel2        int32              `json:"scroll_level_2,omitempty" db:"items.scrolllevel2"`
	Scrolllevel         int32              `json:"scroll_level,omitempty" db:"items.scrolllevel"`
	_                   int32              `db:"items.unk157"`
	Serialized          sql.NullTime       `json:"serialized,omitempty" db:"items.serialized"`
	Verified            sql.NullTime       `json:"verified,omitempty" db:"items.verified"`
	Serialization       sql.NullString     `json:"serialization,omitempty" db:"items.serialization"`
	Source              string             `json:"source,omitempty" db:"items.source"`
	_                   int32              `db:"items.unk033"`
	Lorefile            string             `json:"lore_file,omitempty" db:"items.lorefile"`
	_                   int32              `db:"items.unk014"`
	Svcorruption        int32              `json:"resist_corruption,omitempty" db:"items.svcorruption"`
	Skillmodmax         int32              `json:"skill_mod_max,omitempty" db:"items.skillmodmax"`
	_                   int32              `db:"items.unk060"`
	_                   int32              `db:"items.augslot1unk2"`
	_                   int32              `db:"items.augslot2unk2"`
	_                   int32              `db:"items.augslot3unk2"`
	_                   int32              `db:"items.augslot4unk2"`
	_                   int32              `db:"items.augslot5unk2"`
	_                   int32              `db:"items.augslot6unk2"`
	_                   int32              `db:"items.unk120"`
	_                   int32              `db:"items.unk121"`
	Questitemflag       int32              `json:"quest_item_flag,omitempty" db:"items.questitemflag"`
	_                   sql.NullString     `db:"items.unk132"`
	_                   int32              `db:"items.clickunk5"`
	_                   string             `db:"items.clickunk6"`
	_                   int32              `db:"items.clickunk7"`
	_                   int32              `db:"items.procunk1"`
	_                   int32              `db:"items.procunk2"`
	_                   int32              `db:"items.procunk3"`
	_                   int32              `db:"items.procunk4"`
	_                   string             `db:"items.procunk6"`
	_                   int32              `db:"items.procunk7"`
	_                   int32              `db:"items.wornunk1"`
	_                   int32              `db:"items.wornunk2"`
	_                   int32              `db:"items.wornunk3"`
	_                   int32              `db:"items.wornunk4"`
	_                   int32              `db:"items.wornunk5"`
	_                   string             `db:"items.wornunk6"`
	_                   int32              `db:"items.wornunk7"`
	_                   int32              `db:"items.focusunk1"`
	_                   int32              `db:"items.focusunk2"`
	_                   int32              `db:"items.focusunk3"`
	_                   int32              `db:"items.focusunk4"`
	_                   int32              `db:"items.focusunk5"`
	_                   string             `db:"items.focusunk6"`
	_                   int32              `db:"items.focusunk7"`
	_                   uint32             `db:"items.scrollunk1"`
	_                   int32              `db:"items.scrollunk2"`
	_                   int32              `db:"items.scrollunk3"`
	_                   int32              `db:"items.scrollunk4"`
	_                   int32              `db:"items.scrollunk5"`
	_                   string             `db:"items.scrollunk6"`
	_                   int32              `db:"items.scrollunk7"`
	_                   int32              `db:"items.unk193"`
	Purity              int32              `json:"purity,omitempty" db:"items.purity"`
	Evoitem             int32              `json:"evolving_item,omitempty" db:"items.evoitem"`
	Evoid               int32              `json:"evolving_id,omitempty" db:"items.evoid"`
	Evolvinglevel       int32              `json:"evolving_level,omitempty" db:"items.evolvinglevel"`
	Evomax              int32              `json:"evolving_max,omitempty" db:"items.evomax"`
	Clickname           string             `json:"click_name,omitempty" db:"items.clickname"`
	Procname            string             `json:"proc_name,omitempty" db:"items.procname"`
	Wornname            string             `json:"worn_name,omitempty" db:"items.wornname"`
	Focusname           string             `json:"focus_name,omitempty" db:"items.focusname"`
	Scrollname          string             `json:"scroll_name,omitempty" db:"items.scrollname"`
	Dsmitigation        int16              `json:"damage_shield_mitigation,omitempty" db:"items.dsmitigation"`
	HeroicStr           int16              `json:"heroic_str,omitempty" db:"items.heroicstr"`
	HeroicInt           int16              `json:"heroic_int,omitempty" db:"items.heroicint"`
	HeroicWis           int16              `json:"heroic_wis,omitempty" db:"items.heroicwis"`
	HeroicAgi           int16              `json:"heroic_agi,omitempty" db:"items.heroicagi"`
	HeroicDex           int16              `json:"heroic_dex,omitempty" db:"items.heroicdex"`
	HeroicSta           int16              `json:"heroic_sta,omitempty" db:"items.heroicsta"`
	HeroicCha           int16              `json:"heroic_cha,omitempty" db:"items.heroiccha"`
	HeroicPr            int16              `json:"heroic_poison_resist,omitempty" db:"items.heroicpr"`
	HeroicDr            int16              `json:"heroic_disease_resist,omitempty" db:"items.heroicdr"`
	HeroicFr            int16              `json:"heroic_fire_resist,omitempty" db:"items.heroicfr"`
	HeroicCr            int16              `json:"heroic_cold_resist,omitempty" db:"items.heroiccr"`
	HeroicMr            int16              `json:"heroic_magic_resist,omitempty" db:"items.heroicmr"`
	HeroicSvcorrup      int16              `json:"heroic_corruption_resist,omitempty" db:"items.heroicsvcorrup"`
	Healamt             int16              `json:"heal_amount,omitempty" db:"items.healamt"`
	Spelldmg            int16              `json:"spell_damage,omitempty" db:"items.spelldmg"`
	Clairvoyance        int16              `json:"clairvoyance,omitempty" db:"items.clairvoyance"`
	Backstabdmg         int16              `json:"backstab_damage,omitempty" db:"items.backstabdmg"`
	Created             string             `json:"created,omitempty" db:"items.created"`
	Elitematerial       int16              `json:"elite_material,omitempty" db:"items.elitematerial"`
	Ldonsellbackrate    int16              `json:"ldon_sellback_rate,omitempty" db:"items.ldonsellbackrate"`
	Scriptfileid        int32              `json:"script_field_id" db:"items.scriptfileid"`
	Expendablearrow     int16              `json:"expendable_arrow,omitempty" db:"items.expendablearrow"`
	Powersourcecapacity int32              `json:"power_source_capacity,omitempty" db:"items.powersourcecapacity"`
	Bardeffect          int32              `json:"bard_effect,omitempty" db:"items.bardeffect"`
	Bardeffecttype      int16              `json:"bard_effect_type,omitempty" db:"items.bardeffecttype"`
	Bardlevel2          int16              `json:"bard_level_2,omitempty" db:"items.bardlevel2"`
	Bardlevel           int16              `json:"bard_level,omitempty" db:"items.bardlevel"`
	_                   int16              `db:"items.bardunk1"`
	_                   int16              `db:"items.bardunk2"`
	_                   int16              `db:"items.bardunk3"`
	_                   int16              `db:"items.bardunk4"`
	_                   int16              `db:"items.bardunk5"`
	_                   string             `db:"items.bardname"`
	_                   int16              `db:"items.bardunk7"`
	_                   int16              `db:"items.unk214"`
	Subtype             int32              `json:"subtype" db:"items.subtype"`
	_                   int32              `db:"items.unk220"`
	_                   int32              `db:"items.unk221"`
	Heirloom            int32              `json:"heirloom,omitempty" db:"items.heirloom"`
	_                   int32              `db:"items.unk223"`
	_                   int32              `db:"items.unk224"`
	_                   int32              `db:"items.unk225"`
	_                   int32              `db:"items.unk226"`
	_                   int32              `db:"items.unk227"`
	_                   int32              `db:"items.unk228"`
	_                   int32              `db:"items.unk229"`
	_                   int32              `db:"items.unk230"`
	_                   int32              `db:"items.unk231"`
	_                   int32              `db:"items.unk232"`
	_                   int32              `db:"items.unk233"`
	_                   int32              `db:"items.unk234"`
	Placeable           int32              `json:"placeable,omitempty" db:"items.placeable"`
	_                   int32              `db:"items.unk236"`
	_                   int32              `db:"items.unk237"`
	_                   int32              `db:"items.unk238"`
	_                   int32              `db:"items.unk239"`
	_                   int32              `db:"items.unk240"`
	_                   int32              `db:"items.unk241"`
	Epicitem            int32              `json:"epic_item" db:"items.epicitem"`
}

type DiscoveredItem struct {
	*Item
	ItemID         uint32 `json:"item_id,omitempty"`
	CharName       string `json:"-"`
	DiscoveredDate uint32 `json:"-"`
	AccountStatus  int32  `json:"-"`
}

// MarshalJSON implements a custom item marshaler for displaying the item with JSON encoding.
func (t *Item) MarshalJSON() ([]byte, error) {
	type Alias Item
	i := struct {
		Alias
		Slots library.Slots `json:"slots"`
	}{
		Alias: Alias(*t),
		Slots: library.SlotsFromBitmask(t.Slots),
	}

	return json.Marshal(i)
}

func (t *Item) Identifier() string {
	return "item"
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
	return library.SlotsFromBitmask(t.Slots).String()
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
	return "https://everquest.allakhazam.com/pgfx/"
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
