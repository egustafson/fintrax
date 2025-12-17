package locker

import "fmt"

const (
	NullDecryptorType    = "null-decryptor"
	PwAESDecryptorType   = "pw-aes-decryptor"
	YubiKeyDecryptorType = "yk-decryptor"
)

var (
	ErrUnknownDecryptorType = fmt.Errorf("unknown decryptor type")
	ErrMissingDecryptorType = fmt.Errorf("missing decryptor type")
)

type DecryptorConfig struct {
	Decryptor DecryptorType `yaml:",inline"`
}

type DecryptorType interface {
	GetType() string
}

type NullDecryptorConfig struct{}

func (n NullDecryptorConfig) GetType() string {
	return NullDecryptorType
}

type PwAESDecryptorConfig struct {
	Password string `yaml:"password"`
}

func (p PwAESDecryptorConfig) GetType() string {
	return PwAESDecryptorType
}

type YubiKeyDecryptorConfig struct {
	Slot int    `yaml:"slot"`
	PIN  string `yaml:"pin"`
}

func (y YubiKeyDecryptorConfig) GetType() string {
	return YubiKeyDecryptorType
}

func (dc *DecryptorConfig) YAMLUNmarshal(unmarshal func(interface{}) error) error {
	var typeHolder struct {
		DecryptorType string `yaml:"type"`
	}
	if err := unmarshal(&typeHolder); err != nil {
		return err
	}
	if typeHolder.DecryptorType == "" {
		return ErrMissingDecryptorType
	}

	switch typeHolder.DecryptorType {
	case NullDecryptorType:
		var cfg NullDecryptorConfig
		if err := unmarshal(&cfg); err != nil {
			return err
		}
		dc.Decryptor = &cfg
	case PwAESDecryptorType:
		var cfg PwAESDecryptorConfig
		if err := unmarshal(&cfg); err != nil {
			return err
		}
		dc.Decryptor = &cfg
	case YubiKeyDecryptorType:
		var cfg YubiKeyDecryptorConfig
		if err := unmarshal(&cfg); err != nil {
			return err
		}
		dc.Decryptor = &cfg
	default:
		return ErrUnknownDecryptorType
	}
	return nil
}
