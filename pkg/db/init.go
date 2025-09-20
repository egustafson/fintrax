package db

import (
	"errors"
	"log/slog"

	"github.com/egustafson/fintrax/pkg/config"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB

	ErrorUninitalizedDB = errors.New("db connector uninitialized")
)

func Init(dbConfig *config.DBConfig) (err error) {

	if db, err = sqlx.Open("pgx", dbConfig.DSN()); err != nil {
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
		return ErrorUninitalizedDB
	}
	return db.Ping()
}

// TODO: func Status() StatusST {}

func Shutdown() {
	db.Close()
}
