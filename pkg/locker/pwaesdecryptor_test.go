package locker_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/egustafson/fintrax/pkg/locker"
)

func TestMakePwAESDecryptor_Success(t *testing.T) {
	config := &locker.PwAESDecryptorConfig{
		Password: "testpassword123",
	}

	dec, err := locker.MakePwAESDecryptor(config)
	require.NoError(t, err)
	assert.NotNil(t, dec)
}

func TestMakePwAESDecryptor_EmptyPassword(t *testing.T) {
	config := &locker.PwAESDecryptorConfig{
		Password: "",
	}

	dec, err := locker.MakePwAESDecryptor(config)
	assert.Error(t, err)
	assert.Nil(t, dec)
	assert.Contains(t, err.Error(), "password cannot be empty")
}

func TestMakePwAESDecryptor_ShortPassword(t *testing.T) {
	config := &locker.PwAESDecryptorConfig{
		Password: "short",
	}

	dec, err := locker.MakePwAESDecryptor(config)
	assert.Error(t, err)
	assert.Nil(t, dec)
	assert.Contains(t, err.Error(), "at least")
}

func TestPwAESDecryptor_EncryptDecrypt_RoundTrip(t *testing.T) {
	config := &locker.PwAESDecryptorConfig{
		Password: "testpassword123",
	}

	dec, err := locker.MakePwAESDecryptor(config)
	require.NoError(t, err)

	testCases := []struct {
		name      string
		plaintext string
	}{
		{"simple text", "Hello, World!"},
		{"empty string", ""},
		{"special chars", "!@#$%^&*()_+-=[]{}|;:',.<>?"},
		{"unicode", "Hello 世界 🌍"},
		{"multiline", "Line 1\nLine 2\nLine 3"},
		{"long text", "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Sed do eiusmod tempor incididunt ut labore et dolore magna aliqua."},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Encrypt
			encrypted, err := dec.Encrypt(tc.plaintext)
			require.NoError(t, err)
			assert.NotEmpty(t, encrypted)
			assert.NotEqual(t, tc.plaintext, encrypted)

			// Decrypt
			decrypted, err := dec.Decrypt(encrypted)
			require.NoError(t, err)
			assert.Equal(t, tc.plaintext, decrypted)
		})
	}
}

func TestPwAESDecryptor_Encrypt_ProducesUniqueOutput(t *testing.T) {
	config := &locker.PwAESDecryptorConfig{
		Password: "testpassword123",
	}

	dec, err := locker.MakePwAESDecryptor(config)
	require.NoError(t, err)

	plaintext := "same plaintext"

	// Encrypt the same plaintext multiple times
	encrypted1, err := dec.Encrypt(plaintext)
	require.NoError(t, err)

	encrypted2, err := dec.Encrypt(plaintext)
	require.NoError(t, err)

	encrypted3, err := dec.Encrypt(plaintext)
	require.NoError(t, err)

	// Each encryption should produce different ciphertext (due to random nonce)
	assert.NotEqual(t, encrypted1, encrypted2)
	assert.NotEqual(t, encrypted2, encrypted3)
	assert.NotEqual(t, encrypted1, encrypted3)

	// But all should decrypt to the same plaintext
	decrypted1, err := dec.Decrypt(encrypted1)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted1)

	decrypted2, err := dec.Decrypt(encrypted2)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted2)

	decrypted3, err := dec.Decrypt(encrypted3)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted3)
}

func TestPwAESDecryptor_DifferentPasswords_ProduceDifferentKeys(t *testing.T) {
	plaintext := "secret data"

	config1 := &locker.PwAESDecryptorConfig{
		Password: "password123",
	}
	dec1, err := locker.MakePwAESDecryptor(config1)
	require.NoError(t, err)

	config2 := &locker.PwAESDecryptorConfig{
		Password: "differentpassword456",
	}
	dec2, err := locker.MakePwAESDecryptor(config2)
	require.NoError(t, err)

	// Encrypt with first password
	encrypted1, err := dec1.Encrypt(plaintext)
	require.NoError(t, err)

	// Try to decrypt with second password (should fail)
	_, err = dec2.Decrypt(encrypted1)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "decryption failed")

	// Encrypt with second password
	encrypted2, err := dec2.Encrypt(plaintext)
	require.NoError(t, err)

	// Verify ciphertexts are different
	assert.NotEqual(t, encrypted1, encrypted2)

	// Decrypt with correct passwords
	decrypted1, err := dec1.Decrypt(encrypted1)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted1)

	decrypted2, err := dec2.Decrypt(encrypted2)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted2)
}

func TestPwAESDecryptor_SamePassword_ProducesSameKey(t *testing.T) {
	plaintext := "consistent data"
	password := "samepassword123"

	config1 := &locker.PwAESDecryptorConfig{
		Password: password,
	}
	dec1, err := locker.MakePwAESDecryptor(config1)
	require.NoError(t, err)

	config2 := &locker.PwAESDecryptorConfig{
		Password: password,
	}
	dec2, err := locker.MakePwAESDecryptor(config2)
	require.NoError(t, err)

	// Encrypt with first instance
	encrypted, err := dec1.Encrypt(plaintext)
	require.NoError(t, err)

	// Decrypt with second instance (same password)
	decrypted, err := dec2.Decrypt(encrypted)
	require.NoError(t, err)
	assert.Equal(t, plaintext, decrypted)
}

func TestPwAESDecryptor_Decrypt_InvalidBase64(t *testing.T) {
	config := &locker.PwAESDecryptorConfig{
		Password: "testpassword123",
	}

	dec, err := locker.MakePwAESDecryptor(config)
	require.NoError(t, err)

	// Invalid base64
	_, err = dec.Decrypt("not-valid-base64!!!")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "failed to decode base64")
}

func TestPwAESDecryptor_Decrypt_TooShort(t *testing.T) {
	config := &locker.PwAESDecryptorConfig{
		Password: "testpassword123",
	}

	dec, err := locker.MakePwAESDecryptor(config)
	require.NoError(t, err)

	// Valid base64 but too short to contain nonce
	_, err = dec.Decrypt("YWJj") // "abc" in base64
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "ciphertext too short")
}

func TestPwAESDecryptor_Decrypt_CorruptedData(t *testing.T) {
	config := &locker.PwAESDecryptorConfig{
		Password: "testpassword123",
	}

	dec, err := locker.MakePwAESDecryptor(config)
	require.NoError(t, err)

	// Encrypt some data
	plaintext := "original data"
	encrypted, err := dec.Encrypt(plaintext)
	require.NoError(t, err)

	// Corrupt the encrypted data by changing a character
	corrupted := encrypted[:len(encrypted)-5] + "XXXXX"

	// Try to decrypt corrupted data
	_, err = dec.Decrypt(corrupted)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "decryption failed")
}

func TestPwAESDecryptor_Decrypt_RandomData(t *testing.T) {
	config := &locker.PwAESDecryptorConfig{
		Password: "testpassword123",
	}

	dec, err := locker.MakePwAESDecryptor(config)
	require.NoError(t, err)

	// Valid base64 but random data
	_, err = dec.Decrypt("SGVsbG8gV29ybGQhIFRoaXMgaXMgbm90IGVuY3J5cHRlZCBkYXRhLg==")
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "decryption failed")
}
