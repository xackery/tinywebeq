package library

import (
	"encoding/json"
	"strings"
)

type (
	Class   int
	Classes []Class
)

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

func (c Class) String() string {
	return classToString[c]
}

func (c Class) MarshalJSON() ([]byte, error) {
	return json.Marshal(c.String())
}

func (c Classes) String() string {
	str := &strings.Builder{}

	if len(c) == 16 {
		return "ALL"
	}

	for i, class := range c {
		str.WriteString(classToShortString[class])
		if i != len(c)-1 {
			str.WriteString(" ")
		}
	}
	return str.String()
}

func (c Classes) MarshalJSON() ([]byte, error) {
	return []byte(`"` + c.String() + `"`), nil
}

func ClassesFromBitmask(mask int32) Classes {
	var classes Classes
	var i int32

	for i = 1; i <= mask; i <<= 1 {
		if i&mask != 0 {
			classes = append(classes, Class(i))
		}
	}
	return classes
}
