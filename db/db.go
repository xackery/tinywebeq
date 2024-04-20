package db

import (
	"context"
	"fmt"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	Instance *sqlx.DB
	mu       sync.RWMutex
)

// Init initializes the database
func Init(ctx context.Context) error {
	err := connect()
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	go loop(ctx)
	return nil
}

func Query(ctx context.Context, query string, args map[string]interface{}) (*sqlx.Rows, error) {
	tlog.Debugf("querying `%s`, args: %v", query, args)
	return Instance.NamedQueryContext(ctx, query, args)
}

func connect() error {
	var err error
	var db *sqlx.DB
	tlog.Debugf("Connecting to database %s:%d", config.Get().Database.Host, config.Get().Database.Port)

	db, err = sqlx.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&timeout=5s", config.Get().Database.Username, config.Get().Database.Password, config.Get().Database.Host, config.Get().Database.Port, config.Get().Database.Name))
	if err != nil {
		return fmt.Errorf("sql.Open: %w", err)
	}
	Instance = db
	// ping instance
	err = Instance.Ping()
	if err != nil {
		return fmt.Errorf("ping: %w", err)
	}
	return nil
}

func loop(ctx context.Context) {
	var err error
	for {
		select {
		case <-time.After(60 * time.Second):
			if err = Instance.Ping(); err != nil {
				err = connect()
				if err != nil {
					fmt.Println("db.loop: connect:", err)
				}
			}
		case <-ctx.Done():
			if Instance != nil {
				Instance.Close()
			}
		}
	}
}
