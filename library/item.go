package library

type ItemType int32

const (
	ItemType1HandSlashing ItemType = iota
	ItemType2HandSlashing
	ItemTypePiercing
	ItemType1HandBlunt
	ItemType2HandBlunt
	ItemTypeArchery
	_
	ItemTypeThrowingRange
	ItemTypeShield
	_
	ItemTypeArmor
	ItemTypeGems
	ItemTypeLockPicks
	_
	ItemTypeFood
	ItemTypeDrink
	ItemTypeLight
	ItemTypeCombinable
	ItemTypeBandages
	ItemTypeThrowing
	ItemTypeScroll
	ItemTypePotion
	_
	ItemTypeWindInstrument
	ItemTypeStringedInstrument
	ItemTypeBrassInstrument
	ItemTypePercussionInstrument
	ItemTypeArrow
	_
	ItemTypeJewelry
	ItemTypeSkill
	ItemTypeTome
	ItemTypeNote
	ItemTypeKey
	ItemTypeCoin
	ItemType2HandPiercing
	ItemTypeFishingPole
	ItemTypeFishingBait
	ItemTypeAlcohol
	ItemTypeKeyBIS
	ItemTypeCompass
	_
	ItemTypePoison
	_
	_
	ItemTypeMartial
	_
	_
	_
	_
	_
	_
	_
	ItemTypeCharm
	_
	ItemTypeAugmentation
)

var itemTypeToString = map[ItemType]string{
	ItemType1HandSlashing:        "1H Slashing",
	ItemType2HandSlashing:        "2H Slashing",
	ItemTypePiercing:             "Piercing",
	ItemType1HandBlunt:           "1H Blunt",
	ItemType2HandBlunt:           "2H Blunt",
	ItemTypeArchery:              "Archery",
	ItemTypeThrowingRange:        "Throwing",
	ItemTypeShield:               "Shield",
	ItemTypeArmor:                "Armor",
	ItemTypeGems:                 "Gem",
	ItemTypeLockPicks:            "Lock Picks",
	ItemTypeFood:                 "Food",
	ItemTypeDrink:                "Drink",
	ItemTypeLight:                "Light Source",
	ItemTypeCombinable:           "Combinable",
	ItemTypeBandages:             "Bandages",
	ItemTypeThrowing:             "Throwing",
	ItemTypeScroll:               "Scroll",
	ItemTypePotion:               "Potion",
	ItemTypeWindInstrument:       "Wind Instrument",
	ItemTypeStringedInstrument:   "Stringed Instrument",
	ItemTypeBrassInstrument:      "Brass Instrument",
	ItemTypePercussionInstrument: "Percussion Instrument",
	ItemTypeArrow:                "Arrow",
	ItemTypeJewelry:              "Jewelry",
	ItemTypeSkill:                "Skill",
	ItemTypeTome:                 "Tome",
	ItemTypeNote:                 "Note",
	ItemTypeKey:                  "Key",
	ItemTypeCoin:                 "Coin",
	ItemType2HandPiercing:        "2H Piercing",
	ItemTypeFishingPole:          "Fishing Pole",
	ItemTypeFishingBait:          "Fishing Bait",
	ItemTypeAlcohol:              "Alcohol",
	ItemTypeKeyBIS:               "Key",
	ItemTypeCompass:              "Compass",
	ItemTypePoison:               "Poison",
	ItemTypeMartial:              "Martial",
	ItemTypeCharm:                "Charm",
	ItemTypeAugmentation:         "Augmentation",
}

// IsWeapon returns true if the item is some type of weapon.
func (i ItemType) IsWeapon() bool {
	weaponTypes := []ItemType{
		ItemType1HandSlashing,
		ItemType2HandSlashing,
		ItemTypePiercing,
		ItemType1HandBlunt,
		ItemType2HandBlunt,
		ItemTypeArchery,
		ItemTypeThrowingRange,
		ItemTypeThrowing,
		ItemTypeArrow,
		ItemType2HandPiercing,
	}

	for _, typ := range weaponTypes {
		if i == typ {
			return true
		}
	}

	return false
}

// WeaponSkillName returns a proper skill name for weapon item types.
// If the item is not classified as a weapon type, an empty string is returned.
func (i ItemType) WeaponSkillName() string {
	if i.IsWeapon() {
		if i == ItemTypeArrow {
			return "Archery"
		}

		if i == ItemTypeThrowing || i == ItemTypeThrowingRange {
			return "Throwing"
		}

		return i.String()
	}

	return ""
}

// IsInstrument returns true if the item is some type of instrument.
func (i ItemType) IsInstrument() bool {
	instrumentTypes := []ItemType{
		ItemTypeWindInstrument,
		ItemTypeStringedInstrument,
		ItemTypeBrassInstrument,
		ItemTypePercussionInstrument,
	}

	for _, typ := range instrumentTypes {
		if i == typ {
			return true
		}
	}

	return false
}

func (i ItemType) String() string {
	return itemTypeToString[i]
}
