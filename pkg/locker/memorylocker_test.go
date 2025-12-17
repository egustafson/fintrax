package locker_test

import (
	"testing"

	"github.com/egustafson/fintrax/pkg/locker"
	"github.com/stretchr/testify/assert"
)

func TestMemoryLocker_GetAndGetRaw(t *testing.T) {
	decryptor := locker.NullDecryptor{}
	ml, err := locker.MakeMemoryLocker(decryptor)
	assert.NoError(t, err)

	// Initially, getting a non-existent key should return an error
	_, err = ml.Get("test-key")
	assert.Error(t, err)

	_, err = ml.GetRaw("test-key")
	assert.Error(t, err)
}

func TestMemoryLocker_List_Empty(t *testing.T) {
	decryptor := locker.NullDecryptor{}
	ml, err := locker.MakeMemoryLocker(decryptor)
	assert.NoError(t, err)

	keys, err := ml.List()
	assert.NoError(t, err)
	assert.Empty(t, keys)
}

func TestMemoryLocker_Put(t *testing.T) {
	decryptor := locker.NullDecryptor{}
	ml, err := locker.MakeMemoryLocker(decryptor)
	assert.NoError(t, err)

	// Cast to LockerWr to access Put method
	mlWr, ok := ml.(locker.LockerWr)
	assert.True(t, ok)

	// Put a value
	err = mlWr.Put("test-key", "test-value")
	assert.NoError(t, err)

	// Verify we can retrieve it
	value, err := ml.Get("test-key")
	assert.NoError(t, err)
	assert.Equal(t, "test-value", value)

	// Put another value with the same key (should overwrite)
	err = mlWr.Put("test-key", "new-value")
	assert.NoError(t, err)

	value, err = ml.Get("test-key")
	assert.NoError(t, err)
	assert.Equal(t, "new-value", value)
}

func TestMemoryLocker_Delete(t *testing.T) {
	decryptor := locker.NullDecryptor{}
	ml, err := locker.MakeMemoryLocker(decryptor)
	assert.NoError(t, err)

	mlWr, ok := ml.(locker.LockerWr)
	assert.True(t, ok)

	// Put a value
	err = mlWr.Put("test-key", "test-value")
	assert.NoError(t, err)

	// Delete it
	err = mlWr.Delete("test-key")
	assert.NoError(t, err)

	// Verify it's gone
	_, err = ml.Get("test-key")
	assert.Error(t, err)

	// Try to delete non-existent key
	err = mlWr.Delete("non-existent")
	assert.Error(t, err)
}

func TestMemoryLocker_List(t *testing.T) {
	decryptor := locker.NullDecryptor{}
	ml, err := locker.MakeMemoryLocker(decryptor)
	assert.NoError(t, err)

	mlWr, ok := ml.(locker.LockerWr)
	assert.True(t, ok)

	// Put several values
	err = mlWr.Put("key1", "value1")
	assert.NoError(t, err)
	err = mlWr.Put("key2", "value2")
	assert.NoError(t, err)
	err = mlWr.Put("key3", "value3")
	assert.NoError(t, err)

	// List should return all keys
	keys, err := ml.List()
	assert.NoError(t, err)
	assert.Len(t, keys, 3)
	assert.Contains(t, keys, "key1")
	assert.Contains(t, keys, "key2")
	assert.Contains(t, keys, "key3")
}

func TestMemoryLocker_MakeMemoryLocker_NilDecryptor(t *testing.T) {
	_, err := locker.MakeMemoryLocker(nil)
	assert.Error(t, err)
}
