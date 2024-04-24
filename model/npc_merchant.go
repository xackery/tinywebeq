package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type NpcMerchant struct {
	Entries    []*NpcMerchantEntry
	key        string
	expiration int64
}

type NpcMerchantEntry struct {
	ID        int    `db:"id"`
	Name      string `db:"name"`
	Price     int    `db:"price"`
	Ldonprice int    `db:"ldonprice"`
	Icon      string `db:"icon"`
}

func (t *NpcMerchant) Identifier() string {
	return "npc_merchant"
}

func (t *NpcMerchant) Key() string {
	return t.key
}

func (t *NpcMerchant) SetKey(key string) {
	t.key = key
}

func (t *NpcMerchant) SetExpiration(expiration int64) {
	t.expiration = expiration
}

func (t *NpcMerchant) Expiration() int64 {
	return t.expiration
}

func (t *NpcMerchant) Serialize() string {
	return serialize(t)
}

func (t *NpcMerchant) Deserialize(data string) error {
	decoded, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return fmt.Errorf("base64 decode: %w", err)
	}
	buf := bytes.NewBuffer(decoded)
	d := gob.NewDecoder(buf)

	err = d.Decode(&t)
	if err != nil {
		return fmt.Errorf("gob decode: %w", err)
	}
	return nil
}
