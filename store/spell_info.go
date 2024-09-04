package store

import (
	"fmt"

	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/eq/types"
	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/models"
	"github.com/xackery/tinywebeq/tlog"
)

func SpellInfo(id int32, level int32) (int32, []string) {
	lines := []string{}
	if !config.Get().Spell.IsSpellInfoEnabled {
		return 0, lines
	}

	se := SpellByID(id)
	if se == nil {
		tlog.Warnf("SpellInfo: spell not found: %d", id)
		return 0, lines
	}

	// based on https://github.com/macroquest/macroquest/blob/b5e5fe972b89642c64589751a7b3e1daab43bc49/src/plugins/itemdisplay/MQ2ItemDisplay.cpp#L881
	lines = append(lines, fmt.Sprintf("Spell Info for Effect: %s", se.Name))
	lines = append(lines, fmt.Sprintf("ID: %d", se.ID))

	maxLevel := int32(60)
	if config.Get().MaxLevel > 0 {
		maxLevel = config.Get().MaxLevel
	}
	if level != 0 {
		maxLevel = level
	}

	duration := spellDuration(se, maxLevel)
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

	return se.SpellIcon, lines
}

func spellEffect(se *models.Spell, index int) string {
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
	level := int32(60)
	if config.Get().MaxLevel > 0 {
		level = config.Get().MaxLevel
	}
	minSpellLvl := int32(1) // calcMinSpellLevel(se)
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
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_INVISIBILITY:
		out += spellEffectName
	case SPA_SEE_INVIS, SPA_ENDURING_BREATH:
		out += formatBase(spellEffectName, value, value, "")
	case SPA_MANA:
		out += formatRange(spellEffectName, value, extendedRange, "")
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
		out += formatRange(spellEffectName, value, extendedRange, "")
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
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_DETECT_HOSTILE, SPA_DETECT_MAGIC, SPA_NO_TWINCAST, SPA_INVULNERABILITY, SPA_BANISH, SPA_SHADOW_STEP, SPA_BERSERK, SPA_LYCANTHROPY, SPA_VAMPIRISM:
		out += spellEffectName
	case SPA_RESIST_FIRE, SPA_RESIST_COLD, SPA_RESIST_POISON, SPA_RESIST_DISEASE, SPA_RESIST_MAGIC:
		out += formatRange(spellEffectName, value, extendedRange, "")
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_DETECT_TRAPS, SPA_DETECT_UNDEAD, SPA_DETECT_SUMMONED, SPA_DETECT_ANIMALS:
		out += spellEffectName
	case SPA_STONESKIN:
		out += formatRange(spellEffectName, value, extendedRange, "")
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
		out += formatRange(spellEffectName, value, extendedRange, "")
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
		out += formatRange(spellEffectName, value, extendedRange, "")
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_ENCHANT_LIGHT:
		out += spellEffectName
	case SPA_RESURRECT:
		tmp = fmt.Sprintf("and restore %d%% experience", value)
		out += formatString(spellEffectName, tmp, "")
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
		out += formatRange(spellEffectName, value, extendedRange, "")
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
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_ROOT:
		out += spellEffectName
	case SPA_HEALDOT:
		out += formatRange(spellEffectName, value, extendedRange, "")
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
		out += formatRange(spellEffectName, value, extendedRange, "")
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
		out += formatRange(spellEffectName, value, extendedRange, "")
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
	case SPA_FOCUS_LEVEL_MAX:
		if base2 > 0 {
			out += fmt.Sprintf("%s (%d) (lose %d%% per level over cap)", spellEffectName, base, base2)
		} else {
			out += formatBase(spellEffectName, base, base, spellRange)
		}
	case SPA_FOCUS_RESIST_TYPE:
		out += formatLimits(spellEffectName, value, ResistTypeName(base))
		out += spellRange
	case SPA_FOCUS_TARGET_TYPE:
		out += formatLimits(spellEffectName, value, TargetTypeLimitsName(base))
		out += spellRange
	case SPA_FOCUS_WHICH_SPA:
		out += formatLimits(spellEffectName, value, spellEffectNameByID(base))
		out += spellRange
	case SPA_FOCUS_BENEFICIAL:
		switch base {
		case SpellType_Detrimental:
			tmp = "Detrimental Only"
		case SpellType_Beneficial:
			tmp = "Beneficial Only"
		case SpellType_BeneficialGroupOnly:
			tmp = "Beneficial Group Only"
		}
		out += formatLimits(spellEffectName, value, tmp)
		out += spellRange
	case SPA_FOCUS_WHICH_SPELL:
		out += formatLimits(spellEffectName, value, spellEffectNameByID(base))
		out += spellRange
	case SPA_FOCUS_DURATION_MIN:
		out += formatSeconds(spellEffectName, value*6, false)
		out += spellRange
	case SPA_FOCUS_INSTANT_ONLY:
		out += spellEffectName
	case SPA_FOCUS_LEVEL_MIN:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_FOCUS_CASTTIME_MIN, SPA_FOCUS_CASTTIME_MAX:
		out += formatSeconds(spellEffectName, value/1000, false)
		out += spellRange
	case SPA_NPC_PORTAL_WARDER_BANISH:
		out += fmt.Sprintf("%s to %d, %d, %d in %s", spellEffectName, se.Bases[0], se.Bases[1], se.Bases[2], ZoneLongNameByShortName(se.TeleportZone))
	case SPA_PORTAL_LOCATIONS:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_PERCENT_HEAL:
		out += formatMax(spellEffectName, value, max)
	case SPA_STACKING_BLOCK:
		//blockEffect := spellEffect(base, )
		blockEffect := fmt.Sprintf("Spell %d", base)
		out += formatStacking(spellEffectName, base2, value, max, spa, blockEffect)
	case SPA_STRIP_VIRTUAL_SLOT:
		tmpMax := max
		if max > 1000 {
			tmpMax = max - 1000
		}
		blockEffect := fmt.Sprintf("Spell %d", base)
		out += formatStacking(spellEffectName, calc-200, value, tmpMax, spa, blockEffect)
	case SPA_DIVINE_INTERVENTION:
		tmp = "Unknown"
		if base == 1 {
			tmp = "Partial"
		}
		if base == 2 {
			tmp = "Full"
		}
		tmp = fmt.Sprintf("Restore %s Health", tmp)
		out += formatExtra(spellEffectName, tmp, tmp, spellRange)
	case SPA_POCKET_PET:
		tmp = "(Unknown)"
		if base == 1 {
			tmp = "(Partial)"
		}
		if base == 2 {
			tmp = "(Full)"
		}
		out += formatExtra(spellEffectName, tmp, tmp, spellRange)
	case SPA_PET_SWARM:
		tmp = fmt.Sprintf("%s x%d for %dsec", se.TeleportZone, value, finish)
		out += formatExtra(spellEffectName, tmp, tmp, spellRange)
	case SPA_HEALTH_BALANCE:
		out += formatPenaltyChance(spellEffectName, value, "Penalty")
	case SPA_CANCEL_NEGATIVE_MAGIC:
		out += formatPenaltyChance(spellEffectName, base/10, "Chance")
	case SPA_POP_RESURRECT, SPA_MIRROR:
		out += spellEffectName
	case SPA_FEEDBACK:
		out += formatRange(spellEffectName, -value, extendedRange, spellRange)
	case SPA_REFLECT:
		out += formatPercent(spellEffectName, value, finish, false, false, percent)
	case SPA_MODIFY_ALL_STATS:
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_CHANGE_SOBRIETY:
		out += formatAT(spellEffectName, value, spellRange, "if Alchohol Tolerance is below")
	case SPA_SPELL_GUARD, SPA_MELEE_GUARD:
		out += formatPercent(spellEffectName, value, value, false, false, percent)
		if max > 0 {
			out += fmt.Sprintf(" until %d absorbed", max)
		}
	case SPA_ABSORB_HIT:
		tmp = fmt.Sprintf("up to %d from the next %d melee strikes or direct damage spells", max, value)
		out += formatString(spellEffectName, tmp, spellRange)
	case SPA_OBJECT_SENSE_TRAP, SPA_OBJECT_DISARM_TRAP, SPA_OBJECT_PICKLOCK:
		out += spellEffectName
	case SPA_FOCUS_PET:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_DEFENSIVE:
		out += formatPercent(spellEffectName, -value, -finish, false, false, spellEffectName)
	case SPA_CRITICAL_MELEE:
		out += formatSkills(spellEffectName, value, finish, base2, spellRange, "for")
	case SPA_CRITICAL_SPELL, SPA_CRIPPLING_BLOW, SPA_EVASION, SPA_RIPOSTE, SPA_DODGE, SPA_PARRY, SPA_DUAL_WIELD:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_DOUBLE_ATTACK:
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_MELEE_LIFETAP:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
		out += " heal"
	case SPA_PURETONE, SPA_SANCTIFICATION, SPA_FEARLESS, SPA_HUNDRED_HANDS:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_SKILL_INCREASE_CHANCE, SPA_ACCURACY, SPA_SKILL_DAMAGE_MOD, SPA_MIN_DAMAGE_DONE_MOD:
		out += formatSkills(spellEffectName, value, finish, base2, spellRange, "")
	case SPA_MANA_BALANCE:
		out += formatPenaltyChance(spellEffectName, value, "Penalty")
	case SPA_BLOCK:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_ENDURANCE:
		out += formatRange(spellEffectName, value, extendedRange, "")
		out += repeating
	case SPA_INCREASE_MAX_ENDURANCE:
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_AMNESIA:
		out += spellEffectName
	case SPA_HATE_OVER_TIME:
		out += formatRange(spellEffectName, value, extendedRange, "")
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_SKILL_ATTACK:
		out += formatSkillAttack(spellEffectName, value, max, base2, base, spellRange)
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_FADE:
		out += spellEffectName
	case SPA_STUN_RESIST, SPA_STRIKETHROUGH1:
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_SKILL_DAMAGE_TAKEN:
		out += formatSkills(spellEffectName, value, finish, base2, spellRange, "")
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_INSTANT_ENDURANCE:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_TAUNT:
		out += spellEffectName
	case SPA_PROC_CHANCE:
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_RANGE_ABILITY:
		out += formatExtra(spellEffectName, formatRateMod(spellEffectName, base, base2), spellRange, "")
	case SPA_ILLUSION_OTHERS, SPA_MASS_GROUP_BUFF:
		out += spellEffectName
	case SPA_GROUP_FEAR_IMMUNITY:
		out += formatSeconds(spellEffectName, value, true)
	case SPA_RAMPAGE, SPA_AE_TAUNT, SPA_FLESH_TO_BONE:
		out += spellEffectName
	case SPA_PURGE_POISON:
	case SPA_CANCEL_BENEFICIAL:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_SHIELD_CASTER:
		out += formatSeconds(spellEffectName, value*1.0, false)
	case SPA_DESTRUCTIVE_FORCE:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_FOCUS_FRENZIED_DEVASTATION:
		out += formatSeconds(spellEffectName, value, true)
	case SPA_PET_PCT_MAX_HP: //Pet HP
	case SPA_HP_MAX_HP:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_PET_PCT_AVOIDANCE:
	case SPA_MELEE_ACCURACY:
	case SPA_HEADSHOT:
	case SPA_PET_CRIT_MELEE:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_SLAY_UNDEAD:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_INCREASE_SKILL_DAMAGE:
		out += formatSkills(spellEffectName, value, finish, base2, percent, "")
	case SPA_REDUCE_WEIGHT:
	case SPA_BLOCK_BEHIND:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_DOUBLE_RIPOSTE:
	case SPA_ADD_RIPOSTE:
	case SPA_GIVE_DOUBLE_ATTACK:
	case SPA_2H_BASH:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_REDUCE_SKILL_TIMER:
		out += formatRefreshTimer(spellEffectName, -value, -finish, base2, spellRange)
	case SPA_ACROBATICS:
	case SPA_CAST_THROUGH_STUN:
	case SPA_EXTENDED_SHIELDING:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_BASH_CHANCE:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_DIVINE_SAVE:
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_METABOLISM:
		out += formatRange(spellEffectName, -value, extendedRange, spellRange)
	case SPA_POISON_MASTERY, SPA_FOCUS_CHANNELING, SPA_FREE_PET, SPA_PET_AFFINITY, SPA_PERM_ILLUSION, SPA_STONEWALL, SPA_STRING_UNBREAKABLE, SPA_IMPROVE_RECLAIM_ENERGY, SPA_INCREASE_CHANGE_MEMWIPE:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_ENHANCED_CHARM:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_ENHANCED_ROOT, SPA_TRAP_CIRCUMVENTION:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_INCREASE_AIR_SUPPLY:
		out += formatRange(spellEffectName, value, spellRange, "")
	case SPA_INCREASE_MAX_SKILL, SPA_EXTRA_SPECIALIZATION, SPA_OFFHAND_MIN_WEAPON_DAMAGE:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_INCREASE_PROC_CHANCE:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_ENDLESS_QUIVER:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_BACKSTAB_FRONT:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_CHAOTIC_STAB:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_NOSPELL, SPA_SHIELDING_DURATION_MOD:
		out += formatBasePercent(spellEffectName, base)
	case SPA_SHROUD_OF_STEALTH, SPA_GIVE_PET_HOLD:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_TRIPLE_BACKSTAB:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_AC_LIMIT_MOD, SPA_ADD_INSTRUMENT_MOD, SPA_SONG_MOD_CAP:
		out += formatBasePercent(spellEffectName, base)
	case SPA_INCREASE_STAT_CAP:
		out += formatStatsCapRange(spellEffectName, value, StatShortName(base2), extendedRange)
	case SPA_TRADESKILL_MASTERY, SPA_REDUCE_AA_TIMER:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_NO_FIZZLE:
		out += spellEffectName
	case SPA_ADD_2H_ATTACK_CHANCE, SPA_ADD_PET_COMMANDS:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_ALCHEMY_FAIL_RATE:
		out += formatSkills(spellEffectName, -value, -finish, base2, percent, "for")
	case SPA_FIRST_AID:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_EXTEND_SONG_RANGE:
		out += formatCount(spellEffectName, value, spellRange, "to")
	case SPA_BASE_RUN_MOD:
		out += formatBasePercent(spellEffectName, base)
	case SPA_INCREASE_CASTING_LEVEL:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_DOTCRIT, SPA_HEALCRIT, SPA_MENDCRIT:
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_DUAL_WIELD_AMT, SPA_EXTRA_DI_CHANCE, SPA_FINISHING_BLOW:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_FLURRY, SPA_PET_FLURRY:
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_PET_FEIGN, SPA_INCREASE_BANDAGE_AMT, SPA_WU_ATTACK, SPA_IMPROVE_LOH, SPA_NIMBLE_EVASION:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_FOCUS_DAMAGE_AMT:
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_FOCUS_DURATION_AMT:
		out += formatCount(spellEffectName, value, spellRange, percent)
		out += " tick(s)"
	case SPA_ADD_PROC_HIT:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_DOOM_EFFECT:
		out += formatExtra(spellEffectName, SpellName(base), spellRange, " on Fade")
	case SPA_INCREASE_RUN_SPEED_CAP, SPA_PURIFY, SPA_STRIKETHROUGH, SPA_STUN_RESIST2:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_SPELL_CRIT_CHANCE:
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_REDUCE_SPECIAL_TIMER:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_FOCUS_DAMAGE_MOD_DETRIMENTAL:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_FOCUS_DAMAGE_AMT_DETRIMENTAL:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_TINY_COMPANION:
		out += formatPercent(spellEffectName, -value, -finish, false, false, spellRange)
	case SPA_WAKE_DEAD:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_DOPPELGANGER:
		out += spellEffectName
	case SPA_INCREASE_RANGE_DMG:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_FOCUS_DAMAGE_MOD_CRIT, SPA_FOCUS_DAMAGE_AMT_CRIT:
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_SECONDARY_RIPOSTE_MOD:
		out += formatBasePercent(spellEffectName, base)
	case SPA_DAMAGE_SHIELD_MOD:
		out += formatPercent(spellEffectName, -value, -finish, false, false, spellRange)
	case SPA_WEAK_DEAD_2:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_APPRAISAL, SPA_ZONE_SUSPEND_MINION, SPA_TELEPORT_CASTERS_BINDPOINT:
		out += spellEffectName
	case SPA_FOCUS_REUSE_TIMER:
		out += formatTimer(spellEffectName, float64(-base/1000.0))
	case SPA_FOCUS_COMBAT_SKILL:
		out += spellEffectName
	case SPA_OBSERVER, SPA_FORAGE_MASTER:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_IMPROVED_INVIS, SPA_IMPROVED_INVIS_UNDEAD, SPA_IMPROVED_INVIS_ANIMALS:
		out += spellEffectName
	case SPA_INCREASE_WORN_HP_REGEN_CAP, SPA_INCREASE_WORN_MANA_REGEN_CAP:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_CRITICAL_HP_REGEN, SPA_SHIELD_BLOCK_CHANCE:
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_REDUCE_TARGET_HATE:
		out += formatBasePercent(spellEffectName, base)
	case SPA_GATE_STARTING_CITY:
		out += spellEffectName
	case SPA_DEFENSIVE_PROC:
		out += formatExtra(spellEffectName, formatRateMod(spellEffectName, base, base2), spellRange, "")
	case SPA_HP_FOR_MANA:
		out += formatBasePercent(spellEffectName, base)
	case SPA_NO_BREAK_AE_SNEAK, SPA_ADD_SPELL_SLOTS, SPA_ADD_BUFF_SLOTS:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_INCREASE_NEGATIVE_HP_LIMIT:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_MANA_ABSORB_PCT_DMG:
		out += formatCount(spellEffectName, value, "up to", "%")
	case SPA_CRIT_ATTACK_MODIFIER:
		out += formatSkills(spellEffectName, value, finish, base2, spellRange, "")
	case SPA_FAIL_ALCHEMY_ITEM_RECOVERY:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_SUMMON_TO_CORPSE:
		out += spellEffectName
	case SPA_DOOM_RUNE_EFFECT:
		out += formatExtra(spellEffectName, SpellName(base), spellRange, " on Fade")
	case SPA_NO_MOVE_HP:
		out += formatRange(spellEffectName, value, extendedRange, "")
		out += repeating
		out += " if target is not moving"
	case SPA_FOCUSED_IMMUNITY:
		out += spellEffectName
	case SPA_ILLUSIONARY_TARGET:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_INCREASE_EXP_MOD:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_EXPEDIENT_RECOVERY:
		out += spellEffectName
	case SPA_FOCUS_CASTING_PROC, SPA_CHANCE_SPELL:
		out += formatExtra(spellEffectName, formatSpellChance(spellEffectName, base, base2), spellRange, " on Cast")
	case SPA_WORN_ATTACK_CAP:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_NO_PANIC:
		out += spellEffectName
	case SPA_SPELL_INTERRUPT:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_ITEM_CHANNELING, SPA_ASSASSINATE_MAX_LEVEL, SPA_HEADSHOT_MAX_LEVEL, SPA_DOUBLE_RANGED_ATTACK, SPA_FOCUS_MANA_MIN, SPA_INCREASE_SHIELD_DMG:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_MANABURN:
		out += formatCount(spellEffectName, value*4, spellRange, "for")
	case SPA_SPAWN_INTERACTIVE_OBJECT:
		objectID := se.ID
		if se.SpellGroup == 3 {
			objectID += 1
		}
		objectName := SpellName(objectID)
		out += formatExtra(spellEffectName, objectName, spellRange, "")
	case SPA_INCREASE_TRAP_COUNT, SPA_INCREASE_SOI_COUNT, SPA_DEACTIVATE_ALL_TRAPS, SPA_LEARN_TRAP:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_CHANGE_TRIGGER_TYPE:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_FOCUS_MUTE:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_INSTANT_MANA:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_PASSIVE_SENSE_TRAP:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_PROC_ON_KILL_SHOT, SPA_PROC_ON_DEATH:
		out += formatExtra(spellEffectName, formatSpellChance(spellEffectName, base, base2), spellRange, "")
	case SPA_POTION_BELT, SPA_BANDOLIER:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_ADD_TRIPLE_ATTACK_CHANCE:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_PROC_ON_SPELL_KILL_SHOT:
		out += formatExtra(spellEffectName, formatSpellChance(spellEffectName, base, base2), spellRange, "")
	case SPA_GROUP_SHIELDING:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_MODIFY_BODY_TYPE:
		tmp = "Unknown"
		if base == 3 {
			tmp = "Undead"
		}
		if base == 21 {
			tmp = "Animal"
		}
		if base == 25 {
			tmp = "Plant"
		}
		out += formatString(spellEffectName, tmp, spellRange)
	case SPA_MODIFY_FACTION:
		out += formatExtra(spellEffectName, fmt.Sprintf("Faction ID %d", base), spellRange, "")
	case SPA_CORRUPTION, SPA_RESIST_CORRUPTION:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_SLOW:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_GRANT_FORAGING:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_DOOM_ALWAYS:
		out += formatExtra(spellEffectName, SpellName(base), spellRange, " on Fade")
	case SPA_TRIGGER_SPELL:
		out += formatExtra(spellEffectName, formatSpellChance(spellEffectName, base, base2), spellRange, "")
	case SPA_CRIT_DOT_DMG_MOD:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_FLING:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_DOOM_ENTITY:
		out += formatExtra(spellEffectName, SpellName(base), spellRange, " on Fade")
	case SPA_RESIST_OTHER_SPA:
		out += formatExtra(spellEffectName, formatResists(spellEffectName, base, base2), spellRange, "")
	case SPA_DIRECTIONAL_TELEPORT:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_EXPLOSIVE_KNOCKBACK:
		tmp = fmt.Sprintf("(%d) and Toss Up (%d)", base, base2)
		out += formatString(spellEffectName, tmp, spellRange)
	case SPA_FLING_TOWARD:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_SUPPRESSION:
		tmp = fmt.Sprintf("%s Effect", spellEffectNameByID(base2))
		out += formatExtra(spellEffectName, tmp, spellRange, "")
	case SPA_FOCUS_CASTING_PROC_NORMALIZED:
		out += formatExtra(spellEffectName, formatSpellChance(spellEffectName, base, base2), spellRange, " on Cast")
	case SPA_FLING_AT:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_FOCUS_WHICH_GROUP:
		out += formatExtra(spellEffectName, fmt.Sprintf("Spell Group ID %d", base), spellRange, "")
	case SPA_DOOM_DISPELLER:
		out += formatExtra(spellEffectName, SpellName(base), spellRange, " on Curer")
	case SPA_DOOM_DISPELLEE:
		out += formatExtra(spellEffectName, SpellName(base), spellRange, " on Fade")
	case SPA_SUMMON_ALL_CORPSES:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_REFRESH_SPELL_TIMER:
		out += formatCount(spellEffectName, -value, spellRange, "to")
	case SPA_LOCKOUT_SPELL_TIMER, SPA_FOCUS_MANA_MAX:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_FOCUS_HEAL_AMT:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_FOCUS_HEAL_MOD_BENEFICIAL:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_FOCUS_HEAL_AMT_BENEFICIAL:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_FOCUS_HEAL_MOD_CRIT:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_FOCUS_HEAL_AMT_CRIT:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_ADD_PET_AC:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_FOCUS_SWARM_PET_DURATION:
		out += formatSecondsCount(spellEffectName, float64(value/1000.0), spellRange)
	case SPA_FOCUS_TWINCAST_CHANCE:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_HEALBURN:
		tmp = fmt.Sprintf("use up to %d mana to heal your group", value)
		out += formatString(spellEffectName, tmp, spellRange)
	case SPA_MANA_IGNITE, SPA_ENDURANCE_IGNITE:
		out += formatCount(spellEffectName, -value, spellRange, "by up to")
	case SPA_FOCUS_SPELL_CLASS, SPA_FOCUS_SPELL_SUBCLASS, SPA_STAFF_BLOCK_CHANCE:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_DOOM_LIMIT_USE:
		out += formatExtra(spellEffectName, SpellName(base), spellRange, " on Max Hits")
	case SPA_DOOM_FOCUS_USED:
		out += formatExtra(spellEffectName, SpellName(base), spellRange, " on Focus Used")
	case SPA_LIMIT_HP, SPA_LIMIT_MANA, SPA_LIMIT_ENDURANCE:
		out += formatCount(spellEffectName, value, spellRange, "to")
	case SPA_FOCUS_LIMIT_CLASS:
		out += formatExtra(spellEffectName, types.NewClassesBitmask(base).String(), spellRange, "")
	case SPA_FOCUS_LIMIT_RACE:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_FOCUS_BASE_EFFECTS:
		out += formatCount(spellEffectName, value, "by", "%")
	case SPA_FOCUS_LIMIT_SKILL, SPA_FOCUS_LIMIT_ITEM_CLASS:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_AC2, SPA_MANA2:
		out += formatRange(spellEffectName, value, extendedRange, "")
	case SPA_FOCUS_INCREASE_SKILL_DMG_2:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_PROC_EFFECT_2:
		out += formatExtra(spellEffectName, formatRateMod(spellEffectName, base, base2), spellRange, "")
	case SPA_FOCUS_LIMIT_USE, SPA_FOCUS_LIMIT_USE_AMT, SPA_FOCUS_LIMIT_USE_MIN, SPA_FOCUS_LIMIT_USE_TYPE:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_GRAVITATE:
		out += formatBase(spellEffectName, base, base, spellRange)
		if maxTargets != "" {
			out += maxTargets
		}
	case SPA_FLY:
		out += spellEffectName
	case SPA_ADD_EXTENDED_TARGET_SLOTS:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_SKILL_PROC:
		out += formatExtra(spellEffectName, formatRateMod(spellEffectName, base, base2), spellRange, "")
	case SPA_PROC_SKILL_MODIFIER:
		tmp = "All Skills"
		if base >= 0 {
			tmp = library.SkillName(base)
		}
		out += formatExtra(spellEffectName, tmp, spellRange, "")
	case SPA_SKILL_PROC_SUCCESS:
		out += formatExtra(spellEffectName, formatRateMod(spellEffectName, base, base2), spellRange, "")
	case SPA_POST_EFFECT, SPA_POST_EFFECT_DATA, SPA_EXPAND_MAX_ACTIVE_TROPHY_BENEFITS:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_ADD_NORMALIZED_SKILL_MIN_DMG_AMT:
		out += formatExtra(spellEffectName, formatRateMod(spellEffectName, base, base2), spellRange, "")
	case SPA_ADD_NORMALIZED_SKILL_MIN_DMG_AMT_2:
	case SPA_FRAGILE_DEFENSE:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_FREEZE_BUFF_TIMER:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_TELEPORT_TO_ANCHOR:
	case SPA_TRANSLOCATE_TO_ANCHOR:
		tmp = "Unknown"
		if base == 50874 {
			tmp = "Guild Hall"
		}
		if base == 52584 {
			tmp = "Primary"
		}
		if base == 52585 {
			tmp = "Secondary"
		}
		out += formatString(spellEffectName, tmp, spellRange)
	case SPA_ASSASSINATE, SPA_FINISHING_BLOW_MAX, SPA_DISTANCE_REMOVAL:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_REQUIRE_TARGET_DOOM, SPA_REQUIRE_CASTER_DOOM:
		out += formatExtra(spellEffectName, SpellName(base), spellRange, "")
		if base2 > 0 {
			//GetSpellRestrictions(se, i, szTemp, sizeof(szTemp)
			out += fmt.Sprintf(" -- Restrictions: %d", index)
		}
	case SPA_IMPROVED_TAUNT:
		tmp = fmt.Sprintf(" up to %d and Reduce Ally Hate Generation by %d%%", base, base2)
		out += formatString(spellEffectName, tmp, spellRange)
	case SPA_ADD_MERC_SLOT, SPA_STACKER_A, SPA_STACKER_B, SPA_STACKER_C, SPA_STACKER_D:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_DOT_GUARD:
		tmp = fmt.Sprintf(" absorbing %d%% damage to a total of %d", value, max)
		out += formatString(spellEffectName, tmp, spellRange)
	case SPA_MELEE_THRESHOLD_GUARD, SPA_SPELL_THRESHOLD_GUARD:
		threshold := "spell"
		if spa == SPA_MELEE_THRESHOLD_GUARD {
			threshold = "melee"
		}
		tmp = fmt.Sprintf("absorbing %d%% of incoming %s damage in excess of %d to a total of %d", value, threshold, base2, max)
		out += formatString(spellEffectName, tmp, spellRange)
	case SPA_MELEE_THRESHOLD_DOOM:
		out += formatExtra(spellEffectName, SpellName(base), spellRange, "")
		out += fmt.Sprintf(" on %d Melee Damage Taken", base2)
	case SPA_SPELL_THRESHOLD_DOOM:
		out += formatExtra(spellEffectName, SpellName(base), spellRange, "")
		out += fmt.Sprintf(" on %d Spell Damage Taken", base2)
	case SPA_ADD_HATE_PCT, SPA_ADD_HATE_OVER_TIME_PCT:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_RESOURCE_TAP:
		tmp = "unknown"
		if base == 0 {
			tmp = "hit points"
		}
		if base == 1 {
			tmp = "mana"
		}
		if base == 2 {
			tmp = "endurance"
		}
		out += fmt.Sprintf("Return %.2f%% of direct damage as %s", float64(value/10.0), tmp)
	case SPA_FACTION_MOD:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_SKILL_DAMAGE_MOD_2:
		out += formatSkills(spellEffectName, value, finish, base2, spellRange, "")
	case SPA_OVERRIDE_NOT_FOCUSABLE:
		out += spellEffectName
	case SPA_FOCUS_DAMAGE_MOD_2, SPA_FOCUS_DAMAGE_AMT_2, SPA_SHIELD:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_PC_PET_RAMPAGE, SPA_PC_PET_AE_RAMPAGE, SPA_PC_PET_FLURRY:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_DAMAGE_SHIELD_MITIGATION_AMT:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_DAMAGE_SHIELD_MITIGATION_PCT:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_CHANCE_BEST_IN_SPELL_GROUP, SPA_TRIGGER_BEST_IN_SPELL_GROUP:
		out += formatExtra(spellEffectName, formatSpellGroupChance(spellEffectName, base, base2), spellRange, " on Cast")
	case SPA_DOUBLE_MELEE_ATTACKS:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_AA_BUY_NEXT_RANK:
		out += spellEffectName
	case SPA_DOUBLE_BACKSTAB_FRONT, SPA_PET_MELEE_CRIT_DMG_MOD:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_TRIGGER_SPELL_NON_ITEM:
		out += formatExtra(spellEffectName, SpellName(base2), spellRange, " on Cast")
	case SPA_WEAPON_STANCE, SPA_HATELIST_TO_TOP:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_HATELIST_TO_TAIL:
		out += formatExtra(spellEffectName, SpellName(base), spellRange, " on Cast")
	case SPA_FOCUS_LIMIT_MIN_VALUE, SPA_FOCUS_LIMIT_MAX_VALUE:
		tmpMax := "Min"
		if base < 0 {
			tmpMax = "Max"
		}
		tmp = fmt.Sprintf("%s %s", spellEffectName, tmpMax)
		out += formatMinMaxBase(tmp, base, base2)
	case SPA_FOCUS_CAST_SPELL_ON_LAND:
		out += formatExtra(spellEffectName, SpellName(base2), spellRange, " on Land and conditions are met")
	case SPA_SKILL_BASE_DAMAGE_MOD:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_FOCUS_INCOMING_DMG_MOD, SPA_FOCUS_INCOMING_DMG_AMT:
		out += formatRange(spellEffectName, value, extendedRange, " (after crit)")
	case SPA_FOCUS_LIMIT_CASTER_CLASS:
		out += formatExtra(spellEffectName, types.NewClassesBitmask(base).String(), spellRange, "")
	case SPA_FOCUS_LIMIT_SAME_CASTER:
		tmp = "(Different)"
		if base > 0 {
			tmp = "(Same)"
		}
		out += formatExtra(spellEffectName, tmp, spellRange, "")
	case SPA_EXTEND_TRADESKILL_CAP:
		out += fmt.Sprintf("%s (%d, %d, %d)", spellEffectName, base, base2, max)
	case SPA_DEFENDER_MELEE_FORCE_PCT:
		out += formatBase(spellEffectName, -base, max, spellRange)
	case SPA_WORN_ENDURANCE_REGEN_CAP:
		out += formatBase(spellEffectName, base, base, spellRange)
	case SPA_FOCUS_MIN_REUSE_TIME, SPA_FOCUS_MAX_REUSE_TIME:
		out += formatSeconds(spellEffectName, value/1000.0, false)
	case SPA_FOCUS_ENDURANCE_MIN, SPA_FOCUS_ENDURANCE_MAX, SPA_PET_ADD_ATK:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_FOCUS_DURATION_MAX:
		out += formatSeconds(spellEffectName, value*6, false)
	case SPA_CRIT_MELEE_DMG_MOD_MAX:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
		out += " of Base Damage (Non Stacking)"
	case SPA_FOCUS_CAST_PROC_NO_BYPASS:
		out += fmt.Sprintf("%s (%d, %d, %d)", spellEffectName, base, base2, max)
	case SPA_ADD_EXTRA_PRIMARY_ATTACK_PCT, SPA_ADD_EXTRA_SECONDARY_ATTACK_PCT, SPA_FOCUS_CAST_TIME_MOD2:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_FOCUS_CAST_TIME_AMT, SPA_FEARSTUN:
		out += formatSeconds(spellEffectName, value/1000.0, false)
	case SPA_MELEE_DMG_POSITION_MOD:
		out += formatSeconds(spellEffectName, value/10.0, false)
	case SPA_MELEE_DMG_POSITION_AMT:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_DMG_TAKEN_POSITION_MOD:
		out += formatSeconds(spellEffectName, value/10.0, false)
	case SPA_DMG_TAKEN_POSITION_AMT:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_AMPLIFY_MOD:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
		out += " (Before DoT Crit, After Nuke Crit)"
	case SPA_AMPLIFY_AMT:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_HEALTH_TRANSFER:
		out += fmt.Sprintf("%s (%d, %d, %d)", spellEffectName, base, base2, max)
	case SPA_FOCUS_RESIST_INCOMING:
		out += formatCount(spellEffectName, value, spellRange, percent)
	case SPA_ATTACK_ACCURACY_MAX:
		out += formatPercent(spellEffectName, value, finish, false, false, spellRange)
	case SPA_FOCUS_TIMER_MIN, SPA_PROC_TIMER_MOD, SPA_MANA_MAX, SPA_ENDURANCE_MAX, SPA_AC_AVOIDANCE_MAX, SPA_AC_MITIGATION_MAX, SPA_ATTACK_OFFENSE_MAX, SPA_LUCK_AMT, SPA_LUCK_PCT, SPA_ENDURANCE_ABSORB_PCT_DMG, SPA_INSTANT_MANA_PCT, SPA_INSTANT_ENDURANCE_PCT, SPA_DURATION_HP_PCT, SPA_DURATION_MANA_PCT, SPA_DURATION_ENDURANCE_PCT:
		out += fmt.Sprintf("%s (id=%d, base=%d, base2=%d, max=%d, calc=%d, value=%d)", spellEffectName, spa, base, base2, max, calc, value)
	default:
		out += fmt.Sprintf("%s (id=%d, base=%d, base2=%d, max=%d, calc=%d, value=%d)", spellEffectName, spa, base, base2, max, calc, value)
		return out
	}

	out += fmt.Sprintf(" (spa %d)", spa)
	return out
}
