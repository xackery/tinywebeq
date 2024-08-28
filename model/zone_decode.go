package model

import (
	"github.com/xackery/tinywebeq/db/mysql/storage/mysqlc"
)

func (t *Zone) DecodeZone(in mysqlc.Zone) {
	t.ID = in.ID
	t.Zoneidnumber = in.Zoneidnumber
	t.Version = in.Version
	t.ShortName = ""
	if in.ShortName.Valid {
		t.ShortName = in.ShortName.String
	}
	t.LongName = in.LongName
	// t.MinStatus = in.MinStatus
	// t.MapFileName = in.MapFileName
	// t.Note = in.Note
	t.MinExpansion = in.MinExpansion
	t.MaxExpansion = in.MaxExpansion
	// t.ContentFlags = in.ContentFlags
	// t.ContentFlagsDisabled = in.ContentFlagsDisabled
	t.Expansion = in.Expansion
	// t.FileName = in.FileName
	// t.SafeX = in.SafeX
	// t.SafeY = in.SafeY
	// t.SafeZ = in.SafeZ
	// t.SafeHeading = in.SafeHeading
	// t.GraveyardID = in.GraveyardID
	t.MinLevel = in.MinLevel
	t.MaxLevel = in.MaxLevel
	// t.Timezone = in.Timezone
	// t.Maxclients = in.Maxclients
	// t.Ruleset = in.Ruleset
	// t.Underworld = in.Underworld
	// t.Minclip = in.Minclip
	// t.Maxclip = in.Maxclip
	// t.FogMinclip = in.FogMinclip
	// t.FogMaxclip = in.FogMaxclip
	// t.FogBlue = in.FogBlue
	// t.FogRed = in.FogRed
	// t.FogGreen = in.FogGreen
	// t.Sky = in.Sky
	// t.Ztype = in.Ztype
	// t.ZoneExpMultiplier = in.ZoneExpMultiplier
	// t.Walkspeed = in.Walkspeed
	// t.TimeType = in.TimeType
	// t.FogRed1 = in.FogRed1
	// t.FogGreen1 = in.FogGreen1
	// t.FogBlue1 = in.FogBlue1
	// t.FogMinclip1 = in.FogMinclip1
	// t.FogMaxclip1 = in.FogMaxclip1
	// t.FogRed2 = in.FogRed2
	// t.FogGreen2 = in.FogGreen2
	// t.FogBlue2 = in.FogBlue2
	// t.FogMinclip2 = in.FogMinclip2
	// t.FogMaxclip2 = in.FogMaxclip2
	// t.FogRed3 = in.FogRed3
	// t.FogGreen3 = in.FogGreen3
	// t.FogBlue3 = in.FogBlue3
	// t.FogMinclip3 = in.FogMinclip3
	// t.FogMaxclip3 = in.FogMaxclip3
	// t.FogRed4 = in.FogRed4
	// t.FogGreen4 = in.FogGreen4
	// t.FogBlue4 = in.FogBlue4
	// t.FogMinclip4 = in.FogMinclip4
	// t.FogMaxclip4 = in.FogMaxclip4
	// t.FogDensity = in.FogDensity
	// t.FlagNeeded = in.FlagNeeded
	// t.Canbind = in.Canbind
	// t.Cancombat = in.Cancombat
	// t.Canlevitate = in.Canlevitate
	// t.Castoutdoor = in.Castoutdoor
	// t.Hotzone = in.Hotzone
	// t.Insttype = in.Insttype
	// t.Shutdowndelay = in.Shutdowndelay
	// t.Peqzone = in.Peqzone
	// t.BypassExpansionCheck = in.BypassExpansionCheck
	// t.Suspendbuffs = in.Suspendbuffs
	// t.RainChance1 = in.RainChance1
	// t.RainChance2 = in.RainChance2
	// t.RainChance3 = in.RainChance3
	// t.RainChance4 = in.RainChance4
	// t.RainDuration1 = in.RainDuration1
	// t.RainDuration2 = in.RainDuration2
	// t.RainDuration3 = in.RainDuration3
	// t.RainDuration4 = in.RainDuration4
	// t.SnowChance1 = in.SnowChance1
	// t.SnowChance2 = in.SnowChance2
	// t.SnowChance3 = in.SnowChance3
	// t.SnowChance4 = in.SnowChance4
	// t.SnowDuration1 = in.SnowDuration1
	// t.SnowDuration2 = in.SnowDuration2
	// t.SnowDuration3 = in.SnowDuration3
	// t.SnowDuration4 = in.SnowDuration4
	// t.Gravity = in.Gravity
	// t.Type = in.Type
	// t.Skylock = in.Skylock
	// t.FastRegenHp = in.FastRegenHp
	// t.FastRegenMana = in.FastRegenMana
	// t.FastRegenEndurance = in.FastRegenEndurance
	// t.NpcMaxAggroDist = in.NpcMaxAggroDist
	// t.MaxMovementUpdateRange = in.MaxMovementUpdateRange
	// t.UnderworldTeleportIndex = in.UnderworldTeleportIndex
	// t.LavaDamage = in.LavaDamage
	// t.MinLavaDamage = in.MinLavaDamage
	// t.IdleWhenEmpty = in.IdleWhenEmpty
	// t.SecondsBeforeIdle = in.SecondsBeforeIdle
}
