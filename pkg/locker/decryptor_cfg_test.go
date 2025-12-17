package locker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/yaml.v3"

	"github.com/egustafson/fintrax/pkg/locker"
)

func TestDecryptorConfig_YAMLUnmarshal_UnknownType(t *testing.T) {
	yamlData := `
type: unknown-decryptor
`
	var dc locker.DecryptorConfig
	err := dc.YAMLUNmarshal(func(v interface{}) error {
		return yaml.Unmarshal([]byte(yamlData), v)
	})
	assert.Equal(t, locker.ErrUnknownDecryptorType, err)
}

func TestDecryptorConfig_YAMLUnmarshal_MissingType(t *testing.T) {
	yamlData := `
some-other-field: some-value
`
	var dc locker.DecryptorConfig
	err := dc.YAMLUNmarshal(func(v interface{}) error {
		return yaml.Unmarshal([]byte(yamlData), v)
	})
	assert.Equal(t, locker.ErrMissingDecryptorType, err)
}

func TestDecryptorConfig_YAMLUnmarshal_NullDecryptor(t *testing.T) {
	yamlData := `
type: null-decryptor
`
	var dc locker.DecryptorConfig
	err := dc.YAMLUNmarshal(func(v interface{}) error {
		return yaml.Unmarshal([]byte(yamlData), v)
	})
	assert.NoError(t, err)
	_, ok := dc.Decryptor.(*locker.NullDecryptorConfig)
	assert.True(t, ok)
}

func TestDecryptorConfig_YAMLUnmarshal_PwAESDecryptor(t *testing.T) {
	yamlData := `
type: pw-aes-decryptor
password: mysecretpassword
`
	var dc locker.DecryptorConfig
	err := dc.YAMLUNmarshal(func(v interface{}) error {
		return yaml.Unmarshal([]byte(yamlData), v)
	})
	assert.NoError(t, err)
	pwCfg, ok := dc.Decryptor.(*locker.PwAESDecryptorConfig)
	assert.True(t, ok)
	assert.Equal(t, "mysecretpassword", pwCfg.Password)
}

func TestDecryptorConfig_YAMLUnmarshal_YubiKeyDecryptor(t *testing.T) {
	yamlData := `
type: yk-decryptor
slot: 1
pin: 123456
`
	var dc locker.DecryptorConfig
	err := dc.YAMLUNmarshal(func(v interface{}) error {
		return yaml.Unmarshal([]byte(yamlData), v)
	})
	assert.NoError(t, err)
	ykCfg, ok := dc.Decryptor.(*locker.YubiKeyDecryptorConfig)
	assert.True(t, ok)
	assert.Equal(t, 1, ykCfg.Slot)
	assert.Equal(t, "123456", ykCfg.PIN)
}
