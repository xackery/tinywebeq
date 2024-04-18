package util

import (
	"fmt"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/tlog"
)

func (e *Util) SpellInfo(id int) []string {
	lines := []string{}
	if !config.Get().IsSpellInfoEnabled {
		return lines
	}

	mu.RLock()
	defer mu.RUnlock()
	se, ok := spells[id]
	if !ok {
		return lines
	}

	// based on https://github.com/macroquest/macroquest/blob/b5e5fe972b89642c64589751a7b3e1daab43bc49/src/plugins/itemdisplay/MQ2ItemDisplay.cpp#L881
	lines = append(lines, fmt.Sprintf("Spell Info for Effect: %s", se.Name))
	lines = append(lines, fmt.Sprintf("ID: %d", se.ID))

	duration := spellDuration(se, MaxPCLevel)
	if duration == -1 {
		lines = append(lines, "Duration: Permanent")
	} else if duration == -2 {
		lines = append(lines, "Duration: Unknown")
	} else if duration == -4 {
		lines = append(lines, "Duration: Permanent Aura")
	} else if duration == 0 {
	} else {
		lines = append(lines, fmt.Sprintf("Duration: %0.1f minutes", float64(duration*6/60)))
	}

	if se.RecoveryTime > 0 {
		lines = append(lines, fmt.Sprintf("Recovery Time: %0.2f", float64(se.RecoveryTime)/1000.0))
	}
	if se.RecastTime > 0 {
		lines = append(lines, fmt.Sprintf("Recast Time: %0.2f", float64(se.RecastTime)/1000.0))
	}
	if se.Range > 0 {
		lines = append(lines, fmt.Sprintf("Range: %d", se.Range))
	}
	if se.Pushback > 0 {
		lines = append(lines, fmt.Sprintf("Pushback: %0.1f", float64(se.Pushback)))
	}

	for i := 0; i < 12; i++ {
		line := spellEffect(se, i)
		if line != "" {
			lines = append(lines, line)
		}
	}
	out := ""
	for i := 0; i < 16; i++ {
		class := se.Classes[i]
		if class < 1 || class >= 255 {
			continue
		}

		out += fmt.Sprintf("%s (%d), ", classNames[i+1], class)
	}
	if out != "" {
		out = out[:len(out)-2]
		lines = append(lines, out)
	}

	return lines
}

