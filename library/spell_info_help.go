package library

import (
	"fmt"
	"math"

	"github.com/xackery/tinywebeq/db"
)

func spellEffectNameByID(id int) string {
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
		return fmt.Sprintf("(Random: %d to %d)", start, finish)
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
			return fmt.Sprintf("by %d%s (L%d) to %d%s (L%d)", abs(startVal), percent, minLevel, abs(finishVal), percent, maxLevel)
		}
		return fmt.Sprintf("by %d%s", abs(finishVal), percent)
	}
}

func calcMaxSpellLevel(calc int, base int, max int, tick int, minLevel int, level int) int {
	if abs(max) < 1 {
		return level
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

func formatPercent(spellEffectName string, value int, max int, scaling bool, hundreds bool, percent string) string {
	out := ""
	tmp := "Increase"
	if max < 0 {
		tmp = "Decrease"
	}
	if hundreds {
		if value == max {
			if scaling {
				out = fmt.Sprintf("%s %s by %.2f%s", tmp, spellEffectName, math.Abs(float64(max/100)), percent)
			} else {
				out = fmt.Sprintf("%s by %.2f%s", spellEffectName, math.Abs(float64(max/100)), percent)
			}
		} else {
			if scaling {
				out = fmt.Sprintf("%s %s by %.2f%s to %.2f%s", tmp, spellEffectName, math.Abs(float64(value/100)), percent, math.Abs(float64(max/100)), percent)
			} else {
				out = fmt.Sprintf("%s by %.2f%s to %.2f%s", spellEffectName, math.Abs(float64(value/100)), percent, math.Abs(float64(max/100)), percent)
			}
		}
	} else {
		if value == max {
			if scaling {
				out = fmt.Sprintf("%s %s by %d%s", tmp, spellEffectName, abs(max), percent)
			} else {
				out = fmt.Sprintf("%s by %d%s", spellEffectName, abs(max), percent)
			}
		} else {
			if scaling {
				out = fmt.Sprintf("%s %s by %d%s to %d%s", tmp, spellEffectName, abs(value), percent, abs(max), percent)
			} else {
				out = fmt.Sprintf("%s by %d%s to %d%s", spellEffectName, abs(value), percent, abs(max), percent)
			}
		}
	}
	return out
}

func formatRange(spellEffectName string, value int, spellRange string, extra string) string {
	tmp := "Increase"
	if value < 0 {
		tmp = "Decrease"
	}

	return fmt.Sprintf("%s %s %s%s", tmp, spellEffectName, spellRange, extra)
}

func formatString(spellEffectName string, extra string, trigger string) string {
	return fmt.Sprintf("%s %s%s", spellEffectName, extra, trigger)
}

func formatCount(spellEffectName string, value int, preposition string, percent string) string {
	if preposition == "" {
		preposition = "by"
	}
	tmp := "Increase"
	if value < 0 {
		tmp = "Decrease"
	}
	return fmt.Sprintf("%s %s %s %d%s", tmp, spellEffectName, preposition, abs(value), percent)
}

func formatAT(spellEffectName string, value int, preposition string, percent string) string {
	if preposition == "" {
		preposition = "by"
	}
	return fmt.Sprintf("%s %s %d%s", spellEffectName, preposition, abs(value), percent)
}
func formatBase(spellEffectName string, base int, max int, optional string) string {
	if max == 0 {
		if optional != "" {
			return fmt.Sprintf("%s %s (%d)", spellEffectName, optional, base)
		}
		return fmt.Sprintf("%s (%d)", spellEffectName, base)
	}
	if optional != "" {
		return fmt.Sprintf("%s %s (%d to %d)", spellEffectName, optional, base, max)
	}

	return fmt.Sprintf("%s (%d,%d)", spellEffectName, base, max)
}

func formatBasePercent(spellEffectName string, base int) string {
	return fmt.Sprintf("%s (%d%%)", spellEffectName, base)
}

func formatMinMaxBase(spellEffectName string, base int, spa int) string {
	return fmt.Sprintf("%s (%d %s)", spellEffectName, abs(base), szSPATypes[spa])
}

func formatExtra(spellEffectName string, extra string, trigger string, colon string) string {
	if colon == "" {
		colon = ":"
	}
	return fmt.Sprintf("%s%s %s%s", spellEffectName, colon, extra, trigger)
}

func formatLimits(spellEffectName string, value int, extra string) string {
	tmp := "allowed"
	if value < 0 {
		tmp = "excluded"
	}
	return fmt.Sprintf("%s (%s %s)", spellEffectName, extra, tmp)
}

func formatMax(spellEffectName string, value int, max int) string {
	tmp := "Increase"
	if value < 0 {
		tmp = "Decrease"
	}
	return fmt.Sprintf("%s %s by %d (%d%% max)", tmp, spellEffectName, abs(max), value)
}

func formatPenaltyChance(spellEffectName string, value int, penaltychance string) string {
	if value < 100 {
		return fmt.Sprintf("%s (%d%% %s)", spellEffectName, value, penaltychance)
	}
	return spellEffectName
}

func formatRateMod(spellEffectName string, value int, base int) string {
	if base > 0 {
		return fmt.Sprintf("%s (rate mod %d)", SpellName(value), base)
	}
	return SpellName(value)
}

func formatRefreshTimer(spellEffectName string, value int, max int, skill int, preposition string) string {
	tmp := "Increase"
	if max < 0 {
		tmp = "Decrease"
	}
	tmpSkill := "All Skills"
	if skill >= 0 {
		tmpSkill = db.SkillName(skill)
	}
	if preposition == "" {
		preposition = "with"
	}

	if value == max {
		return fmt.Sprintf("%s %s by %d sec %s %s", tmp, spellEffectName, abs(max), preposition, tmpSkill)
	}
	return fmt.Sprintf("%s %s by %d sec to %d sec %s %s", tmp, spellEffectName, abs(value), abs(max), preposition, tmpSkill)
}

func formatResists(spellEffectName string, value int, base int) string {
	if value >= 100 {
		return spellEffectName
	}

	return fmt.Sprintf("%s (%d%% Chance)", spellEffectNameByID(base), value)
}

func formatSeconds(spellEffectName string, value int, tens bool) string {
	if tens {
		return fmt.Sprintf("%s (%d0.00 sec)", spellEffectName, value)
	}
	return fmt.Sprintf("%s (%d sec)", spellEffectName, value)
}

func formatSecondsCount(spellEffectName string, value float64, preposition string) string {
	tmp := "Increase"
	if value < 0 {
		tmp = "Decrease"
	}

	if preposition == "" {
		preposition = "by"
	}

	return fmt.Sprintf("%s %s %s %.2f sec", tmp, spellEffectName, preposition, math.Abs(value))
}

func formatSkillAttack(spellEffectName string, value int, max int, base2 int, skill int, preposition string) string {
	if preposition == "" {
		preposition = "with"
	}
	tmpSkill := "All Skills"
	if skill >= 0 {
		tmpSkill = db.SkillName(skill)
	}
	return fmt.Sprintf("%s %s %s for %d damage", formatPercent(spellEffectName, value, max, false, false, "%"), preposition, tmpSkill, base2)
}

func formatSkills(spellEffectName string, value int, max int, skill int, percent string, preposition string) string {
	if preposition == "" {
		preposition = "with"
	}
	tmpSkill := "All Skills"
	if skill >= 0 {
		tmpSkill = db.SkillName(skill)
	}
	return fmt.Sprintf("%s %s %s", formatPercent(spellEffectName, value, max, true, false, percent), preposition, tmpSkill)
}

func formatSpellChance(spellEffectName string, value int, base int) string {
	if value > 100 {
		return fmt.Sprintf(" (Spell: %s)", SpellName(base))
	}
	return fmt.Sprintf(" (%d%% Chance, Spell: %s)", value, SpellName(base))
}

func formatSpellGroupChance(spellEffectName string, value int, groupId int) string {
	if value > 100 {
		return fmt.Sprintf(" (Spell: %d)", groupId) //GetSpellNameBySpellGroupID(groupId))
	}

	return fmt.Sprintf(" (%d%% Chance, Spell: %d)", value, groupId) // GetSpellNameBySpellGroupID(groupId))
}

func formatStacking(spellEffectName string, slot int, value int, max int, spa int, extra string) string {
	tmp := "new"
	if spa == SPA_STACKING_BLOCK {
		tmp = "existing"
	}

	if max > 0 {
		return fmt.Sprintf("%s %s spell if slot %d is effect '%s' and < %d", spellEffectName, tmp, slot, extra, value)
	}
	return fmt.Sprintf("%s %s spell if slot %d is effect '%s'", spellEffectName, tmp, slot, extra)
}

func formatStatsCapRange(spellEffectName string, value int, stat string, spellRange string) string {
	tmp := "Increase"
	if value < 0 {
		tmp = "Decrease"
	}
	return fmt.Sprintf("%s %s %s%s", tmp, spellEffectName, stat, spellRange)
}

func formatTimer(spellEffectName string, value float64) string {
	return fmt.Sprintf("%s by %.2f sec", spellEffectName, value)
}

func ResistTypeName(in int) string {
	switch in {
	case 0:
		return "Magic"
	case 1:
		return "Fire"
	case 2:
		return "Cold"
	case 3:
		return "Poison"
	case 4:
		return "Disease"
	case 5:
		return "Chromatic"
	case 6:
		return "Prismatic"
	case 7:
		return "Physical"
	case 8:
		return "Corruption"
	}
	return fmt.Sprintf("Unknown Resist Type %d", in)
}

func TargetTypeLimitsName(in int) string {
	switch in {
	case 1:
		return "Line of Sight"
	case 2:
		return "AE PC v1"
	case 3:
		return "Group v1"
	case 4:
		return "PB AE"
	case 5:
		return "Single"
	case 6:
		return "Self"
	case 8:
		return "Targeted AE"
	case 9:
		return "Animal"
	case 10:
		return "Undead"
	case 11:
		return "Summoned"
	case 13:
		return "LifeTap"
	case 14:
		return "Pet"
	case 15:
		return "Corpse"
	case 16:
		return "Plant"
	case 17:
		return "Uber Giants"
	case 18:
		return "Uber Dragons"
	case 20:
		return "Targeted AE Tap"
	case 24:
		return "AE Undead"
	case 25:
		return "AE Summoned"
	case 32:
		return "Hatelist"
	case 33:
		return "Hatelist2"
	case 34:
		return "Chest"
	case 35:
		return "Special Muramites"
	case 36:
		return "Caster PB PC"
	case 37:
		return "Caster PB NPC"
	case 38:
		return "Pet2"
	case 39:
		return "No Pets"
	case 40:
		return "AE PC v2"
	case 41:
		return "Group v2"
	case 42:
		return "Directional AE"
	case 43:
		return "Single in Group"
	case 44:
		return "Beam"
	case 45:
		return "Free Target"
	case 46:
		return "Target of Target"
	case 47:
		return "Pet Owner"
	case 52:
		return "Single Friendly (or Target's Target)"
	case 50:
		return "Target AE No Players Pets"
	}
	return fmt.Sprintf("Unknown Target Type %d", in)
}

func StatShortName(in int) string {
	switch in {
	case 0:
		return "STR"
	case 1:
		return "STA"
	case 2:
		return "AGI"
	case 3:
		return "DEX"
	case 4:
		return "WIS"
	case 5:

		return "INT"
	case 6:
		return "CHA"
	case 7:
		return "MR"
	case 8:
		return "CR"
	case 9:
		return "FR"
	case 10:
		return "PR"
	case 11:
		return "DR"
	}

	return fmt.Sprintf("Unknown Stat %d", in)
}
