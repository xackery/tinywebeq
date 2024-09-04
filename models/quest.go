package models

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type Quest struct {
	ID              int64  `db:"id"`
	Name            string `db:"name"`
	Level           int    `db:"level"`
	Icon            int    `db:"icon"`
	Entries         []*QuestEntry
	CacheKey        string `db:"key"`
	CacheExpiration int64
}

type QuestEntry struct {
	ItemID    int64  `db:"item_id"`
	ItemName  string `db:"item_name"`
	ZoneID    int32  `db:"zone_id"`
	NpcID     int64  `db:"npc_id"`
	NpcName   string `db:"npc_name"`
	Score     int    `db:"score"`
	Expansion int8   `db:"expansion"`
	FileName  string `db:"file_name"`
	UseCase   string `db:"use_case"`
}

func (t *Quest) Identifier() string {
	return "quest"
}

func (t *Quest) Key() string {
	return t.CacheKey
}

func (t *Quest) SetKey(key string) {
	t.CacheKey = key
}

func (t *Quest) SetExpiration(expiration int64) {
	t.CacheExpiration = expiration
}

func (t *Quest) Expiration() int64 {
	return t.CacheExpiration
}

func (t *Quest) Serialize() string {
	return serialize(t)
}

func (t *Quest) Deserialize(data string) error {
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

func (t *Quest) Expansion() int8 {
	expansion := int8(0)
	for _, entry := range t.Entries {
		if entry.Expansion > expansion {
			expansion = entry.Expansion
		}
	}
	return expansion
}
