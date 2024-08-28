package library

import "fmt"

func ItemTypeStr(in int32) string {
	switch in {
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
		return fmt.Sprintf("Unknown %d", in)
	case 7:
		return "Throwing range items"
	case 8:
		return "Shield"
	case 9:
		return fmt.Sprintf("Unknown %d", in)
	case 10:
		return "Armor"
	case 11:
		return "Gems"
	case 12:
		return "Lockpicks"
	case 13:
		return fmt.Sprintf("Unknown %d", in)
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
		return fmt.Sprintf("Unknown %d", in)
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
		return fmt.Sprintf("Unknown %d", in)
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
		return fmt.Sprintf("Unknown %d", in)
	case 42:
		return "Poison"
	case 43:
		return fmt.Sprintf("Unknown %d", in)
	case 44:
		return fmt.Sprintf("Unknown %d", in)
	case 45:
		return "Martial"
	case 46:
		return fmt.Sprintf("Unknown %d", in)
	case 47:
		return fmt.Sprintf("Unknown %d", in)
	case 48:
		return fmt.Sprintf("Unknown %d", in)
	case 49:
		return fmt.Sprintf("Unknown %d", in)
	case 50:
		return fmt.Sprintf("Unknown %d", in)
	case 51:
		return fmt.Sprintf("Unknown %d", in)
	case 52:
		return "Charm"
	case 53:
		return fmt.Sprintf("Unknown %d", in)
	case 54:
		return "Augmentation"
	}
	return fmt.Sprintf("Unknown %d", in)
}
