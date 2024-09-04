package models

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type NpcQuest struct {
	ID              int64
	Entries         []*NpcQuestEntry
	CacheKey        string `db:"key"`
	CacheExpiration int64
}

type NpcQuestEntry struct {
	QuestID   int64  `db:"quest_id"`
	ZoneID    int32  `db:"zone_id"`
	QuestName string `db:"quest_name"`
}

func (t *NpcQuest) Identifier() string {
	return "npc_quest"
}

func (t *NpcQuest) Key() string {
	return t.CacheKey
}

func (t *NpcQuest) SetKey(key string) {
	t.CacheKey = key
}

func (t *NpcQuest) SetExpiration(expiration int64) {
	t.CacheExpiration = expiration
}

func (t *NpcQuest) Expiration() int64 {
	return t.CacheExpiration
}

func (t *NpcQuest) Serialize() string {
	return serialize(t)
}

func (t *NpcQuest) Deserialize(data string) error {
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
