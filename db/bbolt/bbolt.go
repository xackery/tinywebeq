package bbolt

import (
	"context"
	"encoding/binary"
	"fmt"
	"os"
	"path/filepath"
	"time"

	bolt "go.etcd.io/bbolt"
)

// BBolt is a bbolt database
type BBolt struct {
	path string
}

// New creates a new bbolt database
func New(path string) *BBolt {
	return &BBolt{
		path: path,
	}
}

// Open should rarely be used, if you must be sure to close the database
func (b *BBolt) Open() (*bolt.DB, error) {
	err := os.MkdirAll(filepath.Dir(b.path), 0700)
	if err != nil {
		return nil, fmt.Errorf("mkdir all: %w", err)
	}

	return bolt.Open(b.path, 0600, &bolt.Options{Timeout: 1 * time.Second})
}

// itob returns an 8-byte big endian representation of v.
func itob(v int64) []byte {
	b := make([]byte, 8)
	binary.BigEndian.PutUint64(b, uint64(v))
	return b
}

func (b *BBolt) BucketCreate(ctx context.Context, name string) error {
	db, err := b.Open()
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}
	defer db.Close()
	err = db.Update(func(tx *bolt.Tx) error {
		// check if bucket exists

		bucket := tx.Bucket([]byte(name))
		if bucket != nil {
			return nil
		}

		_, err := tx.CreateBucket([]byte(name))
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
