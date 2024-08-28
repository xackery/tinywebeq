package bbolt

import (
	"bytes"
	"context"
	"encoding/gob"
	"fmt"
	"strings"

	"github.com/xackery/tinywebeq/model"
	bolt "go.etcd.io/bbolt"
)

func (b *BBolt) QuestByQuestID(ctx context.Context, questID int64) (*model.Quest, error) {
	db, err := b.Open()
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	quest := &model.Quest{}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("quest"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		v := b.Get(itob(questID))
		if v == nil {
			return nil
		}
		return gob.NewDecoder(bytes.NewReader(v)).Decode(quest)
	})
	if err != nil {
		return nil, fmt.Errorf("view: %w", err)
	}
	return quest, nil
}

func (b *BBolt) QuestReplace(ctx context.Context, questID int64, quest *model.Quest) error {
	db, err := b.Open()
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("quest"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		var buf bytes.Buffer
		err := gob.NewEncoder(&buf).Encode(quest)
		if err != nil {
			return fmt.Errorf("encode: %w", err)
		}
		return b.Put(itob(questID), buf.Bytes())
	})
	if err != nil {
		return fmt.Errorf("update: %w", err)
	}
	return nil

}

func (b *BBolt) QuestsAll(ctx context.Context) (map[int64]*model.Quest, error) {
	db, err := b.Open()
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	quests := map[int64]*model.Quest{}
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("quest"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		return b.ForEach(func(k, v []byte) error {
			quest := &model.Quest{}
			err := gob.NewDecoder(bytes.NewReader(v)).Decode(quest)
			if err != nil {
				return fmt.Errorf("decode: %w", err)
			}
			quests[int64(k[0])] = quest
			return nil
		})
	})
	if err != nil {
		return nil, fmt.Errorf("view: %w", err)
	}
	return quests, nil
}

func (b *BBolt) QuestTruncate(ctx context.Context) error {
	db, err := b.Open()
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		err := tx.DeleteBucket([]byte("quest"))
		if err != nil {
			return fmt.Errorf("delete bucket: %w", err)
		}
		_, err = tx.CreateBucket([]byte("quest"))
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

func (b *BBolt) QuestSearchByName(ctx context.Context, name string) ([]*model.Quest, error) {
	db, err := b.Open()
	if err != nil {
		return nil, fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	quests := []*model.Quest{}
	names := strings.Split(name, " ")
	err = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("quest"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		return b.ForEach(func(k, v []byte) error {
			quest := &model.Quest{}
			err := gob.NewDecoder(bytes.NewReader(v)).Decode(quest)
			if err != nil {
				return fmt.Errorf("decode: %w", err)
			}
			for _, n := range names {
				if strings.Contains(strings.ToLower(quest.Name), strings.ToLower(n)) {
					quests = append(quests, quest)
					return nil
				}
			}
			return nil
		})
	})
	if err != nil {
		return nil, fmt.Errorf("view: %w", err)
	}
	return quests, nil
}

func (b *BBolt) QuestNextID(ctx context.Context) (int64, error) {
	db, err := b.Open()
	if err != nil {
		return 0, fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	var id int64
	err = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("quest"))
		if b == nil {
			return fmt.Errorf("bucket not found")
		}
		nextID, err := b.NextSequence()
		if err != nil {
			return fmt.Errorf("next sequence: %w", err)
		}
		id = int64(nextID)
		return nil
	})
	if err != nil {
		return 0, fmt.Errorf("update: %w", err)
	}
	return id, nil
}
