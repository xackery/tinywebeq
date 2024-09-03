package models

import (
	"github.com/xackery/tinywebeq/db/mysql/storage/mysqlc"
	"github.com/xackery/tinywebeq/library"
)

func (t *NpcLootEntry) DecodeNpcLootsDiscoveredOnly(in mysqlc.NpcLootsDiscoveredOnlyByLootTableIDRow) {
	t.ID = in.ID
	t.Name = in.Name
	t.Itemtype = library.ItemType(in.Itemtype)
	t.Chance = in.Chance
	t.Probability = in.Probability
	t.LootdropID = in.LootdropID
	t.Multiplier = in.Multiplier
}

func (t *NpcLootEntry) DecodeNpcLoots(in mysqlc.NpcLootsByLootTableIDRow) {
	t.ID = in.ID
	t.Name = in.Name
	t.Itemtype = library.ItemType(in.Itemtype)
	t.Chance = in.Chance
	t.Probability = in.Probability
	t.LootdropID = in.LootdropID
	t.Multiplier = in.Multiplier
}
