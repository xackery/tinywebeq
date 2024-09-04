package models

import "github.com/xackery/tinywebeq/db/mysql/storage/mysqlc"

func (t *NpcFactionEntry) DecodeNpcFactionsByFactionIDRow(row mysqlc.NpcFactionsByFactionIDRow) {
	t.ID = row.ID
	t.Name = row.Name
	t.Value = row.Value
}
