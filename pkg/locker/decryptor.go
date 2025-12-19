package locker

import "fmt"

func MakeDecryptor(cfg DecryptorType) (Decryptor, error) {
	switch cfg.GetType() {
	case NullDecryptorType:
		// Handle both pointer and value types
		switch v := cfg.(type) {
		case NullDecryptorConfig:
			return MakeNullDecryptor(&v), nil
		case *NullDecryptorConfig:
			return MakeNullDecryptor(v), nil
		default:
			return nil, fmt.Errorf("invalid config for null-decryptor")
		}
	case PwAESDecryptorType:
		// Handle both pointer and value types
		switch v := cfg.(type) {
		case PwAESDecryptorConfig:
			return MakePwAESDecryptor(&v)
		case *PwAESDecryptorConfig:
			return MakePwAESDecryptor(v)
		default:
			return nil, fmt.Errorf("invalid config for pw-aes-decryptor")
		}
	case YubiKeyDecryptorType:
		// Handle both pointer and value types
		switch v := cfg.(type) {
		case YubiKeyDecryptorConfig:
			return MakeYubiKeyDecryptor(&v)
		case *YubiKeyDecryptorConfig:
			return MakeYubiKeyDecryptor(v)
		default:
			return nil, fmt.Errorf("invalid config for yk-decryptor")
		}
	default:
		return nil, ErrUnknownDecryptorType
	}
}
