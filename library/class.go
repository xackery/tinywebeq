package library

import (
	"fmt"
	"strings"
)

func ClassStr(in int) string {
	switch in {
	case 1:
		return "Warrior"
	case 2:
		return "Cleric"
	case 3:
		return "Paladin"
	case 4:
		return "Ranger"
	case 5:
		return "Shadow Knight"
	case 6:
		return "Druid"
	case 7:
		return "Monk"
	case 8:
		return "Bard"
	case 9:
		return "Rogue"
	case 10:
		return "Shaman"
	case 11:
		return "Necromancer"
	case 12:
		return "Wizard"
	case 13:
		return "Magician"
	case 14:
		return "Enchanter"
	case 15:
		return "Beastlord"
	case 16:
		return "Berserker"
	}
	return fmt.Sprintf("Unknown %d", in)
}

// ClassesFromMask returns a string of classes from a bitmask
func ClassesFromMask(in int) string {
	out := ""

	if in == 65535 {
		return "ALL"
	}
	if in&1 != 0 {
		out += "WAR "
	}
	if in&2 != 0 {
		out += "CLR "
	}
	if in&4 != 0 {
		out += "PAL "
	}
	if in&8 != 0 {
		out += "RNG "
	}
	if in&16 != 0 {
		out += "SHD "
	}
	if in&32 != 0 {
		out += "DRU "
	}
	if in&64 != 0 {
		out += "MNK "
	}
	if in&128 != 0 {
		out += "BRD "
	}
	if in&256 != 0 {
		out += "ROG "
	}
	if in&512 != 0 {
		out += "SHM "
	}
	if in&1024 != 0 {
		out += "NEC "
	}
	if in&2048 != 0 {
		out += "WIZ "
	}
	if in&4096 != 0 {
		out += "MAG "
	}
	if in&8192 != 0 {
		out += "ENC "
	}
	if in&16384 != 0 {
		out += "BST "
	}

	out = strings.TrimSuffix(out, " ")
	return out
}
