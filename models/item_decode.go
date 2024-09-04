package models

import (
	"github.com/xackery/tinywebeq/db/mysql/storage/mysqlc"
)

func (t *Item) DecodeItemsAll(in mysqlc.ItemsAllRow) {
	t.ID = in.ID
	t.ID = in.ID
	t.Name = in.Name
	t.Ac = in.Ac
	t.Reqlevel = in.Reqlevel
	t.Reclevel = in.Reclevel
	t.Hp = in.Hp
	t.Damage = in.Damage
	t.Delay = in.Delay
	t.Mana = in.Mana
}
