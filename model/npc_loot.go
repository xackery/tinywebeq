package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"

	"github.com/xackery/tinywebeq/library"
)

type NpcLoot struct {
	Entries         []*NpcLootEntry
	CacheKey        string `db:"key"`
	CacheExpiration int64
}

type NpcLootEntry struct {
	ID          int     `db:"id"`
	Name        string  `db:"name"`
	Itemtype    int     `db:"itemtype"`
	Chance      float32 `db:"chance"`
	Probability int     `db:"probability"`
	Lootdropid  int     `db:"lootdrop_id"`
	Multiplier  int     `db:"multiplier"`
}

func (t *NpcLoot) Identifier() string {
	return "npc_loot"
}

func (t *NpcLoot) Key() string {
	return t.CacheKey
}

func (t *NpcLoot) SetKey(key string) {
	t.CacheKey = key
}

func (t *NpcLoot) SetExpiration(expiration int64) {
	t.CacheExpiration = expiration
}

func (t *NpcLoot) Expiration() int64 {
	return t.CacheExpiration
}

func (t *NpcLoot) Serialize() string {
	return serialize(t)
}

func (t *NpcLoot) Deserialize(data string) error {
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

func (t *NpcLootEntry) ChanceGlobal() int {
	return int(t.Chance * float32(t.Probability) / 100)
}

func (t *NpcLootEntry) ItemTypeStr() string {
	return library.ItemTypeStr(t.Itemtype)
}
