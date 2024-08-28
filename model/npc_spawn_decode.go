package model

import "github.com/xackery/tinywebeq/db/mysql/storage/mysqlc"

func (t *NpcSpawnEntry) DecodeNpcSpawns(in mysqlc.NpcSpawnsByNpcIDRow) {
	t.LongName = in.LongName
	t.ShortName = in.ShortName
	t.X = in.X
	t.Y = in.Y
	t.Z = in.Z
	t.Spawngroup = in.Spawngroup
	t.Spawngroupid = in.Spawngroupid
	t.Respawntime = in.Respawntime
}
