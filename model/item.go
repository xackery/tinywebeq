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
	key                 string
	expiration          int64
	ID                  int            `db:"id"`
	Minstatus           int            `db:"minstatus"`
	Name                string         `db:"Name"`
	Aagi                int            `db:"aagi"`
	Ac                  int            `db:"ac"`
	Accuracy            int            `db:"accuracy"`
	Acha                int            `db:"acha"`
	Adex                int            `db:"adex"`
	Aint                int            `db:"aint"`
	Artifactflag        int            `db:"artifactflag"`
	Asta                int            `db:"asta"`
	Astr                int            `db:"astr"`
	Attack              int            `db:"attack"`
	Augrestrict         int            `db:"augrestrict"`
	Augslot1type        int            `db:"augslot1type"`
	Augslot1visible     int            `db:"augslot1visible"`
	Augslot2type        int            `db:"augslot2type"`
	Augslot2visible     int            `db:"augslot2visible"`
	Augslot3type        int            `db:"augslot3type"`
	Augslot3visible     int            `db:"augslot3visible"`
	Augslot4type        int            `db:"augslot4type"`
	Augslot4visible     int            `db:"augslot4visible"`
	Augslot5type        int            `db:"augslot5type"`
	Augslot5visible     int            `db:"augslot5visible"`
	Augslot6type        int            `db:"augslot6type"`
	Augslot6visible     int            `db:"augslot6visible"`
	Augtype             int            `db:"augtype"`
	Avoidance           int            `db:"avoidance"`
	Awis                int            `db:"awis"`
	Bagsize             int            `db:"bagsize"`
	Bagslots            int            `db:"bagslots"`
	Bagtype             int            `db:"bagtype"`
	Bagwr               int            `db:"bagwr"`
	Banedmgamt          int            `db:"banedmgamt"`
	Banedmgraceamt      int            `db:"banedmgraceamt"`
	Banedmgbody         int            `db:"banedmgbody"`
	Banedmgrace         int            `db:"banedmgrace"`
	Bardtype            int            `db:"bardtype"`
	Bardvalue           int            `db:"bardvalue"`
	Book                int            `db:"book"`
	Casttime            int            `db:"casttime"`
	Casttime_           int            `db:"casttime_"`
	Charmfile           string         `db:"charmfile"`
	Charmfileid         string         `db:"charmfileid"`
	Classes             int            `db:"classes"`
	Color               int            `db:"color"`
	Combateffects       string         `db:"combateffects"`
	Extradmgskill       int            `db:"extradmgskill"`
	Extradmgamt         int            `db:"extradmgamt"`
	Price               int            `db:"price"`
	Cr                  int            `db:"cr"`
	Damage              int            `db:"damage"`
	Damageshield        int            `db:"damageshield"`
	Deity               int            `db:"deity"`
	Delay               int            `db:"delay"`
	Augdistiller        int            `db:"augdistiller"`
	Dotshielding        int            `db:"dotshielding"`
	Dr                  int            `db:"dr"`
	Clicktype           int            `db:"clicktype"`
	Clicklevel2         int            `db:"clicklevel2"`
	Elemdmgtype         int            `db:"elemdmgtype"`
	Elemdmgamt          int            `db:"elemdmgamt"`
	Endur               int            `db:"endur"`
	Factionamt1         int            `db:"factionamt1"`
	Factionamt2         int            `db:"factionamt2"`
	Factionamt3         int            `db:"factionamt3"`
	Factionamt4         int            `db:"factionamt4"`
	Factionmod1         int            `db:"factionmod1"`
	Factionmod2         int            `db:"factionmod2"`
	Factionmod3         int            `db:"factionmod3"`
	Factionmod4         int            `db:"factionmod4"`
	Filename            string         `db:"filename"`
	Focuseffect         int            `db:"focuseffect"`
	Fr                  int            `db:"fr"`
	Fvnodrop            int            `db:"fvnodrop"`
	Haste               int            `db:"haste"`
	Clicklevel          int            `db:"clicklevel"`
	Hp                  int            `db:"hp"`
	Regen               int            `db:"regen"`
	Icon                int            `db:"icon"`
	Idfile              string         `db:"idfile"`
	Itemclass           int            `db:"itemclass"`
	Itemtype            int            `db:"itemtype"`
	Ldonprice           int            `db:"ldonprice"`
	Ldontheme           int            `db:"ldontheme"`
	Ldonsold            int            `db:"ldonsold"`
	Light               int            `db:"light"`
	Lore                string         `db:"lore"`
	Loregroup           int            `db:"loregroup"`
	Magic               int            `db:"magic"`
	Mana                int            `db:"mana"`
	Manaregen           int            `db:"manaregen"`
	Enduranceregen      int            `db:"enduranceregen"`
	Material            int            `db:"material"`
	Herosforgemodel     int            `db:"herosforgemodel"`
	Maxcharges          int            `db:"maxcharges"`
	Mr                  int            `db:"mr"`
	Nodrop              int            `db:"nodrop"`
	Norent              int            `db:"norent"`
	Pendingloreflag     int            `db:"pendingloreflag"`
	Pr                  int            `db:"pr"`
	Procrate            int            `db:"procrate"`
	Races               int            `db:"races"`
	Range               int            `db:"range"`
	Reclevel            int            `db:"reclevel"`
	Recskill            int            `db:"recskill"`
	Reqlevel            int            `db:"reqlevel"`
	Sellrate            float64        `db:"sellrate"`
	Shielding           int            `db:"shielding"`
	Size                int            `db:"size"`
	Skillmodtype        int            `db:"skillmodtype"`
	Skillmodvalue       int            `db:"skillmodvalue"`
	Slots               int            `db:"slots"`
	Clickeffect         int            `db:"clickeffect"`
	Spellshield         int            `db:"spellshield"`
	Strikethrough       int            `db:"strikethrough"`
	Stunresist          int            `db:"stunresist"`
	Summonedflag        int            `db:"summonedflag"`
	Tradeskills         int            `db:"tradeskills"`
	Favor               int            `db:"favor"`
	Weight              int            `db:"weight"`
	Unk012              int            `db:"UNK012"`
	Unk013              int            `db:"UNK013"`
	Benefitflag         int            `db:"benefitflag"`
	Unk054              int            `db:"UNK054"`
	Unk059              int            `db:"UNK059"`
	Booktype            int            `db:"booktype"`
	Recastdelay         int            `db:"recastdelay"`
	Recasttype          int            `db:"recasttype"`
	Guildfavor          int            `db:"guildfavor"`
	Unk123              int            `db:"UNK123"`
	Unk124              int            `db:"UNK124"`
	Attuneable          int            `db:"attuneable"`
	Nopet               int            `db:"nopet"`
	Updated             sql.NullTime   `db:"updated"`
	Comment             string         `db:"comment"`
	Unk127              int            `db:"UNK127"`
	Pointtype           int            `db:"pointtype"`
	Potionbelt          int            `db:"potionbelt"`
	Potionbeltslots     int            `db:"potionbeltslots"`
	Stacksize           int            `db:"stacksize"`
	Notransfer          int            `db:"notransfer"`
	Stackable           int            `db:"stackable"`
	Unk134              string         `db:"UNK134"`
	Unk137              int            `db:"UNK137"`
	Proceffect          int            `db:"proceffect"`
	Proctype            int            `db:"proctype"`
	Proclevel2          int            `db:"proclevel2"`
	Proclevel           int            `db:"proclevel"`
	Unk142              int            `db:"UNK142"`
	Worneffect          int            `db:"worneffect"`
	Worntype            int            `db:"worntype"`
	Wornlevel2          int            `db:"wornlevel2"`
	Wornlevel           int            `db:"wornlevel"`
	Unk147              int            `db:"UNK147"`
	Focustype           int            `db:"focustype"`
	Focuslevel2         int            `db:"focuslevel2"`
	Focuslevel          int            `db:"focuslevel"`
	Unk152              int            `db:"UNK152"`
	Scrolleffect        int            `db:"scrolleffect"`
	Scrolltype          int            `db:"scrolltype"`
	Scrolllevel2        int            `db:"scrolllevel2"`
	Scrolllevel         int            `db:"scrolllevel"`
	Unk157              int            `db:"UNK157"`
	Serialized          sql.NullTime   `db:"serialized"`
	Verified            sql.NullTime   `db:"verified"`
	Serialization       sql.NullString `db:"serialization"`
	Source              string         `db:"source"`
	Unk033              int            `db:"UNK033"`
	Lorefile            string         `db:"lorefile"`
	Unk014              int            `db:"UNK014"`
	Svcorruption        int            `db:"svcorruption"`
	Skillmodmax         int            `db:"skillmodmax"`
	Unk060              int            `db:"UNK060"`
	Augslot1unk2        int            `db:"augslot1unk2"`
	Augslot2unk2        int            `db:"augslot2unk2"`
	Augslot3unk2        int            `db:"augslot3unk2"`
	Augslot4unk2        int            `db:"augslot4unk2"`
	Augslot5unk2        int            `db:"augslot5unk2"`
	Augslot6unk2        int            `db:"augslot6unk2"`
	Unk120              int            `db:"UNK120"`
	Unk121              int            `db:"UNK121"`
	Questitemflag       int            `db:"questitemflag"`
	Unk132              sql.NullString `db:"UNK132"`
	Clickunk5           int            `db:"clickunk5"`
	Clickunk6           string         `db:"clickunk6"`
	Clickunk7           int            `db:"clickunk7"`
	Procunk1            int            `db:"procunk1"`
	Procunk2            int            `db:"procunk2"`
	Procunk3            int            `db:"procunk3"`
	Procunk4            int            `db:"procunk4"`
	Procunk6            string         `db:"procunk6"`
	Procunk7            int            `db:"procunk7"`
	Wornunk1            int            `db:"wornunk1"`
	Wornunk2            int            `db:"wornunk2"`
	Wornunk3            int            `db:"wornunk3"`
	Wornunk4            int            `db:"wornunk4"`
	Wornunk5            int            `db:"wornunk5"`
	Wornunk6            string         `db:"wornunk6"`
	Wornunk7            int            `db:"wornunk7"`
	Focusunk1           int            `db:"focusunk1"`
	Focusunk2           int            `db:"focusunk2"`
	Focusunk3           int            `db:"focusunk3"`
	Focusunk4           int            `db:"focusunk4"`
	Focusunk5           int            `db:"focusunk5"`
	Focusunk6           string         `db:"focusunk6"`
	Focusunk7           int            `db:"focusunk7"`
	Scrollunk1          int            `db:"scrollunk1"`
	Scrollunk2          int            `db:"scrollunk2"`
	Scrollunk3          int            `db:"scrollunk3"`
	Scrollunk4          int            `db:"scrollunk4"`
	Scrollunk5          int            `db:"scrollunk5"`
	Scrollunk6          string         `db:"scrollunk6"`
	Scrollunk7          int            `db:"scrollunk7"`
	Unk193              int            `db:"UNK193"`
	Purity              int            `db:"purity"`
	Evoitem             int            `db:"evoitem"`
	Evoid               int            `db:"evoid"`
	Evolvinglevel       int            `db:"evolvinglevel"`
	Evomax              int            `db:"evomax"`
	Clickname           string         `db:"clickname"`
	Procname            string         `db:"procname"`
	Wornname            string         `db:"wornname"`
	Focusname           string         `db:"focusname"`
	Scrollname          string         `db:"scrollname"`
	Dsmitigation        int            `db:"dsmitigation"`
	Heroic_str          int            `db:"heroic_str"`
	Heroic_int          int            `db:"heroic_int"`
	Heroic_wis          int            `db:"heroic_wis"`
	Heroic_agi          int            `db:"heroic_agi"`
	Heroic_dex          int            `db:"heroic_dex"`
	Heroic_sta          int            `db:"heroic_sta"`
	Heroic_cha          int            `db:"heroic_cha"`
	Heroic_pr           int            `db:"heroic_pr"`
	Heroic_dr           int            `db:"heroic_dr"`
	Heroic_fr           int            `db:"heroic_fr"`
	Heroic_cr           int            `db:"heroic_cr"`
	Heroic_mr           int            `db:"heroic_mr"`
	Heroic_svcorrup     int            `db:"heroic_svcorrup"`
	Healamt             int            `db:"healamt"`
	Spelldmg            int            `db:"spelldmg"`
	Clairvoyance        int            `db:"clairvoyance"`
	Backstabdmg         int            `db:"backstabdmg"`
	Created             string         `db:"created"`
	Elitematerial       int            `db:"elitematerial"`
	Ldonsellbackrate    int            `db:"ldonsellbackrate"`
	Scriptfileid        int            `db:"scriptfileid"`
	Expendablearrow     int            `db:"expendablearrow"`
	Powersourcecapacity int            `db:"powersourcecapacity"`
	Bardeffect          int            `db:"bardeffect"`
	Bardeffecttype      int            `db:"bardeffecttype"`
	Bardlevel2          int            `db:"bardlevel2"`
	Bardlevel           int            `db:"bardlevel"`
	Bardunk1            int            `db:"bardunk1"`
	Bardunk2            int            `db:"bardunk2"`
	Bardunk3            int            `db:"bardunk3"`
	Bardunk4            int            `db:"bardunk4"`
	Bardunk5            int            `db:"bardunk5"`
	Bardname            string         `db:"bardname"`
	Bardunk7            int            `db:"bardunk7"`
	Unk214              int            `db:"UNK214"`
	Subtype             int            `db:"subtype"`
	Unk220              int            `db:"UNK220"`
	Unk221              int            `db:"UNK221"`
	Heirloom            int            `db:"heirloom"`
	Unk223              int            `db:"UNK223"`
	Unk224              int            `db:"UNK224"`
	Unk225              int            `db:"UNK225"`
	Unk226              int            `db:"UNK226"`
	Unk227              int            `db:"UNK227"`
	Unk228              int            `db:"UNK228"`
	Unk229              int            `db:"UNK229"`
	Unk230              int            `db:"UNK230"`
	Unk231              int            `db:"UNK231"`
	Unk232              int            `db:"UNK232"`
	Unk233              int            `db:"UNK233"`
	Unk234              int            `db:"UNK234"`
	Placeable           int            `db:"placeable"`
	Unk236              int            `db:"UNK236"`
	Unk237              int            `db:"UNK237"`
	Unk238              int            `db:"UNK238"`
	Unk239              int            `db:"UNK239"`
	Unk240              int            `db:"UNK240"`
	Unk241              int            `db:"UNK241"`
	Epicitem            int            `db:"epicitem"`
}

func (t *Item) Identifier() string {
	return "item"
}

func (t *Item) Key() string {
	return t.key
}

func (t *Item) SetKey(key string) {
	t.key = key
}

func (t *Item) SetExpiration(expiration int64) {
	t.expiration = expiration
}

func (t *Item) Expiration() int64 {
	return t.expiration
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
		Bit := 1
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

func (t *Item) AugRestrictType(val int) string {
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

func (t *Item) ProcRateTotal() int {
	return t.Procrate + 100
}
