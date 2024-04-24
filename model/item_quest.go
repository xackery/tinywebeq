package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/xackery/tinywebeq/library"
)

type ItemQuest struct {
	Entries    []*ItemQuestEntry
	key        string
	expiration int64
}

type ItemQuestEntry struct {
	ItemID  int    `db:"item_id"`
	NpcID   int    `db:"npc_id"`
	NpcName string `db:"npc_name"`
	ZoneID  int    `db:"zone_id"`
	UseCase string `db:"use_case"`
}

func (t *ItemQuest) Identifier() string {
	return "item_quest"
}

func (t *ItemQuest) Key() string {
	return t.key
}

func (t *ItemQuest) SetKey(key string) {
	t.key = key
}

func (t *ItemQuest) SetExpiration(expiration int64) {
	t.expiration = expiration
}

func (t *ItemQuest) Expiration() int64 {
	return t.expiration
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
