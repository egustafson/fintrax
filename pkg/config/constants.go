package config

import (
	"fmt"
	"log/slog"
)

// defaults
const (

	// Private
	defaultPort int = 8080
)

// defaults
var (
	DefaultLogLevel = fmt.Sprintf("%d", slog.LevelInfo)
)

// Registered env vars
const (
	ENV_PG_USER   string = "FINTRAX_PG_USER"
	ENV_PG_PASS   string = "FINTRAX_PG_PASS"
	ENV_PG_HOST   string = "FINTRAX_PG_HOST"
	ENV_LOG_LEVEL string = "FINTRAX_LOGLEVEL"
)
