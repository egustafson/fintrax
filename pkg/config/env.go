package config

import "os"

func EnvOrDefault(key, defaultVal string) string {
	v, ok := os.LookupEnv(key)
	if !ok {
		v = defaultVal
	}
	return v
}
