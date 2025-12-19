package locker

import (
	"gopkg.in/yaml.v3"
)

type LockerItem struct {
	Name   string `yaml:"name"`
	Secret string `yaml:"secret"`
}

type YamlLockerFile struct {
	Decryptor *DecryptorConfig `yaml:"decryptor"`
	Items     []LockerItem     `yaml:"items"`
}

type yamlLocker struct{ memoryLocker }

var _ Locker = (*yamlLocker)(nil)

// MakeYAMLLocker creates a new YAML-based locker from raw YAML.
func MakeYAMLLocker(content []byte) (Locker, error) {

	// Unmarshal the YAML content
	var yamlFile YamlLockerFile
	err := yaml.Unmarshal(content, &yamlFile)
	if err != nil {
		return nil, err
	}

	// Check if decryptor config is provided
	if yamlFile.Decryptor == nil || yamlFile.Decryptor.Decryptor == nil {
		return nil, ErrMissingDecryptorType
	}

	// Create a new yamlLocker
	decryptor, err := MakeDecryptor(yamlFile.Decryptor.Decryptor)
	if err != nil {
		return nil, err
	}
	locker := &yamlLocker{
		memoryLocker: memoryLocker{
			decryptor: decryptor,
			store:     make(map[string]string),
		},
	}

	for _, item := range yamlFile.Items {
		locker.memoryLocker.store[item.Name] = item.Secret
	}
	return locker, nil
}
