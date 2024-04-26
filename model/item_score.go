package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"
	"fmt"
)

type ItemScore struct {
	Entries         []*ItemScoreEntry
	CacheKey        string `db:"key"`
	CacheExpiration int64
}

type ItemScoreEntry struct {
	ItemID    int `db:"item_id"`
	Score     int `db:"score"`
	Expansion int `db:"expansion"`
}

func (t *ItemScore) Identifier() string {
	return "item_score"
}

func (t *ItemScore) Key() string {
	return t.CacheKey
}

func (t *ItemScore) SetKey(key string) {
	t.CacheKey = key
}

func (t *ItemScore) SetExpiration(expiration int64) {
	t.CacheExpiration = expiration
}

func (t *ItemScore) Expiration() int64 {
	return t.CacheExpiration
}

func (t *ItemScore) Serialize() string {
	return serialize(t)
}

func (t *ItemScore) Deserialize(data string) error {
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
