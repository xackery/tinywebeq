package model

import (
	"bytes"
	"encoding/base64"
	"encoding/gob"

	"github.com/xackery/tinywebeq/tlog"
)

type CacheIdentifier interface {
	Identifier() string
	Key() string
	SetKey(string)
	SetExpiration(int64)
	Serialize() string
	Expiration() int64
}

func serialize(data CacheIdentifier) string {
	buf := bytes.Buffer{}
	e := gob.NewEncoder(&buf)
	err := e.Encode(data)
	if err != nil {
		tlog.Warnf("gob encode: %v", err)
	}
	return base64.StdEncoding.EncodeToString(buf.Bytes())
}
