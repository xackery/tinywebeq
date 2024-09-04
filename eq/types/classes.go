package types

import "strings"

// ClassesBitmask is an integer that represents a set of EQ classes.
type ClassesBitmask int32

const (
	ClassWarrior Class = 1 << iota
	ClassCleric
	ClassPaladin
	ClassRanger
	ClassShadowKnight
	ClassDruid
	ClassMonk
	ClassBard
	ClassRogue
	ClassShaman
	ClassNecromancer
	ClassWizard
	ClassMagician
	ClassEnchanter
	ClassBeastlord
	ClassBerserker
)

// classToString is the full name representation of class names.
var classToString = map[Class]string{
	ClassWarrior:      "Warrior",
	ClassCleric:       "Cleric",
	ClassPaladin:      "Paladin",
	ClassRanger:       "Ranger",
	ClassShadowKnight: "Shadow Knight",
	ClassDruid:        "Druid",
	ClassMonk:         "Monk",
	ClassBard:         "Bard",
	ClassRogue:        "Rogue",
	ClassShaman:       "Shaman",
	ClassNecromancer:  "Necromancer",
	ClassWizard:       "Wizard",
	ClassMagician:     "Magician",
	ClassEnchanter:    "Enchanter",
	ClassBeastlord:    "Beastlord",
	ClassBerserker:    "Berserker",
}

// classToShortString is the short name representation of class names.
var classToShortString = map[Class]string{
	ClassWarrior:      "WAR",
	ClassCleric:       "CLR",
	ClassPaladin:      "PAL",
	ClassRanger:       "RNG",
	ClassShadowKnight: "SHD",
	ClassDruid:        "DRU",
	ClassMonk:         "MNK",
	ClassBard:         "BRD",
	ClassRogue:        "ROG",
	ClassShaman:       "SHM",
	ClassNecromancer:  "NEC",
	ClassWizard:       "WIZ",
	ClassMagician:     "MAG",
	ClassEnchanter:    "ENC",
	ClassBeastlord:    "BST",
	ClassBerserker:    "BER",
}

func NewClassesBitmask(bitmask int32) Class {
	return Class(bitmask)
}

// list is an internally used function to return a list of class names.
func (c ClassesBitmask) list(typ string) []string {
	var (
		classes []string
		i       int32
	)

	for i = 1; i <= int32(c); i <<= 1 {
		if i&int32(c) != 0 {
			switch typ {
			case "long":
				classes = append(classes, classToString[Class(i)])
			default:
				classes = append(classes, classToShortString[Class(i)])
			}
		}
	}

	return classes
}

// String implements the stringer interface and returns a concatenated list of short name classes.
func (c ClassesBitmask) String() string {
	list := c.list("short")

	if len(list) == 16 {
		return "ALL"
	}

	return strings.Join(list, " ")
}

// NamesList returns a list of class names in long format.
func (c ClassesBitmask) NamesList() []string {
	return c.list("long")
}

// ShortNamesList returns a list of class short names.
func (c ClassesBitmask) ShortNamesList() []string {
	return c.list("short")
}

// MarshalJSON is a custom JSON marshaler that returns the concatenated short form of class names.
func (c Class) MarshalJSON() ([]byte, error) {
	return []byte(`"` + c.String() + `"`), nil
}

type Class uint

func (c Class) String() string {
	return classToString[c]
}
