package repo

import (
	"github.com/jmoiron/sqlx"
)

type Logger interface {
	Info(...any)
	Error(...any)
	Debug(...any)
}

type Repo struct {
	db     *sqlx.DB
	logger Logger
}

const driverName = "mysql"

func New(logger Logger, connString string) (*Repo, error) {
	db, err := sqlx.Connect(driverName, connString)
	if err != nil {
		return nil, err
	}

	return &Repo{
		db:     db,
		logger: logger,
	}, nil
}
