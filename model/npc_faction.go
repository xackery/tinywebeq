package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type NpcFaction struct {
	Entries    []*NpcFactionEntry
	key        string
	expiration int64
}

type NpcFactionEntry struct {
	ID    int    `db:"id"`
	Name  string `db:"name"`
	Value int    `db:"value"`
}

func (t *NpcFaction) Identifier() string {
	return "npc_faction"
}

func (t *NpcFaction) Key() string {
	return t.key
}

func (t *NpcFaction) SetKey(key string) {
	t.key = key
}

func (t *NpcFaction) SetExpiration(expiration int64) {
	t.expiration = expiration
}

func (t *NpcFaction) Expiration() int64 {
	return t.expiration
}

func (t *NpcFaction) Serialize() string {
	return serialize(t)
}

func (t *NpcFaction) Deserialize(data string) error {
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
