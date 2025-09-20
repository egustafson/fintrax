package config

import (
	"github.com/a8m/envsubst"
	"gopkg.in/yaml.v3"
)

func loadConfigFromFile(path string, config any) error {
	bytes, err := envsubst.ReadFile(path)
	if err != nil {
		return err
	}
	if err = yaml.Unmarshal(bytes, config); err != nil {
		return err
	}
	return nil
}

func locateServerConfig() string {
	return EnvOrDefault(ENV_CFG_FILE, defaultServerConfigFile)
}