func spellEffect(se *Spell, index int) string {
	//id := se.ID
	spa := se.Attribs[index]
	base := se.Bases[index]
	base2 := se.Limits[index]
	max := se.Maxes[index]
	calc := se.Calcs[index]
	ticks := se.DurationCap
	targets := se.MaxTargets
	targetType := se.TargetType
	//skill := se.Skill
	spellEffectName := spellEffectName(spa)
	level := 60
	minSpellLvl := 1 // calcMinSpellLevel(se)
	maxSpellLvl := calcMaxSpellLevel(calc, base, max, ticks, minSpellLvl, level)

	tlog.Debugf("spellEffect(%d, %d)", index+1, spa)

	switch spa {
	case SPA_CHA:
		if base <= 1 || base > 255 {
			return ""
		}
	case SPA_NOSPELL, SPA_PORTAL_LOCATIONS:
		return ""
	}
	switch spa {
	case SPA_HASTE, SPA_HEIGHT, SPA_BARD_HASTE:
		base -= 100
		max -= 100
	case SPA_SUMMON_CORPSE:
		max = base
		base = 0
	case SPA_FOCUS_DAMAGE_MOD, SPA_FOCUS_HEAL_MOD, SPA_FOCUS_MANACOST_MOD:
		max = base2
	case SPA_FOCUS_REAGENT_MOD, SPA_FOCUS_DAMAGE_AMT_DETRIMENTAL:
		base = base2
	}

	value := calcValue(calc, base, max, 1, minSpellLvl, minSpellLvl)
	finish := calcValue(calc, base, max, ticks, minSpellLvl, level)

	usePercent := false
	switch spa {
	case SPA_MOVEMENT_RATE, SPA_HASTE, SPA_BARD_HASTE, SPA_FOCUS_DAMAGE_MOD, SPA_FOCUS_HEAL_MOD, SPA_DOUBLE_ATTACK, SPA_STUN_RESIST, SPA_PROC_CHANCE, SPA_DIVINE_SAVE, SPA_METABOLISM, SPA_TRIPLE_BACKSTAB, SPA_DOTCRIT, SPA_HEALCRIT, SPA_MENDCRIT, SPA_FLURRY, SPA_PET_FLURRY, SPA_SPELL_CRIT_CHANCE, SPA_SHIELD_BLOCK_CHANCE, SPA_FOCUS_DAMAGE_MOD_CRIT, SPA_FOCUS_INCOMING_DMG_MOD, SPA_CANCEL_NEGATIVE_MAGIC:
		usePercent = true
	}

	AEEffect := false
	switch targetType {
	case TT_PBAE, TT_TARGETED_AE, TT_AE_PC_V2, TT_DIRECTIONAL:
		AEEffect = true
	}

	spellRange := calcValueRange(calc, base, max, ticks, minSpellLvl, level, usePercent)
	extendedRange := calcExtendedRange(calc, value, finish, minSpellLvl, maxSpellLvl, usePercent, spa == SPA_AC || spa == SPA_AC2)
	repeating := ""
	if ticks != 0 {
		repeating = " per tick "
	}
	maxLevel := ""
	if max != 0 {
		maxLevel = fmt.Sprintf(" up to level %d", max)
	}
	maxTargets := ""
	if targets != 0 && AEEffect {
		maxTargets = fmt.Sprintf(" on up to %d enemies", targets)
	}

	out := fmt.Sprintf("Slot %d: ", index+1)

	tmp := ""
	percent := ""
	if usePercent {
		percent = "%"
	}
	switch spa {
	case SPA_HP:
		tmp = "Increase"
		if value < 0 {
			tmp = "Decrease"
		}
		out += fmt.Sprintf("%s %s %s", tmp, spellEffectName, extendedRange) // formatRange
		if repeating != "" {
			out += repeating
		}
		if maxTargets != "" {
			out += maxTargets
		}
		if base2 != 0 {
			out += fmt.Sprintf(" -- Restrictions: %s", spellRestrictions[base2])
		}
	case SPA_AC, SPA_ATTACK_POWER, SPA_MOVEMENT_RATE, SPA_STR, SPA_DEX, SPA_AGI, SPA_STA, SPA_INT, SPA_WIS, SPA_CHA, SPA_HASTE:
		tmp = "Increase"
		if value < 0 {
			tmp = "Decrease"
		}
		out += fmt.Sprintf("%s %s %s", tmp, spellEffectName, extendedRange) // formatRange
	case SPA_INVISIBILITY:
		out += spellEffectName
	case SPA_SEE_INVIS, SPA_ENDURING_BREATH:
		out += fmt.Sprintf("%s (%d)", spellEffectName, base)
	case SPA_MANA:
		out += fmt.Sprintf("%s %s%d%s%s", tmp, spellEffectName, value, extendedRange, spellRange)
		if repeating != "" {
			out += repeating
		}
		if maxTargets != "" {
			out += maxTargets
		}
		if base2 != 0 {
			out += fmt.Sprintf(" -- Restrictions: %s", spellRestrictions[base2])
		}
	case SPA_NPC_FRENZY, SPA_NPC_AWARENESS, SPA_NPC_AGGRO:
		out += spellEffectName
	case SPA_NPC_FACTION:
		out += fmt.Sprintf("%s %s%d%s%s", tmp, spellEffectName, value, extendedRange, spellRange)
	case SPA_BLINDNESS:
		out += spellEffectName
	case SPA_STUN:
		if base2 != 0 && base != base2 {
			out += fmt.Sprintf(" NPC for %1.fs (PC for %1.fs)%s", float64(base/1000.0), float64(base2/1000.0), maxLevel)
		} else {
			out += fmt.Sprintf(" for %1.fs%s", float64(base/1000.0), maxLevel)
		}
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_CHARM, SPA_FEAR:
		out += fmt.Sprintf("%s %s%s", spellEffectName, maxLevel, spellRange) //FormatString
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_FATIGUE:
		tmp = "Increase"
		if value < 0 {
			tmp = "Decrease"
		}

		out += fmt.Sprintf("%s %s by %d%s", tmp, spellEffectName, abs(value), percent) //FormatCount
	case SPA_BIND_AFFINITY:
		if base == 2 {
			tmp = "(Secondary Bind Point)"
		}
		if base == 3 {
			tmp = "(Tertiary Bind Point)"
		}
		out += fmt.Sprintf("%s %s%s", spellEffectName, tmp, spellRange) //FormatString
	case SPA_GATE:

		if base == 2 {
			tmp = " to Secondary Bind Point"
		}
		if base == 3 {
			tmp = " (Tertiary Bind Point)"
		}
		out += fmt.Sprintf("%s %s%s", spellEffectName, tmp, spellRange) //FormatString
	case SPA_DISPEL_MAGIC:
		out += fmt.Sprintf("%s (%d)", spellEffectName, base)
	case SPA_INVIS_VS_UNDEAD, SPA_INVIS_VS_ANIMALS:
		out += spellEffectName
	case SPA_NPC_AGGRO_RADIUS:
		tmp = "Increase"
		if value < 0 {
			tmp = "Decrease"
		}

		out += fmt.Sprintf("%s %s by %d%s", tmp, spellEffectName, -value, percent) //FormatCount
	case SPA_ENTHRALL:
		out += fmt.Sprintf("%s %s%s", spellEffectName, maxLevel, maxTargets) //FormatString
	case SPA_CREATE_ITEM:
		out += fmt.Sprintf("%s %d (Qty: %d)", spellEffectName, base, calc) //FormatString
	default:
		out += fmt.Sprintf("%s (id=%d, base=%d, base2=%d, max=%d, calc=%d, value=%d)", spellEffectName, spa, base, base2, max, calc, value)
	}

	tlog.Debugf("%s (base=%d, base2=%d, max=%d, calc=%d, value=%d)", spellEffectName, base, base2, max, calc, value)

	return out
}
