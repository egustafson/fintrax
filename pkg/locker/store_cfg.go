package locker

import "fmt"

const (
	MemoryStoreType = "memory-store"
	YAMLStoreType   = "yaml-store"
	DBStoreType     = "db-store"
)

var (
	ErrUnknownStoreType = fmt.Errorf("unknown store type")
	ErrMissingStoreType = fmt.Errorf("missing store type")
)

type StoreConfig struct {
	Store StoreType `yaml:",inline"`
}

type StoreType interface {
	GetType() string
}

type MemoryStoreConfig struct{}

func (m MemoryStoreConfig) GetType() string {
	return MemoryStoreType
}

type YAMLStoreConfig struct {
	FilePath string `yaml:"file-path"`
}

func (y YAMLStoreConfig) GetType() string {
	return YAMLStoreType
}

type DBStoreConfig struct {
	DSN   string `yaml:"dsn"`
	Table string `yaml:"table-name"`
}

func (d DBStoreConfig) GetType() string {
	return DBStoreType
}

func (sc *StoreConfig) YAMLUNmarshal(unmarshal func(interface{}) error) error {
	var typeHolder struct {
		StoreType string `yaml:"type"`
	}
	if err := unmarshal(&typeHolder); err != nil {
		return err
	}
	if typeHolder.StoreType == "" {
		return ErrMissingStoreType
	}

	switch typeHolder.StoreType {
	case MemoryStoreType:
		var cfg MemoryStoreConfig
		if err := unmarshal(&cfg); err != nil {
			return err
		}
		sc.Store = &cfg
	case YAMLStoreType:
		var cfg YAMLStoreConfig
		if err := unmarshal(&cfg); err != nil {
			return err
		}
		sc.Store = &cfg
	case DBStoreType:
		var cfg DBStoreConfig
		if err := unmarshal(&cfg); err != nil {
			return err
		}
		sc.Store = &cfg
	default:
		return ErrUnknownStoreType
	}
	return nil
}
