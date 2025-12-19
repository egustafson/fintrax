package locker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/egustafson/fintrax/pkg/locker"
)

func TestMakeYAMLLocker_ValidYAML(t *testing.T) {
	yamlContent := []byte(`
decryptor:
  type: null-decryptor
items:
  - name: key1
    secret: value1
  - name: key2
    secret: value2
`)

	l, err := locker.MakeYAMLLocker(yamlContent)
	require.NoError(t, err)
	assert.NotNil(t, l)

	// Verify the items were loaded
	keys, err := l.List()
	require.NoError(t, err)
	assert.Len(t, keys, 2)
	assert.Contains(t, keys, "key1")
	assert.Contains(t, keys, "key2")
}

func TestMakeYAMLLocker_Get(t *testing.T) {
	yamlContent := []byte(`
decryptor:
  type: null-decryptor
items:
  - name: username
    secret: admin
  - name: password
    secret: secret123
`)

	l, err := locker.MakeYAMLLocker(yamlContent)
	require.NoError(t, err)

	// Test Get for existing keys
	value, err := l.Get("username")
	require.NoError(t, err)
	assert.Equal(t, "admin", value)

	value, err = l.Get("password")
	require.NoError(t, err)
	assert.Equal(t, "secret123", value)
}

func TestMakeYAMLLocker_GetNonExistentKey(t *testing.T) {
	yamlContent := []byte(`
decryptor:
  type: null-decryptor
items:
  - name: key1
    secret: value1
`)

	l, err := locker.MakeYAMLLocker(yamlContent)
	require.NoError(t, err)

	// Test Get for non-existent key
	_, err = l.Get("nonexistent")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "key not found")
}

func TestMakeYAMLLocker_EmptyItems(t *testing.T) {
	yamlContent := []byte(`
decryptor:
  type: null-decryptor
items: []
`)

	l, err := locker.MakeYAMLLocker(yamlContent)
	require.NoError(t, err)
	assert.NotNil(t, l)

	// Verify no items
	keys, err := l.List()
	require.NoError(t, err)
	assert.Len(t, keys, 0)
}

func TestMakeYAMLLocker_InvalidYAML(t *testing.T) {
	yamlContent := []byte(`
this is not valid yaml: [
`)

	_, err := locker.MakeYAMLLocker(yamlContent)
	assert.Error(t, err)
}

func TestMakeYAMLLocker_MissingDecryptor(t *testing.T) {
	yamlContent := []byte(`
items:
  - name: key1
    secret: value1
`)

	_, err := locker.MakeYAMLLocker(yamlContent)
	assert.Error(t, err)
}

func TestMakeYAMLLocker_InvalidDecryptorType(t *testing.T) {
	yamlContent := []byte(`
decryptor:
  type: invalid-decryptor-type
items:
  - name: key1
    secret: value1
`)

	_, err := locker.MakeYAMLLocker(yamlContent)
	assert.Error(t, err)
}

func TestMakeYAMLLocker_GetRaw(t *testing.T) {
	yamlContent := []byte(`
decryptor:
  type: null-decryptor
items:
  - name: apikey
    secret: sk-1234567890
`)

	l, err := locker.MakeYAMLLocker(yamlContent)
	require.NoError(t, err)

	// Test GetRaw
	value, err := l.GetRaw("apikey")
	require.NoError(t, err)
	assert.Equal(t, "sk-1234567890", value)
}

func TestMakeYAMLLocker_List(t *testing.T) {
	yamlContent := []byte(`
decryptor:
  type: null-decryptor
items:
  - name: alpha
    secret: value_alpha
  - name: beta
    secret: value_beta
  - name: gamma
    secret: value_gamma
`)

	l, err := locker.MakeYAMLLocker(yamlContent)
	require.NoError(t, err)

	keys, err := l.List()
	require.NoError(t, err)
	assert.Len(t, keys, 3)
	assert.Contains(t, keys, "alpha")
	assert.Contains(t, keys, "beta")
	assert.Contains(t, keys, "gamma")
}

func TestMakeYAMLLocker_DuplicateKeys(t *testing.T) {
	yamlContent := []byte(`
decryptor:
  type: null-decryptor
items:
  - name: duplicate
    secret: first_value
  - name: duplicate
    secret: second_value
`)

	l, err := locker.MakeYAMLLocker(yamlContent)
	require.NoError(t, err)

	// The second value should overwrite the first
	value, err := l.Get("duplicate")
	require.NoError(t, err)
	assert.Equal(t, "second_value", value)

	// List should only show one instance
	keys, err := l.List()
	require.NoError(t, err)
	assert.Len(t, keys, 1)
	assert.Contains(t, keys, "duplicate")
}

func TestMakeYAMLLocker_SpecialCharactersInValues(t *testing.T) {
	yamlContent := []byte(`
decryptor:
  type: null-decryptor
items:
  - name: special
    secret: "value with spaces and $pecial @characters!"
  - name: multiline
    secret: |
      line1
      line2
      line3
`)

	l, err := locker.MakeYAMLLocker(yamlContent)
	require.NoError(t, err)

	value, err := l.Get("special")
	require.NoError(t, err)
	assert.Equal(t, "value with spaces and $pecial @characters!", value)

	value, err = l.Get("multiline")
	require.NoError(t, err)
	assert.Contains(t, value, "line1")
	assert.Contains(t, value, "line2")
	assert.Contains(t, value, "line3")
}