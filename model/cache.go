package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"

	"github.com/xackery/tinywebeq/tlog"
)

type CacheSerializer interface {
	Serialize() string
	Deserialize(string) error
}

type CacheIdentifier interface {
	Identifier() string
	Key() string
	SetKey(string)
	SetExpiration(int64)
	Expiration() int64
	CacheSerializer
}

func serialize(data CacheSerializer) string {
	buf := bytes.Buffer{}
	e := gob.NewEncoder(&buf)
	err := e.Encode(data)
	if err != nil {
		tlog.Warnf("gob encode: %v", err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
