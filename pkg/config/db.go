package config

import "fmt"

type DBConfig struct {
	Username   string `yaml:"user" json:"user"`
	Password   string `yaml:"pass" json:"pass"`
	Hostname   string `yaml:"host" json:"host"`
	DBName     string `yaml:"dbname" json:"dbname"`
	TLSEnabled bool   `yaml:"tls-enabled" json:"tls-enabled"`
}

// DSN returns the Data Source Name
func (db *DBConfig) DSN() string {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s/%s",
		db.Username,
		db.Password,
		db.Hostname,
		db.DBName,
	)
	if !db.TLSEnabled {
		dsn = dsn + "?sslmode=disable"
	}
	return dsn
}
