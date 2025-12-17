package locker

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"

	"golang.org/x/crypto/pbkdf2"
)

const (
	// PBKDF2 iterations - higher is more secure but slower
	// 600,000 iterations is OWASP recommendation as of 2023 for PBKDF2-SHA256
	pbkdf2Iterations = 600000
	// AES-256 requires 32 bytes
	aesKeySize = 32
	// Salt size in bytes
	saltSize = 16
	// Minimum password length
	minPasswordLength = 8
)

type PwAESDecryptor struct {
	config *PwAESDecryptorConfig
	key    []byte
}

func MakePwAESDecryptor(config *PwAESDecryptorConfig) (*PwAESDecryptor, error) {
	if config.Password == "" {
		return nil, fmt.Errorf("password cannot be empty")
	}
	if len(config.Password) < minPasswordLength {
		return nil, fmt.Errorf("password must be at least %d characters long", minPasswordLength)
	}

	// Derive AES key from password using PBKDF2
	// Using a fixed salt derived from a constant - this makes keys deterministic
	// In production, you may want to store salt separately or derive it differently
	salt := sha256.Sum256([]byte("fintrax-aes-salt-v1"))

	key := pbkdf2.Key(
		[]byte(config.Password),
		salt[:saltSize],
		pbkdf2Iterations,
		aesKeySize,
		sha256.New,
	)

	return &PwAESDecryptor{
		config: config,
		key:    key,
	}, nil
}

func (p PwAESDecryptor) Decrypt(data string) (string, error) {
	// Decode base64 encoded data
	ciphertext, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", fmt.Errorf("failed to decode base64 data: %w", err)
	}

	// Create AES cipher block
	block, err := aes.NewCipher(p.key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	// Extract nonce and ciphertext
	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]

	// Decrypt and verify
	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", fmt.Errorf("decryption failed: %w", err)
	}

	return string(plaintext), nil
}

func (p PwAESDecryptor) Encrypt(data string) (string, error) {
	// Create AES cipher block
	block, err := aes.NewCipher(p.key)
	if err != nil {
		return "", fmt.Errorf("failed to create cipher: %w", err)
	}

	// Create GCM mode
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", fmt.Errorf("failed to create GCM: %w", err)
	}

	// Generate random nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err := rand.Read(nonce); err != nil {
		return "", fmt.Errorf("failed to generate nonce: %w", err)
	}

	// Encrypt and authenticate
	ciphertext := gcm.Seal(nonce, nonce, []byte(data), nil)

	// Encode as base64
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}
