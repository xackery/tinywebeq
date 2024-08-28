// db contains all database related functions
package db

import (
	"context"
	"fmt"
)

// Init initializes the database
func Init(ctx context.Context) error {
	err := initMySQL(ctx)
	if err != nil {
		return fmt.Errorf("init mysql: %w", err)
	}

	err = initBbolt(ctx)
	if err != nil {
		return fmt.Errorf("init bbolt: %w", err)
	}

	return nil
}
