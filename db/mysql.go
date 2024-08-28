package db

import (
	"context"
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/tinywebeq/config"
	"github.com/xackery/tinywebeq/db/mysql"
	"github.com/xackery/tinywebeq/tlog"
)

var (
	Mysql = &mysql.Mysql{}
)

func initMySQL(ctx context.Context) error {
	err := mysqlConnect()
	if err != nil {
		return fmt.Errorf("connect: %w", err)
	}
	go mysqlLoop(ctx)
	return nil
}

func mysqlConnect() error {
	var err error
	var db *sqlx.DB
	tlog.Debugf("Connecting to database %s:%d", config.Get().Database.Host, config.Get().Database.Port)

	db, err = Mysql.Open(fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true&readTimeout=5s", config.Get().Database.Username, config.Get().Database.Password, config.Get().Database.Host, config.Get().Database.Port, config.Get().Database.Name))
	if err != nil {
		return fmt.Errorf("open: %w", err)
	}

	// ping instance
	err = db.Ping()
	if err != nil {
		return fmt.Errorf("ping: %w", err)
	}
	return nil
}

func mysqlLoop(ctx context.Context) {
	var err error
	for {
		select {
		case <-time.After(60 * time.Second):
			if err = Mysql.Ping(); err != nil {
				err = mysqlConnect()
				if err != nil {
					fmt.Println("db.loop: connect:", err)
				}
			}
		case <-ctx.Done():
			if Mysql != nil {
				Mysql.Close()
			}
		}
	}
}
