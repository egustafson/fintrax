package locker_test

import (
	"testing"

	"github.com/egustafson/fintrax/pkg/locker"
	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"
)

func TestStoreConfig_YAMLUnmarshal_UnknownType(t *testing.T) {
	yamlData := `
type: unknown-store
`
	var sc locker.StoreConfig
	err := sc.YAMLUNmarshal(func(v interface{}) error {
		return yaml.Unmarshal([]byte(yamlData), v)
	})
	assert.Equal(t, locker.ErrUnknownStoreType, err)
}

func TestStoreConfig_YAMLUnmarshal_MissingType(t *testing.T) {
	yamlData := `
some-other-field: some-value
`
	var sc locker.StoreConfig
	err := sc.YAMLUNmarshal(func(v interface{}) error {
		return yaml.Unmarshal([]byte(yamlData), v)
	})
	assert.Equal(t, locker.ErrMissingStoreType, err)
}

func TestStoreConfig_YAMLUnmarshal_MemoryStore(t *testing.T) {
	yamlData := `
type: memory-store
`
	var sc locker.StoreConfig
	err := sc.YAMLUNmarshal(func(v interface{}) error {
		return yaml.Unmarshal([]byte(yamlData), v)
	})
	assert.NoError(t, err)
	memCfg, ok := sc.Store.(*locker.MemoryStoreConfig)
	assert.True(t, ok)
	assert.Equal(t, locker.MemoryStoreType, memCfg.GetType())
}

func TestStoreConfig_YAMLUnmarshal_YAMLStore(t *testing.T) {
	yamlData := `
type: yaml-store
file-path: /path/to/store.yaml
`
	var sc locker.StoreConfig
	err := sc.YAMLUNmarshal(func(v interface{}) error {
		return yaml.Unmarshal([]byte(yamlData), v)
	})
	assert.NoError(t, err)
	yamlCfg, ok := sc.Store.(*locker.YAMLStoreConfig)
	assert.True(t, ok)
	assert.Equal(t, "/path/to/store.yaml", yamlCfg.FilePath)
}

func TestStoreConfig_YAMLUnmarshal_DBStore(t *testing.T) {
	yamlData := `
type: db-store
dsn: user:password@tcp(localhost:3306)/dbname
table-name: lockers
`
	var sc locker.StoreConfig
	err := sc.YAMLUNmarshal(func(v interface{}) error {
		return yaml.Unmarshal([]byte(yamlData), v)
	})
	assert.NoError(t, err)
	dbCfg, ok := sc.Store.(*locker.DBStoreConfig)
	assert.True(t, ok)
	assert.Equal(t, "user:password@tcp(localhost:3306)/dbname", dbCfg.DSN)
	assert.Equal(t, "lockers", dbCfg.Table)
}
