package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type NpcSpawn struct {
	Entries    []*NpcSpawnEntry
	key        string
	expiration int64
}

type NpcSpawnEntry struct {
	Spawngroupid int     `db:"spawngroupid"`
	Spawngroup   string  `db:"spawngroup"`
	LongName     string  `db:"long_name"`
	ShortName    string  `db:"short_name"`
	X            float32 `db:"x"`
	Y            float32 `db:"y"`
	Z            float32 `db:"z"`
	Respawntime  int     `db:"respawntime"`
}

func (t *NpcSpawn) Identifier() string {
	return "npc_spawn"
}

func (t *NpcSpawn) Key() string {
	return t.key
}

func (t *NpcSpawn) SetKey(key string) {
	t.key = key
}

func (t *NpcSpawn) SetExpiration(expiration int64) {
	t.expiration = expiration
}

func (t *NpcSpawn) Expiration() int64 {
	return t.expiration
}

func (t *NpcSpawn) Serialize() string {
	return serialize(t)
}

func (t *NpcSpawn) Deserialize(data string) error {
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
