package library

import (
	"fmt"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/tlog"
)

func SpellInfo(id int) []string {
	lines := []string{}
	if !config.Get().Spell.IsSpellInfoEnabled {
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

	if se.Mana > 0 {
		lines = append(lines, fmt.Sprintf("Mana: %d", se.Mana))
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
	spellEffectName := spellEffectNameByID(spa)
	level := 60
	minSpellLvl := 1 // calcMinSpellLevel(se)
	maxSpellLvl := calcMaxSpellLevel(calc, base, max, ticks, minSpellLvl, level)

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
		maxLevel = fmt.Sprintf("up to level %d", max)
	}
	maxTargets := ""
	if targets != 0 && AEEffect {
		maxTargets = fmt.Sprintf("on up to %d enemies", targets)
	}

	out := fmt.Sprintf("Slot %d: ", index+1)
	tlog.Debugf("%s (base=%d, base2=%d, max=%d, calc=%d, value=%d)", spellEffectName, base, base2, max, calc, value)

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
		out += formatRange(spellEffectName, value, extendedRange, spellRange)
	case SPA_INVISIBILITY:
		out += spellEffectName
	case SPA_SEE_INVIS, SPA_ENDURING_BREATH:
		out += formatBase(spellEffectName, value, value, "")
	case SPA_MANA:
		out += formatRange(spellEffectName, value, extendedRange, spellRange)
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
		out += formatRange(spellEffectName, value, extendedRange, spellRange)
	case SPA_BLINDNESS:
		out += spellEffectName
	case SPA_STUN:
		tmp = fmt.Sprintf("for %1.fs%s", float64(base/1000.0), maxLevel)
		if base2 != 0 && base != base2 {
			tmp = fmt.Sprintf("NPC for %1.fs (PC for %1.fs)%s", float64(base/1000.0), float64(base2/1000.0), maxLevel)
		}
		out += formatString(spellEffectName, tmp, spellRange)
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_CHARM, SPA_FEAR:
		out += formatString(spellEffectName, maxLevel, spellRange)
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_FATIGUE:
		out += formatCount(spellEffectName, value, spellRange, "")
	case SPA_BIND_AFFINITY:
		if base == 2 {
			tmp = "(Secondary Bind Point)"
		}
		if base == 3 {
			tmp = "(Tertiary Bind Point)"
		}
		out += formatString(spellEffectName, tmp, spellRange)
	case SPA_GATE:
		if base == 2 {
			tmp = "to Secondary Bind Point"
		}
		if base == 3 {
			tmp = "(Tertiary Bind Point)"
		}
		out += formatString(spellEffectName, tmp, spellRange)
	case SPA_DISPEL_MAGIC:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_INVIS_VS_UNDEAD, SPA_INVIS_VS_ANIMALS:
		out += spellEffectName
	case SPA_NPC_AGGRO_RADIUS:
		out += formatCount(spellEffectName, -value, spellRange, "")
		out += maxLevel
	case SPA_ENTHRALL:
		out += formatString(spellEffectName, maxLevel, spellRange)
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_CREATE_ITEM:
		out += fmt.Sprintf("%s %d (Qty: %d)", spellEffectName, base, calc)
	case SPA_SUMMON_PET:
		out += formatExtra(spellEffectName, se.TeleportZone, spellRange, "")
	case SPA_CONFUSE:
		out += spellEffectName
	case SPA_DISEASE, SPA_POISON:
		out += formatRange(spellEffectName, value, extendedRange, spellRange)
	case SPA_DETECT_HOSTILE, SPA_DETECT_MAGIC, SPA_NO_TWINCAST, SPA_INVULNERABILITY, SPA_BANISH, SPA_SHADOW_STEP, SPA_BERSERK, SPA_LYCANTHROPY, SPA_VAMPIRISM:
		out += spellEffectName
	case SPA_RESIST_FIRE, SPA_RESIST_COLD, SPA_RESIST_POISON, SPA_RESIST_DISEASE, SPA_RESIST_MAGIC:
		out += formatRange(spellEffectName, value, extendedRange, spellRange)
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_DETECT_TRAPS, SPA_DETECT_UNDEAD, SPA_DETECT_SUMMONED, SPA_DETECT_ANIMALS:
		out += spellEffectName
	case SPA_STONESKIN:
		out += formatRange(spellEffectName, value, extendedRange, spellRange)
	case SPA_TRUE_NORTH:
		out += spellEffectName
	case SPA_LEVITATION:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_CHANGE_FORM:
		out += formatExtra(spellEffectName, fmt.Sprintf("Race ID %d", base), spellRange, "")
	case SPA_DAMAGE_SHIELD:
		out += formatRange(spellEffectName, -value, extendedRange, spellRange)
	case SPA_TRANSFER_ITEM, SPA_ITEM_LORE, SPA_ITEM_IDENTIFY:
		out += spellEffectName
	case SPA_NPC_WIPE_HATE_LIST:
		out += formatPenaltyChance(spellEffectName, value+40, "Chance")
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_SPIN_STUN, SPA_INFRAVISION, SPA_ULTRAVISION, SPA_EYE_OF_ZOMM, SPA_RECLAIM_ENERGY:
		out += spellEffectName
	case SPA_MAX_HP:
		out += formatRange(spellEffectName, value, extendedRange, spellRange)
	case SPA_CORPSE_BOMB:
		out += spellEffectName
	case SPA_CREATE_UNDEAD:
		out += formatExtra(spellEffectName, se.TeleportZone, spellRange, "")
	case SPA_PRESERVE_CORPSE, SPA_BIND_SIGHT:
		out += spellEffectName
	case SPA_FEIGN_DEATH:
		out += formatPenaltyChance(spellEffectName, value, "Chance")
	case SPA_VENTRILOQUISM, SPA_SENTINEL, SPA_LOCATE_CORPSE:
		out += spellEffectName
	case SPA_SPELL_SHIELD, SPA_INSTANT_HP:
		out += formatRange(spellEffectName, value, extendedRange, spellRange)
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_ENCHANT_LIGHT:
		out += spellEffectName
	case SPA_RESURRECT:
		tmp = fmt.Sprintf("and restore %d%% experience", value)
		out += formatString(spellEffectName, tmp, spellRange)
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_SUMMON_TARGET:
		out += spellEffectName
	case SPA_PORTAL:
		context := "Self"
		if targetType != 6 {
			context = "Group"
		}
		out += fmt.Sprintf("%s %s to %d, %d, %d in %s", spellEffectName, context, se.Bases[0], se.Bases[1], se.Bases[2], ZoneLongNameByShortName(se.TeleportZone))
	case SPA_HP_NPC_ONLY:
		tmp := "Up"
		if base < 0 {
			tmp = "Down"
		}
		out += formatBase(spellEffectName, abs(base), abs(base), tmp)
	case SPA_MELEE_PROC:
		out += formatExtra(spellEffectName, formatRateMod(spellEffectName, base, base2), spellRange, "")
	case SPA_NPC_HELP_RADIUS:
		out += formatCount(spellEffectName, -value, spellRange, "")
		out += maxLevel
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_MAGNIFICATION:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_EVACUATE:
		out += fmt.Sprintf("%s to %d, %d, %d in %s", spellEffectName, se.Bases[0], se.Bases[1], se.Bases[2], ZoneLongNameByShortName(se.TeleportZone))
	case SPA_HEIGHT:
		out += fmt.Sprintf("%s %d%s", spellEffectName, value, percent)
	case SPA_IGNORE_PET, SPA_SUMMON_CORPSE:
		out += formatString(spellEffectName, maxLevel, spellRange)
	case SPA_HATE:
		out += formatRange(spellEffectName, value, extendedRange, spellRange)
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_WEATHER_CONTROL, SPA_FRAGILE, SPA_SACRIFICE:
		out += spellEffectName
	case SPA_SILENCE:
		out += spellEffectName
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_MAX_MANA:
	case SPA_BARD_HASTE:
		out += formatRange(spellEffectName, value, extendedRange, spellRange)
	case SPA_ROOT:
		out += spellEffectName
	case SPA_HEALDOT:
		out += formatRange(spellEffectName, value, extendedRange, spellRange)
		out += repeating
	case SPA_COMPLETEHEAL:
	case SPA_PET_FEARLESS:
	case SPA_CALL_PET:
		out += spellEffectName
	case SPA_TRANSLOCATE:
		if se.TeleportZone == "" {
			out += fmt.Sprintf("%s to Bind Point", spellEffectName)
		} else {
			out += fmt.Sprintf("%s to %d, %d, %d in %s", spellEffectName, se.Bases[0], se.Bases[1], se.Bases[2], ZoneLongNameByShortName(se.TeleportZone))
		}
	case SPA_NPC_ANTI_GATE:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_BEASTLORD_PET:
		out += formatExtra(spellEffectName, se.TeleportZone, spellRange, "")
	case SPA_ALTER_PET_LEVEL:
		out += spellEffectName
	case SPA_FAMILIAR:
		out += formatExtra(spellEffectName, se.TeleportZone, spellRange, "")
	case SPA_CREATE_ITEM_IN_BAG:
		out += fmt.Sprintf("%s %d", spellEffectName, base)
	case SPA_ARCHERY:
		out += formatCount(spellEffectName, value, spellRange, "")
	case SPA_RESIST_ALL:
		out += formatRange(spellEffectName, value, extendedRange, spellRange)
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_FIZZLE_SKILL:
		out += formatCount(spellEffectName, value, spellRange, "")
	case SPA_SUMMON_MOUNT:
		out += formatExtra(spellEffectName, se.TeleportZone, spellRange, "")
	case SPA_MODIFY_HATE:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_CORNUCOPIA:
		out += spellEffectName
	case SPA_CURSE:
		out += formatCount(spellEffectName, value, spellRange, "")
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_HIT_MAGIC:
		out += spellEffectName
	case SPA_AMPLIFICATION:
		out += formatRange(spellEffectName, value, extendedRange, spellRange)
	case SPA_ATTACK_SPEED_MAX:
	case SPA_HEALMOD:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_IRONMAIDEN:
		out += formatBase(spellEffectName, -base, -base, spellRange)
	case SPA_REDUCESKILL:
		out += formatSkills(spellEffectName, value, finish, base2, spellRange, "")
	case SPA_IMMUNITY:
		out += spellEffectName
	case SPA_FOCUS_DAMAGE_MOD:
	case SPA_FOCUS_HEAL_MOD:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_FOCUS_RESIST_MOD:
	case SPA_FOCUS_CAST_TIME_MOD:
		out += formatPercent(spellEffectName, -value, -finish, false, false, spellRange)
	case SPA_FOCUS_DURATION_MOD:
	case SPA_FOCUS_RANGE_MOD:
	case SPA_FOCUS_HATE_MOD:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_FOCUS_REAGENT_MOD:
	case SPA_FOCUS_MANACOST_MOD:
	case SPA_FOCUS_STUNTIME_MOD:
		out += formatPercent(spellEffectName, -value, -finish, false, false, spellRange)

	default:
		out += fmt.Sprintf("%s (id=%d, base=%d, base2=%d, max=%d, calc=%d, value=%d)", spellEffectName, spa, base, base2, max, calc, value)
		return out
	}

	out += fmt.Sprintf(" (spa %d)", spa)
	return out
}
