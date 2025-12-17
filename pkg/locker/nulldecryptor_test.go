package locker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/egustafson/fintrax/pkg/locker"
)

func TestMakeNullDecryptor(t *testing.T) {
	config := &locker.NullDecryptorConfig{}
	dec := locker.MakeNullDecryptor(config)
	assert.NotNil(t, dec)
}

func TestNullDecryptor_Decrypt(t *testing.T) {
	dec := locker.MakeNullDecryptor(&locker.NullDecryptorConfig{})

	testData := "test data"
	result, err := dec.Decrypt(testData)
	require.NoError(t, err)
	assert.Equal(t, testData, result)
}

func TestNullDecryptor_Encrypt(t *testing.T) {
	dec := locker.MakeNullDecryptor(&locker.NullDecryptorConfig{})

	testData := "test data"
	result, err := dec.Encrypt(testData)
	require.NoError(t, err)
	assert.Equal(t, testData, result)
}
