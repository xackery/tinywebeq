package bbolt

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"

	"github.com/xackery/tinywebeq/model"
	bolt "go.etcd.io/bbolt"
)

func (b *BBolt) NpcQuestByNpcID(ctx context.Context, npcID int64) (*model.NpcQuest, error) {
	db, err := b.Open()
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	npcQuest := &model.NpcQuest{}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("npc_quest"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		v := b.Get(itob(npcID))
		if v == nil {
			return nil
		}
		return gob.NewDecoder(bytes.NewReader(v)).Decode(npcQuest)
	})
	if err != nil {
		return nil, fmt.Errorf("view: %w", err)
	}
	return npcQuest, nil
}

func (b *BBolt) NpcQuestReplace(ctx context.Context, npcID int64, npcQuest *model.NpcQuest) error {
	db, err := b.Open()
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("npc_quest"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		var buf bytes.Buffer
		err := gob.NewEncoder(&buf).Encode(npcQuest)
		if err != nil {
			return fmt.Errorf("encode: %w", err)
		}
		return b.Put(itob(npcID), buf.Bytes())
	})
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	return nil
}

func (b *BBolt) NpcQuestTruncate(ctx context.Context) error {
	db, err := b.Open()
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte("npc_quest"))
		if err != nil {
			return fmt.Errorf("delete bucket: %w", err)
		}
		_, err = tx.CreateBucket([]byte("npc_quest"))
		if err != nil {
			return fmt.Errorf("create bucket: %w", err)
		}
		return nil
	})
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	return nil
}
