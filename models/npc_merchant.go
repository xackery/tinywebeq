package models

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type NpcMerchant struct {
	Entries         []*NpcMerchantEntry
	CacheKey        string `db:"key"`
	CacheExpiration int64
}

type NpcMerchantEntry struct {
	ID        int32
	Name      string
	Price     int32
	Ldonprice int32
	Icon      int32
}

func (t *NpcMerchant) Identifier() string {
	return "npc_merchant"
}

func (t *NpcMerchant) Key() string {
	return t.CacheKey
}

func (t *NpcMerchant) SetKey(key string) {
	t.CacheKey = key
}

func (t *NpcMerchant) SetExpiration(expiration int64) {
	t.CacheExpiration = expiration
}

func (t *NpcMerchant) Expiration() int64 {
	return t.CacheExpiration
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
