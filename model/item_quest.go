package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/xackery/tinywebeq/library"
	"github.com/xackery/tinywebeq/tlog"
)

type ItemQuest struct {
	ID              int
	Entries         []*ItemQuestEntry
	CacheKey        string `db:"key"`
	CacheExpiration int64
}

type ItemQuestEntry struct {
	ItemID    int    `db:"item_id"`
	NpcID     int    `db:"npc_id"`
	NpcName   string `db:"npc_name"`
	ZoneID    int    `db:"zone_id"`
	UseCase   string `db:"use_case"`
	QuestID   int    `db:"quest_id"`
	QuestName string `db:"quest_name"`
}

func (t *ItemQuest) Identifier() string {
	return "item_quest"
}

func (t *ItemQuest) Key() string {
	return t.CacheKey
}

func (t *ItemQuest) SetKey(key string) {
	t.CacheKey = key
}

func (t *ItemQuest) SetExpiration(expiration int64) {
	t.CacheExpiration = expiration
}

func (t *ItemQuest) Expiration() int64 {
	return t.CacheExpiration
}

func (t *ItemQuest) Serialize() string {
	return serialize(t)
}

func (t *ItemQuest) Deserialize(data string) error {
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
	for _, entry := range t.Entries {
		tlog.Debugf("entry: %+v", entry)
	}
	return nil
}

func (t *ItemQuestEntry) ZoneLongName() string {
	return library.ZoneLongNameByID(t.ZoneID)
}

func (t *ItemQuestEntry) NpcCleanName() string {
	out := t.NpcName
	out = strings.ReplaceAll(out, "_", " ")
	out = strings.ReplaceAll(out, "-", "`")
	out = strings.ReplaceAll(out, "#", "")
	out = strings.ReplaceAll(out, "!", "")
	out = strings.ReplaceAll(out, "~", "")
	return out
}

func (t *ItemQuest) RewardEntries() []*ItemQuestEntry {
	var out []*ItemQuestEntry
	for _, entry := range t.Entries {
		if entry.UseCase == "success" {
			out = append(out, entry)
		}
	}
	return out
}

func (t *ItemQuest) ComponentEntries() []*ItemQuestEntry {
	var out []*ItemQuestEntry
	for _, entry := range t.Entries {
		if entry.UseCase == "component" {
			out = append(out, entry)
		}
	}
	return out
}
