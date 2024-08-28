package mysql

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/tinywebeq/db/mysql/storage/mysqlc"
)

var ()

// Mysql is a mysql database
type Mysql struct {
	db    *sqlx.DB
	conn  string
	query *mysqlc.Queries
}

// Open should rarely be used, if you must be sure to close the database
func (b *Mysql) Open(conn string) (*sqlx.DB, error) {
	if b.db != nil {
		return b.db, nil
	}
	b.conn = conn
	var err error
	b.db, err = sqlx.Open("mysql", conn)
	if err != nil {
		return b.db, err
	}
	b.query = mysqlc.New(b.db.DB)
	return b.db, err
}

// Close should be called when done with the database
func (b *Mysql) Close() error {
	if b.db == nil {
		return nil
	}
	return b.db.Close()
}

func (b *Mysql) Ping() error {
	var err error
	if b.db == nil {
		b.db, err = b.Open(b.conn)
		if err != nil {
			return err
		}

	}
	return b.db.Ping()
}
