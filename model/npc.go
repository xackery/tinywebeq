package model

type Npc struct {
	key        string
	expiration int64
	ID         int    `db:"id"`
	Name       string `db:"name"`
}

func (t *Npc) Identifier() string {
	return "npc"
}

func (t *Npc) Key() string {
	return t.key
}

func (t *Npc) SetKey(key string) {
	t.key = key
}

func (t *Npc) SetExpiration(expiration int64) {
	t.expiration = expiration
}

func (t *Npc) Expiration() int64 {
	return t.expiration
}

func (t *Npc) Serialize() string {
	return serialize(t)
}
