package db

import (
	"database/sql"
	"fmt"
	"strings"
)

type Item struct {
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
	return "Item"
}

func (t *Item) ClassStr() string {
	return ClassesFromMask(t.Classes)
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
	switch t.Banedmgrace {
	case 0:
		return "Invalid"
	case 1:
		return "Human"
	case 2:
		return "Barbarian"
	case 3:
		return "Erudite"
	case 4:
		return "Wood Elf"
	case 5:
		return "High Elf"
	case 6:
		return "Dark Elf"
	case 7:
		return "Half Elf"
	case 8:
		return "Dwarf"
	case 9:
		return "Troll"
	case 10:
		return "Ogre"
	case 11:
		return "Halfling"
	case 12:
		return "Gnome"
	case 13:
		return "Aviak"
	case 14:
		return "Werewolf"
	case 15:
		return "Brownie"
	case 16:
		return "Centaur"
	case 17:
		return "Golem"
	case 18:
		return "Giant"
	case 19:
		return "Trakanon"
	case 20:
		return "Venril Sathir"
	case 21:
		return "Evil Eye"
	case 22:
		return "Beetle"
	case 23:
		return "Kerran"
	case 24:
		return "Fish"
	case 25:
		return "Fairy"
	case 26:
		return "Froglok"
	case 27:
		return "Froglok"
	case 28:
		return "Fungusman"
	case 29:
		return "Gargoyle"
	case 30:
		return "Gasbag"
	case 31:
		return "Gelatinous Cube"
	case 32:
		return "Ghost"
	case 33:
		return "Ghoul"
	case 34:
		return "Bat"
	case 35:
		return "Eel"
	case 36:
		return "Rat"
	case 37:
		return "Snake"
	case 38:
		return "Spider"
	case 39:
		return "Gnoll"
	case 40:
		return "Goblin"
	case 41:
		return "Gorilla"
	case 42:
		return "Wolf"
	case 43:
		return "Bear"
	case 44:
		return "Guard"
	case 45:
		return "Demi Lich"
	case 46:
		return "Imp"
	case 47:
		return "Griffin"
	case 48:
		return "Kobold"
	case 49:
		return "Dragon"
	case 50:
		return "Lion"
	case 51:
		return "Lizard Man"
	case 52:
		return "Mimic"
	case 53:
		return "Minotaur"
	case 54:
		return "Orc"
	case 55:
		return "Beggar"
	case 56:
		return "Pixie"
	case 57:
		return "Drachnid"
	case 58:
		return "Solusek Ro"
	case 59:
		return "Goblin"
	case 60:
		return "Skeleton"
	case 61:
		return "Shark"
	case 62:
		return "Tunare"
	case 63:
		return "Tiger"
	case 64:
		return "Treant"
	case 65:
		return "Vampire"
	case 66:
		return "Rallos Zek"
	case 67:
		return "Human"
	case 68:
		return "Tentacle Terror"
	case 69:
		return "Will-O-Wisp"
	case 70:
		return "Zombie"
	case 71:
		return "Human"
	case 72:
		return "Ship"
	case 73:
		return "Launch"
	case 74:
		return "Piranha"
	case 75:
		return "Elemental"
	case 76:
		return "Puma"
	case 77:
		return "Dark Elf"
	case 78:
		return "Erudite"
	case 79:
		return "Bixie"
	case 80:
		return "Reanimated Hand"
	case 81:
		return "Halfling"
	case 82:
		return "Scarecrow"
	case 83:
		return "Skunk"
	case 84:
		return "Snake Elemental"
	case 85:
		return "Spectre"
	case 86:
		return "Sphinx"
	case 87:
		return "Armadillo"
	case 88:
		return "Clockwork Gnome"
	case 89:
		return "Drake"
	case 90:
		return "Barbarian"
	case 91:
		return "Alligator"
	case 92:
		return "Troll"
	case 93:
		return "Ogre"
	case 94:
		return "Dwarf"
	case 95:
		return "Cazic Thule"
	case 96:
		return "Cockatrice"
	case 97:
		return "Daisy Man"
	case 98:
		return "Vampire"
	case 99:
		return "Amygdalan"
	case 100:
		return "Dervish"
	case 101:
		return "Efreeti"
	case 102:
		return "Tadpole"
	case 103:
		return "Kedge"
	case 104:
		return "Leech"
	case 105:
		return "Swordfish"
	case 106:
		return "Guard"
	case 107:
		return "Mammoth"
	case 108:
		return "Eye"
	case 109:
		return "Wasp"
	case 110:
		return "Mermaid"
	case 111:
		return "Harpy"
	case 112:
		return "Guard"
	case 113:
		return "Drixie"
	case 114:
		return "Ghost Ship"
	case 115:
		return "Clam"
	case 116:
		return "Seahorse"
	case 117:
		return "Ghost"
	case 118:
		return "Ghost"
	case 119:
		return "Sabertooth"
	case 120:
		return "Wolf"
	case 121:
		return "Gorgon"
	case 122:
		return "Dragon"
	case 123:
		return "Innoruuk"
	case 124:
		return "Unicorn"
	case 125:
		return "Pegasus"
	case 126:
		return "Djinn"
	case 127:
		return "Invisible Man"
	case 128:
		return "Iksar"
	case 129:
		return "Scorpion"
	case 130:
		return "Vah Shir"
	case 131:
		return "Sarnak"
	case 132:
		return "Draglock"
	case 133:
		return "Drolvarg"
	case 134:
		return "Mosquito"
	case 135:
		return "Rhinoceros"
	case 136:
		return "Xalgoz"
	case 137:
		return "Goblin"
	case 138:
		return "Yeti"
	case 139:
		return "Iksar"
	case 140:
		return "Giant"
	case 141:
		return "Boat"
	case 142:
		return "Uknown"
	case 143:
		return "Uknown"
	case 144:
		return "Burynai"
	case 145:
		return "Goo"
	case 146:
		return "Sarnak Spirit"
	case 147:
		return "Iksar Spirit"
	case 148:
		return "Fish"
	case 149:
		return "Scorpion"
	case 150:
		return "Erollisi"
	case 151:
		return "Tribunal"
	case 152:
		return "Bertoxxulous"
	case 153:
		return "Bristlebane"
	case 154:
		return "Fay Drake"
	case 155:
		return "Undead Sarnak"
	case 156:
		return "Ratman"
	case 157:
		return "Wyvern"
	case 158:
		return "Wurm"
	case 159:
		return "Devourer"
	case 160:
		return "Iksar Golem"
	case 161:
		return "Undead Iksar"
	case 162:
		return "Man-Eating Plant"
	case 163:
		return "Raptor"
	case 164:
		return "Sarnak Golem"
	case 165:
		return "Dragon"
	case 166:
		return "Animated Hand"
	case 167:
		return "Succulent"
	case 168:
		return "Holgresh"
	case 169:
		return "Brontotherium"
	case 170:
		return "Snow Dervish"
	case 171:
		return "Dire Wolf"
	case 172:
		return "Manticore"
	case 173:
		return "Totem"
	case 174:
		return "Ice Spectre"
	case 175:
		return "Enchanted Armor"
	case 176:
		return "Snow Rabbit"
	case 177:
		return "Walrus"
	case 178:
		return "Geonid"
	case 179:
		return "Uknown"
	case 180:
		return "Uknown"
	case 181:
		return "Yakkar"
	case 182:
		return "Faun"
	case 183:
		return "Coldain"
	case 184:
		return "Dragon"
	case 185:
		return "Hag"
	case 186:
		return "Hippogriff"
	case 187:
		return "Siren"
	case 188:
		return "Giant"
	case 189:
		return "Giant"
	case 190:
		return "Othmir"
	case 191:
		return "Ulthork"
	case 192:
		return "Dragon"
	case 193:
		return "Abhorrent"
	case 194:
		return "Sea Turtle"
	case 195:
		return "Dragon"
	case 196:
		return "Dragon"
	case 197:
		return "Ronnie Test"
	case 198:
		return "Dragon"
	case 199:
		return "Shik'Nar"
	case 200:
		return "Rockhopper"
	case 201:
		return "Underbulk"
	case 202:
		return "Grimling"
	case 203:
		return "Worm"
	case 204:
		return "Evan Test"
	case 205:
		return "Shadel"
	case 206:
		return "Owlbear"
	case 207:
		return "Rhino Beetle"
	case 208:
		return "Vampire"
	case 209:
		return "Earth Elemental"
	case 210:
		return "Air Elemental"
	case 211:
		return "Water Elemental"
	case 212:
		return "Fire Elemental"
	case 213:
		return "Wetfang Minnow"
	case 214:
		return "Thought Horror"
	case 215:
		return "Tegi"
	case 216:
		return "Horse"
	case 217:
		return "Shissar"
	case 218:
		return "Fungal Fiend"
	case 219:
		return "Vampire"
	case 220:
		return "Stonegrabber"
	case 221:
		return "Scarlet Cheetah"
	case 222:
		return "Zelniak"
	case 223:
		return "Lightcrawler"
	case 224:
		return "Shade"
	case 225:
		return "Sunflower"
	case 226:
		return "Sun Revenant"
	case 227:
		return "Shrieker"
	case 228:
		return "Galorian"
	case 229:
		return "Netherbian"
	case 230:
		return "Akheva"
	case 231:
		return "Grieg Veneficus"
	case 232:
		return "Sonic Wolf"
	case 233:
		return "Ground Shaker"
	case 234:
		return "Vah Shir Skeleton"
	case 235:
		return "Wretch"
	case 236:
		return "Seru"
	case 237:
		return "Recuso"
	case 238:
		return "Vah Shir"
	case 239:
		return "Guard"
	case 240:
		return "Teleport Man"
	case 241:
		return "Werewolf"
	case 242:
		return "Nymph"
	case 243:
		return "Dryad"
	case 244:
		return "Treant"
	case 245:
		return "Fly"
	case 246:
		return "Tarew Marr"
	case 247:
		return "Solusek Ro"
	case 248:
		return "Clockwork Golem"
	case 249:
		return "Clockwork Brain"
	case 250:
		return "Banshee"
	case 251:
		return "Guard of Justice"
	case 252:
		return "Mini POM"
	case 253:
		return "Diseased Fiend"
	case 254:
		return "Solusek Ro Guard"
	case 255:
		return "Bertoxxulous"
	case 256:
		return "The Tribunal"
	case 257:
		return "Terris Thule"
	case 258:
		return "Vegerog"
	case 259:
		return "Crocodile"
	case 260:
		return "Bat"
	case 261:
		return "Hraquis"
	case 262:
		return "Tranquilion"
	case 263:
		return "Tin Soldier"
	case 264:
		return "Nightmare Wraith"
	case 265:
		return "Malarian"
	case 266:
		return "Knight of Pestilence"
	case 267:
		return "Lepertoloth"
	case 268:
		return "Bubonian"
	case 269:
		return "Bubonian Underling"
	case 270:
		return "Pusling"
	case 271:
		return "Water Mephit"
	case 272:
		return "Stormrider"
	case 273:
		return "Junk Beast"
	case 274:
		return "Broken Clockwork"
	case 275:
		return "Giant Clockwork"
	case 276:
		return "Clockwork Beetle"
	case 277:
		return "Nightmare Goblin"
	case 278:
		return "Karana"
	case 279:
		return "Blood Raven"
	case 280:
		return "Nightmare Gargoyle"
	case 281:
		return "Mouth of Insanity"
	case 282:
		return "Skeletal Horse"
	case 283:
		return "Saryrn"
	case 284:
		return "Fennin Ro"
	case 285:
		return "Tormentor"
	case 286:
		return "Soul Devourer"
	case 287:
		return "Nightmare"
	case 288:
		return "Rallos Zek"
	case 289:
		return "Vallon Zek"
	case 290:
		return "Tallon Zek"
	case 291:
		return "Air Mephit"
	case 292:
		return "Earth Mephit"
	case 293:
		return "Fire Mephit"
	case 294:
		return "Nightmare Mephit"
	case 295:
		return "Zebuxoruk"
	case 296:
		return "Mithaniel Marr"
	case 297:
		return "Undead Knight"
	case 298:
		return "The Rathe"
	case 299:
		return "Xegony"
	case 300:
		return "Fiend"
	case 301:
		return "Test Object"
	case 302:
		return "Crab"
	case 303:
		return "Phoenix"
	case 304:
		return "Dragon"
	case 305:
		return "Bear"
	case 306:
		return "Giant"
	case 307:
		return "Giant"
	case 308:
		return "Giant"
	case 309:
		return "Giant"
	case 310:
		return "Giant"
	case 311:
		return "Giant"
	case 312:
		return "Giant"
	case 313:
		return "War Wraith"
	case 314:
		return "Wrulon"
	case 315:
		return "Kraken"
	case 316:
		return "Poison Frog"
	case 317:
		return "Nilborien"
	case 318:
		return "Valorian"
	case 319:
		return "War Boar"
	case 320:
		return "Efreeti"
	case 321:
		return "War Boar"
	case 322:
		return "Valorian"
	case 323:
		return "Animated Armor"
	case 324:
		return "Undead Footman"
	case 325:
		return "Rallos Zek Minion"
	case 326:
		return "Arachnid"
	case 327:
		return "Crystal Spider"
	case 328:
		return "Zebuxoruk's Cage"
	case 329:
		return "BoT Portal"
	case 330:
		return "Froglok"
	case 331:
		return "Troll"
	case 332:
		return "Troll"
	case 333:
		return "Troll"
	case 334:
		return "Ghost"
	case 335:
		return "Pirate"
	case 336:
		return "Pirate"
	case 337:
		return "Pirate"
	case 338:
		return "Pirate"
	case 339:
		return "Pirate"
	case 340:
		return "Pirate"
	case 341:
		return "Pirate"
	case 342:
		return "Pirate"
	case 343:
		return "Frog"
	case 344:
		return "Troll Zombie"
	case 345:
		return "Luggald"
	case 346:
		return "Luggald"
	case 347:
		return "Luggalds"
	case 348:
		return "Drogmore"
	case 349:
		return "Froglok Skeleton"
	case 350:
		return "Undead Froglok"
	case 351:
		return "Knight of Hate"
	case 352:
		return "Arcanist of Hate"
	case 353:
		return "Veksar"
	case 354:
		return "Veksar"
	case 355:
		return "Veksar"
	case 356:
		return "Chokadai"
	case 357:
		return "Undead Chokadai"
	case 358:
		return "Undead Veksar"
	case 359:
		return "Vampire"
	case 360:
		return "Vampire"
	case 361:
		return "Rujarkian Orc"
	case 362:
		return "Bone Golem"
	case 363:
		return "Synarcana"
	case 364:
		return "Sand Elf"
	case 365:
		return "Vampire"
	case 366:
		return "Rujarkian Orc"
	case 367:
		return "Skeleton"
	case 368:
		return "Mummy"
	case 369:
		return "Goblin"
	case 370:
		return "Insect"
	case 371:
		return "Froglok Ghost"
	case 372:
		return "Dervish"
	case 373:
		return "Shade"
	case 374:
		return "Golem"
	case 375:
		return "Evil Eye"
	case 376:
		return "Box"
	case 377:
		return "Barrel"
	case 378:
		return "Chest"
	case 379:
		return "Vase"
	case 380:
		return "Table"
	case 381:
		return "Weapon Rack"
	case 382:
		return "Coffin"
	case 383:
		return "Bones"
	case 384:
		return "Jokester"
	case 385:
		return "Nihil"
	case 386:
		return "Trusik"
	case 387:
		return "Stone Worker"
	case 388:
		return "Hynid"
	case 389:
		return "Turepta"
	case 390:
		return "Cragbeast"
	case 391:
		return "Stonemite"
	case 392:
		return "Ukun"
	case 393:
		return "Ixt"
	case 394:
		return "Ikaav"
	case 395:
		return "Aneuk"
	case 396:
		return "Kyv"
	case 397:
		return "Noc"
	case 398:
		return "Ra`tuk"
	case 399:
		return "Taneth"
	case 400:
		return "Huvul"
	case 401:
		return "Mutna"
	case 402:
		return "Mastruq"
	case 403:
		return "Taelosian"
	case 404:
		return "Discord Ship"
	case 405:
		return "Stone Worker"
	case 406:
		return "Mata Muram"
	case 407:
		return "Lightning Warrior"
	case 408:
		return "Succubus"
	case 409:
		return "Bazu"
	case 410:
		return "Feran"
	case 411:
		return "Pyrilen"
	case 412:
		return "Chimera"
	case 413:
		return "Dragorn"
	case 414:
		return "Murkglider"
	case 415:
		return "Rat"
	case 416:
		return "Bat"
	case 417:
		return "Gelidran"
	case 418:
		return "Discordling"
	case 419:
		return "Girplan"
	case 420:
		return "Minotaur"
	case 421:
		return "Dragorn Box"
	case 422:
		return "Runed Orb"
	case 423:
		return "Dragon Bones"
	case 424:
		return "Muramite Armor Pile"
	case 425:
		return "Crystal Shard"
	case 426:
		return "Portal"
	case 427:
		return "Coin Purse"
	case 428:
		return "Rock Pile"
	case 429:
		return "Murkglider Egg Sack"
	case 430:
		return "Drake"
	case 431:
		return "Dervish"
	case 432:
		return "Drake"
	case 433:
		return "Goblin"
	case 434:
		return "Kirin"
	case 435:
		return "Dragon"
	case 436:
		return "Basilisk"
	case 437:
		return "Dragon"
	case 438:
		return "Dragon"
	case 439:
		return "Puma"
	case 440:
		return "Spider"
	case 441:
		return "Spider Queen"
	case 442:
		return "Animated Statue"
	case 443:
		return "Uknown"
	case 444:
		return "Uknown"
	case 445:
		return "Dragon Egg"
	case 446:
		return "Dragon Statue"
	case 447:
		return "Lava Rock"
	case 448:
		return "Animated Statue"
	case 449:
		return "Spider Egg Sack"
	case 450:
		return "Lava Spider"
	case 451:
		return "Lava Spider Queen"
	case 452:
		return "Dragon"
	case 453:
		return "Giant"
	case 454:
		return "Werewolf"
	case 455:
		return "Kobold"
	case 456:
		return "Sporali"
	case 457:
		return "Gnomework"
	case 458:
		return "Orc"
	case 459:
		return "Corathus"
	case 460:
		return "Coral"
	case 461:
		return "Drachnid"
	case 462:
		return "Drachnid Cocoon"
	case 463:
		return "Fungus Patch"
	case 464:
		return "Gargoyle"
	case 465:
		return "Witheran"
	case 466:
		return "Dark Lord"
	case 467:
		return "Shiliskin"
	case 468:
		return "Snake"
	case 469:
		return "Evil Eye"
	case 470:
		return "Minotaur"
	case 471:
		return "Zombie"
	case 472:
		return "Clockwork Boar"
	case 473:
		return "Fairy"
	case 474:
		return "Witheran"
	case 475:
		return "Air Elemental"
	case 476:
		return "Earth Elemental"
	case 477:
		return "Fire Elemental"
	case 478:
		return "Water Elemental"
	case 479:
		return "Alligator"
	case 480:
		return "Bear"
	case 481:
		return "Scaled Wolf"
	case 482:
		return "Wolf"
	case 483:
		return "Spirit Wolf"
	case 484:
		return "Skeleton"
	case 485:
		return "Spectre"
	case 486:
		return "Bolvirk"
	case 487:
		return "Banshee"
	case 488:
		return "Banshee"
	case 489:
		return "Elddar"
	case 490:
		return "Forest Giant"
	case 491:
		return "Bone Golem"
	case 492:
		return "Horse"
	case 493:
		return "Pegasus"
	case 494:
		return "Shambling Mound"
	case 495:
		return "Scrykin"
	case 496:
		return "Treant"
	case 497:
		return "Vampire"
	case 498:
		return "Ayonae Ro"
	case 499:
		return "Sullon Zek"
	case 500:
		return "Banner"
	case 501:
		return "Flag"
	case 502:
		return "Rowboat"
	case 503:
		return "Bear Trap"
	case 504:
		return "Clockwork Bomb"
	case 505:
		return "Dynamite Keg"
	case 506:
		return "Pressure Plate"
	case 507:
		return "Puffer Spore"
	case 508:
		return "Stone Ring"
	case 509:
		return "Root Tentacle"
	case 510:
		return "Runic Symbol"
	case 511:
		return "Saltpetter Bomb"
	case 512:
		return "Floating Skull"
	case 513:
		return "Spike Trap"
	case 514:
		return "Totem"
	case 515:
		return "Web"
	case 516:
		return "Wicker Basket"
	case 517:
		return "Nightmare/Unicorn"
	case 518:
		return "Horse"
	case 519:
		return "Nightmare/Unicorn"
	case 520:
		return "Bixie"
	case 521:
		return "Centaur"
	case 522:
		return "Drakkin"
	case 523:
		return "Giant"
	case 524:
		return "Gnoll"
	case 525:
		return "Griffin"
	case 526:
		return "Giant Shade"
	case 527:
		return "Harpy"
	case 528:
		return "Mammoth"
	case 529:
		return "Satyr"
	case 530:
		return "Dragon"
	case 531:
		return "Dragon"
	case 532:
		return "Dyn'Leth"
	case 533:
		return "Boat"
	case 534:
		return "Weapon Rack"
	case 535:
		return "Armor Rack"
	case 536:
		return "Honey Pot"
	case 537:
		return "Jum Jum Bucket"
	case 538:
		return "Toolbox"
	case 539:
		return "Stone Jug"
	case 540:
		return "Small Plant"
	case 541:
		return "Medium Plant"
	case 542:
		return "Tall Plant"
	case 543:
		return "Wine Cask"
	case 544:
		return "Elven Boat"
	case 545:
		return "Gnomish Boat"
	case 546:
		return "Barrel Barge Ship"
	case 547:
		return "Goo"
	case 548:
		return "Goo"
	case 549:
		return "Goo"
	case 550:
		return "Merchant Ship"
	case 551:
		return "Pirate Ship"
	case 552:
		return "Ghost Ship"
	case 553:
		return "Banner"
	case 554:
		return "Banner"
	case 555:
		return "Banner"
	case 556:
		return "Banner"
	case 557:
		return "Banner"
	case 558:
		return "Aviak"
	case 559:
		return "Beetle"
	case 560:
		return "Gorilla"
	case 561:
		return "Kedge"
	case 562:
		return "Kerran"
	case 563:
		return "Shissar"
	case 564:
		return "Siren"
	case 565:
		return "Sphinx"
	case 566:
		return "Human"
	case 567:
		return "Campfire"
	case 568:
		return "Brownie"
	case 569:
		return "Dragon"
	case 570:
		return "Exoskeleton"
	case 571:
		return "Ghoul"
	case 572:
		return "Clockwork Guardian"
	case 573:
		return "Mantrap"
	case 574:
		return "Minotaur"
	case 575:
		return "Scarecrow"
	case 576:
		return "Shade"
	case 577:
		return "Rotocopter"
	case 578:
		return "Tentacle Terror"
	case 579:
		return "Wereorc"
	case 580:
		return "Worg"
	case 581:
		return "Wyvern"
	case 582:
		return "Chimera"
	case 583:
		return "Kirin"
	case 584:
		return "Puma"
	case 585:
		return "Boulder"
	case 586:
		return "Banner"
	case 587:
		return "Elven Ghost"
	case 588:
		return "Human Ghost"
	case 589:
		return "Chest"
	case 590:
		return "Chest"
	case 591:
		return "Crystal"
	case 592:
		return "Coffin"
	case 593:
		return "Guardian CPU"
	case 594:
		return "Worg"
	case 595:
		return "Mansion"
	case 596:
		return "Floating Island"
	case 597:
		return "Cragslither"
	case 598:
		return "Wrulon"
	case 599:
		return "Spell Particle 1"
	case 600:
		return "Invisible Man of Zomm"
	case 601:
		return "Robocopter of Zomm"
	case 602:
		return "Burynai"
	case 603:
		return "Frog"
	case 604:
		return "Dracolich"
	case 605:
		return "Iksar Ghost"
	case 606:
		return "Iksar Skeleton"
	case 607:
		return "Mephit"
	case 608:
		return "Muddite"
	case 609:
		return "Raptor"
	case 610:
		return "Sarnak"
	case 611:
		return "Scorpion"
	case 612:
		return "Tsetsian"
	case 613:
		return "Wurm"
	case 614:
		return "Balrog"
	case 615:
		return "Hydra Crystal"
	case 616:
		return "Crystal Sphere"
	case 617:
		return "Gnoll"
	case 618:
		return "Sokokar"
	case 619:
		return "Stone Pylon"
	case 620:
		return "Demon Vulture"
	case 621:
		return "Wagon"
	case 622:
		return "God of Discord"
	case 623:
		return "Wrulon Mount"
	case 624:
		return "Ogre NPC - Male"
	case 625:
		return "Sokokar Mount"
	case 626:
		return "Giant (Rallosian mats)"
	case 627:
		return "Sokokar (w saddle)"
	case 628:
		return "10th Anniversary Banner"
	case 629:
		return "10th Anniversary Cake"
	case 630:
		return "Wine Cask"
	case 631:
		return "Hydra Mount"
	case 632:
		return "Hydra NPC"
	case 633:
		return "Wedding Flowers"
	case 634:
		return "Wedding Arbor"
	case 635:
		return "Wedding Altar"
	case 636:
		return "Powder Keg"
	case 637:
		return "Apexus"
	case 638:
		return "Bellikos"
	case 639:
		return "Brell's First Creation"
	case 640:
		return "Brell"
	case 641:
		return "Crystalskin Ambuloid"
	case 642:
		return "Cliknar Queen"
	case 643:
		return "Cliknar Soldier"
	case 644:
		return "Cliknar Worker"
	case 645:
		return "Coldain"
	case 646:
		return "Coldain"
	case 647:
		return "Crystalskin Sessiloid"
	case 648:
		return "Genari"
	case 649:
		return "Gigyn"
	case 650:
		return "Greken - Young Adult"
	case 651:
		return "Greken - Young"
	case 652:
		return "Cliknar Mount"
	case 653:
		return "Telmira"
	case 654:
		return "Spider Mount"
	case 655:
		return "Bear Mount"
	case 656:
		return "Rat Mount"
	case 657:
		return "Sessiloid Mount"
	case 658:
		return "Morell Thule"
	case 659:
		return "Marionette"
	case 660:
		return "Book Dervish"
	case 661:
		return "Topiary Lion"
	case 662:
		return "Rotdog"
	case 663:
		return "Amygdalan"
	case 664:
		return "Sandman"
	case 665:
		return "Grandfather Clock"
	case 666:
		return "Gingerbread Man"
	case 667:
		return "Beefeater"
	case 668:
		return "Rabbit"
	case 669:
		return "Blind Dreamer"
	case 670:
		return "Cazic Thule"
	case 671:
		return "Topiary Lion Mount"
	case 672:
		return "Rot Dog Mount"
	case 673:
		return "Goral Mount"
	case 674:
		return "Selyran Mount"
	case 675:
		return "Sclera Mount"
	case 676:
		return "Braxy Mount"
	case 677:
		return "Kangon Mount"
	case 678:
		return "Erudite"
	case 679:
		return "Wurm Mount"
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

func (t *Item) ItemTypeStr() string {
	switch t.Itemtype {
	case 0:
		return "1HS"
	case 1:
		return "2HS"
	case 2:
		return "Piercing"
	case 3:
		return "1HB"
	case 4:
		return "2HB"
	case 5:
		return "Archery"
	case 6:
		return "Unknown"
	case 7:
		return "Throwing range items"
	case 8:
		return "Shield"
	case 9:
		return "Unknown"
	case 10:
		return "Armor"
	case 11:
		return "Gems"
	case 12:
		return "Lockpicks"
	case 13:
		return "Unknown"
	case 14:
		return "Food"
	case 15:
		return "Drink"
	case 16:
		return "Light"
	case 17:
		return "Combinable"
	case 18:
		return "Bandages"
	case 19:
		return "Throwing"
	case 20:
		return "Scroll"
	case 21:
		return "Potion"
	case 22:
		return "Unknown"
	case 23:
		return "Wind Instrument"
	case 24:
		return "Stringed Instrument"
	case 25:
		return "Brass Instrument"
	case 26:
		return "Percussion Instrument"
	case 27:
		return "Arrow"
	case 28:
		return "Unknown"
	case 29:
		return "Jewelry"
	case 30:
		return "Skull"
	case 31:
		return "Tome"
	case 32:
		return "Note"
	case 33:
		return "Key"
	case 34:
		return "Coin"
	case 35:
		return "2H Piercing"
	case 36:
		return "Fishing Pole"
	case 37:
		return "Fishing Bait"
	case 38:
		return "Alcohol"
	case 39:
		return "Key (bis)"
	case 40:
		return "Compass"
	case 41:
		return "Unknown"
	case 42:
		return "Poison"
	case 43:
		return "Unknown"
	case 44:
		return "Unknown"
	case 45:
		return "Martial"
	case 46:
		return "Unknown"
	case 47:
		return "Unknown"
	case 48:
		return "Unknown"
	case 49:
		return "Unknown"
	case 50:
		return "Unknown"
	case 51:
		return "Unknown"
	case 52:
		return "Charm"
	case 53:
		return "Unknown"
	case 54:
		return "Augmentation"
	}
	return "Unknown"
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
	return SkillName(t.Extradmgskill)
}

func (t *Item) SkillModTypeStr() string {
	return SkillName(t.Skillmodtype)
}
