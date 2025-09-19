package db

import (
	"errors"
	"log/slog"

	"github.com/egustafson/fintrax/pkg/config"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB

	UninitalizedDbError = errors.New("db connector uninitialized")
)

func Init(dbConfig *config.DBConfig) (err error) {

	if db, err = sqlx.Connect("postgres", dbConfig.DSN()); err != nil {
		slog.Error("failed to connect to db", "error", err)
		return
	}
	if err = db.Ping(); err != nil {
		slog.Error("failed to connect to db", "error", err)
		return
	}
	slog.Info("db connected")
	return
}

func Healthz() error {
	if db == nil {
		return UninitalizedDbError
	}
	return db.Ping()
}

// TODO: func Status() StatusST {}

func Shutdown() {
	db.Close()
}
