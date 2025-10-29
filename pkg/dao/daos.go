package dao

import (
	"github.com/egustafson/fintrax/pkg/config"
	"github.com/egustafson/fintrax/pkg/mx"
	"github.com/jmoiron/sqlx"
)

type Factory interface {
	// Statusable allows the factory to report its status
	mx.Statusable
	// DB returns the underlying DB instance
	DB() *sqlx.DB
	// Shutdown closes any resources held by the factory
	Shutdown() error
}

type daoFactory struct {
	db *sqlx.DB
}

func NewFactory(dbConfig *config.DBConfig) (Factory, error) {
	db, err := initDB(dbConfig)
	if err != nil {
		return nil, err
	}
	return &daoFactory{
		db: db,
	}, nil
}

func (f *daoFactory) DB() *sqlx.DB {
	return f.db
}

func (f *daoFactory) TypeID() string {
	return "dao-factory"
}

func (f *daoFactory) Status() mx.StatusObj {
	if f.db == nil {
		return mx.StatusObj{Health: mx.UNKNOWN}
	}
	if f.db.Ping() != nil {
		return mx.StatusObj{Health: mx.IMPAIRED}
	}
	return mx.StatusObj{Health: mx.OK}
}

func (f *daoFactory) Shutdown() error {
	if f.db != nil {
		return f.db.Close()
	}
	return nil
}
