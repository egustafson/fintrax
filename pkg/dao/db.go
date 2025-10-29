package dao

import (
	"errors"
	"log/slog"

	"github.com/egustafson/fintrax/pkg/config"
	"github.com/jmoiron/sqlx"
)

var (
	ErrorDBUninitalized = errors.New("db connector uninitialized")
	ErrorDBDisabled     = errors.New("db disabled in config")
)

func initDB(dbConfig *config.DBConfig) (db *sqlx.DB, err error) {

	if dbConfig.Disabled {
		return nil, ErrorDBDisabled
	}

	if db, err = sqlx.Open("pgx", dbConfig.DSN()); err != nil {
		slog.Error("failed to connect to db", "error", err)
		return nil, err
	}
	if err = db.Ping(); err != nil {
		db.Close()
		slog.Error("failed to connect to db", "error", err)
		return nil, err
	}
	slog.Info("db connected")
	return
}
