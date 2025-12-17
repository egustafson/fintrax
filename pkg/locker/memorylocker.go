package locker

import "fmt"

type memoryLocker struct {
	decryptor Decryptor
	store     map[string]string
}

func MakeMemoryLocker(decryptor Decryptor) (Locker, error) {
	if decryptor == nil {
		return nil, fmt.Errorf("decryptor cannot be nil")
	}
	return &memoryLocker{
		decryptor: decryptor,
		store:     make(map[string]string),
	}, nil
}

// Get retrieves a value by name from the memory store
func (m *memoryLocker) Get(name string) (string, error) {
	// look up the value
	value, exists := m.store[name]
	if !exists {
		return "", fmt.Errorf("key not found: %s", name)
	}
	// decrypt the value
	decryptedValue, err := m.decryptor.Decrypt(value)
	if err != nil {
		return "", fmt.Errorf("failed to decrypt value for key %s: %w", name, err)
	}
	return decryptedValue, nil
}

// GetRaw retrieves a raw value by name from the memory store (same as Get for memory)
func (m *memoryLocker) GetRaw(name string) (string, error) {
	return m.Get(name)
}

// List returns all keys stored in the memory store
func (m *memoryLocker) List() ([]string, error) {
	keys := make([]string, 0, len(m.store))
	for key := range m.store {
		keys = append(keys, key)
	}
	return keys, nil
}

// Put stores a value with the given name in the memory store
func (m *memoryLocker) Put(name string, value string) error {
	// Encrypt the value
	encryptedValue, err := m.decryptor.Encrypt(value)
	if err != nil {
		return fmt.Errorf("failed to encrypt value for key %s: %w", name, err)
	}
	// Store the encrypted value
	m.store[name] = encryptedValue
	return nil
}

// Delete removes a value with the given name from the memory store
func (m *memoryLocker) Delete(name string) error {
	if _, exists := m.store[name]; !exists {
		return fmt.Errorf("key not found: %s", name)
	}
	delete(m.store, name)
	return nil
}
