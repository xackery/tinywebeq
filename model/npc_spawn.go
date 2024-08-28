package model

import (
	"bytes"
	"database/sql"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type NpcSpawn struct {
	Entries         []*NpcSpawnEntry
	CacheKey        string `db:"key"`
	CacheExpiration int64
}

type NpcSpawnEntry struct {
	LongName     string
	ShortName    sql.NullString
	X            float64
	Y            float64
	Z            float64
	Spawngroup   string
	Spawngroupid int32
	Respawntime  int32
}

func (t *NpcSpawn) Identifier() string {
	return "npc_spawn"
}

func (t *NpcSpawn) Key() string {
	return t.CacheKey
}

func (t *NpcSpawn) SetKey(key string) {
	t.CacheKey = key
}

func (t *NpcSpawn) SetExpiration(expiration int64) {
	t.CacheExpiration = expiration
}

func (t *NpcSpawn) Expiration() int64 {
	return t.CacheExpiration
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
