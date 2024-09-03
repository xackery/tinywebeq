package models

import "github.com/xackery/tinywebeq/db/mysql/storage/mysqlc"

func (t *NpcSpellEntry) DecodeNpcSpells(in mysqlc.NpcSpellsByNpcSpellIDRow) {
	t.ID = in.ID
	t.Name = in.Name
	t.ProcChance = in.ProcChance
	t.AttackProc = in.AttackProc
	t.RangeProc = in.RangeProc
	t.RprocChance = in.RprocChance
	t.DefensiveProc = in.DefensiveProc
	t.DprocChance = in.DprocChance
}
