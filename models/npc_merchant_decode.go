package models

import "github.com/xackery/tinywebeq/db/mysql/storage/mysqlc"

func (t *NpcMerchantEntry) DecodeNpcMerchantsByMerchantIDRow(in mysqlc.NpcMerchantsByMerchantIDRow) {
	t.ID = in.ID
	t.Name = in.Name
	t.Price = in.Price
	t.Ldonprice = in.Ldonprice
	t.Icon = in.Icon
}
