package util

import (
	"fmt"
	"math"
)

func spellEffectName(id int) string {
	absID := int(math.Abs(float64(id)))
	if absID >= len(szSPATypes) {
		return fmt.Sprintf("Unknown SPA[%03d]", absID)
	}
	return szSPATypes[absID]
}

func calcValue(calc int, base int, max int, tick int, minLevel int, level int) int {
	if calc == 0 {
		return base
	}

	if calc == 100 {
		if max > 0 && ((base > max) || (level > minLevel)) {
			return max
		}
		return base
	}

	var change int
	var adjustment int

	switch calc {
	case 100:
	case 101:
		change = level / 2
	case 102:
		change = level
	case 103:
		change = level * 2
	case 104:
		change = level * 3
	case 105:
		change = level * 4
	case 106:
		change = level * 5
	case 107:
		change = -1 * tick
	case 108:
		change = -2 * tick
	case 109:
		change = level / 4
	case 110:
		change = level / 6
	case 111:
		if level > 16 {
			change = (level - 16) * 6
		}
	case 112:
		if level > 24 {
			change = (level - 24) * 8
		}
	case 113:
		if level > 34 {
			change = (level - 34) * 10
		}
	case 114:
		if level > 44 {
			change = (level - 44) * 15
		}
	case 115:
		if level > 15 {
			change = (level - 15) * 7
		}
	case 116:
		if level > 24 {
			change = (level - 24) * 10
		}
	case 117:
		if level > 34 {
			change = (level - 34) * 13
		}
	case 118:
		if level > 44 {
			change = (level - 44) * 20
		}
	case 119:
		change = level / 8
	case 120:
		change = -5 * tick
	case 121:
		change = level / 3
	case 122:
		change = -12 * tick
	case 123:
		if tick > 1 {
			change = abs(max) - abs(base)
		}
	case 124:
		if level > 50 {
			change = (level - 50)
		}
	case 125:
		if level > 50 {
			change = (level - 50) * 2
		}
	case 126:
		if level > 50 {
			change = (level - 50) * 3
		}
	case 127:
		if level > 50 {
			change = (level - 50) * 4
		}
	case 128:
		if level > 50 {
			change = (level - 50) * 5
		}
	case 129:
		if level > 50 {
			change = (level - 50) * 10
		}
	case 130:
		if level > 50 {
			change = (level - 50) * 15
		}
	case 131:
		if level > 50 {
			change = (level - 50) * 20
		}
	case 132:
		if level > 50 {

			change = (level - 50) * 25
		}
	case 139:
		if level > 30 {
			change = (level - 30) / 2
		}
	case 140:
		if level > 30 {
			change = (level - 30)
		}
	case 141:
		if level > 30 {
			change = 3 * (level - 30) / 2
		}
	case 142:
		if level > 30 {
			change = 2 * (level - 30)
		}
	case 143:
		change = 3 * level / 4
	case 3000:
		return base
	default:
		if calc > 0 && calc < 1000 {
			change = level * calc
		}
		if calc >= 1000 && calc < 2000 {
			change = tick * (calc - 1000) * -1
		}
		if calc >= 2000 {
			change = level * (calc - 2000)
		}
	}

	value := abs(base) + adjustment + change

	if max != 0 && value > abs(max) {
		value = abs(max)
	}

	if base < 0 {
		value = -value
	}

	return value
}

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func calcValueRange(calc int, base int, max int, duration int, minLevel int, level int, usePercent bool) string {
	start := calcValue(calc, base, max, 1, minLevel, minLevel)
	finish := calcValue(calc, base, max, duration, minLevel, level)
	var t string
	if abs(start) < abs(finish) {
		t = "Growing"
	} else {
		t = "Decaying"
	}

	switch calc {
	case SpellValueRangeCalc_DecayTick1:
		return fmt.Sprintf(" (%s to %d @ 1/tick)", t, finish)
	case SpellValueRangeCalc_DecayTick2:
		return fmt.Sprintf(" (%s to %d @ 2/tick)", t, finish)
	case SpellValueRangeCalc_DecayTick5:
		return fmt.Sprintf(" (%s to %d @ 5/tick)", t, finish)
	case SpellValueRangeCalc_DecayTick12:
		return fmt.Sprintf(" (%s to %d @ 12/tick)", t, finish)
	case SpellValueRangeCalc_Random:
		if start < 0 {
			return fmt.Sprintf(" (Random: %d to %d)", start, finish*-1)
		}
		return fmt.Sprintf(" (Random: %d to %d)", start, finish)
	default:
		if calc > 0 && calc < 1000 {
			percent := ""
			if usePercent {
				percent = "%"
			}
			return fmt.Sprintf(" to %d%s", start, percent)
		}
		if calc >= 1000 && calc < 2000 {
			return fmt.Sprintf(" (%s to %d @ %d/tick)", t, finish, calc-1000)
		}
	}
	return ""
}

func calcExtendedRange(calc int, start int, finish int, minLevel int, maxLevel int, usePercent bool, useACMod bool) string {
	percent := ""
	if usePercent {
		percent = "%"
	}
	switch calc {
	case SpellValueRangeCalc_Random:
		if start < 0 {
			finish = finish * -1
		}
		return fmt.Sprintf(" (Random: %d to %d)", start, finish)
	default:
		startVal := start
		if useACMod {
			startVal = int(float64(start) / float64(10/3))
		}
		finishVal := finish
		if useACMod {
			finishVal = int(float64(finish) / float64(10/3))
		}

		if abs(start) < abs(finish) {
			return fmt.Sprintf(" by %d%s (L%d) to %d%s (L%d)", abs(startVal), percent, minLevel, abs(finishVal), percent, maxLevel)
		}
		return fmt.Sprintf(" by %d%s", abs(finishVal), percent)
	}
}

func calcMaxSpellLevel(calc int, base int, max int, tick int, minLevel int, level int) int {
	if abs(max) < 1 {
		return MaxPCLevel
	}
	for maxLevel := 1; maxLevel < level; maxLevel++ {
		val := calcValue(calc, base, max, tick, minLevel, maxLevel)
		if abs(val) >= abs(max) {
			return maxLevel
		}
	}
	return level
}

func spellDuration(se *Spell, level int) int {
	val := spellDurationCalc(se.DurationCalc, level)
	if val > 0 && val < se.DurationCap {
		return se.DurationCap
	}
	return val
}

func spellDurationCalc(calc int, level int) int {
	switch calc {
	case 1:
		if level > 3 {
			return level / 2
		}
		return 1
	case 2:
		if level > 3 {
			return level/2 + 5
		}
		return 6
	case 3:
		return 30 * level
	case 4:
		return 50
	case 5:
		return 2
	case 6:
		if level > 1 {
			return level/2 + 2
		}
		return 1
	case 7:
		return level
	case 8:
		return level + 10
	case 9:
		return 2*level + 10
	case 10:
		return 3*level + 10
	case 11:
		return 30 * (level + 3)
	case 12:
		if level > 7 {
			return level / 4
		}
		return 1
	case 13:
		return 4*level + 10
	case 14:
		return 5 * (level + 2)
	case 15:
		return 10 * (level + 10)
	case 50:
		return -1
	case 51:
		return -4
	default:
		if calc < 200 {
			return 0
		}
	}
	return calc
}
