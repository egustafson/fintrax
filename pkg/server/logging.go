package server

import (
	"log/slog"
	"os"
	"strconv"
	"strings"

	"github.com/egustafson/fintrax/pkg/config"
)

var (
	rootLogger *slog.Logger
)

func init() { // bootstrap logging: pre-config load
	logWr := os.Stderr
	levelStr := config.EnvOrDefault(config.ENV_LOG_LEVEL, config.DefaultLogLevel)

	rootLogger = slog.New(slog.NewTextHandler(logWr, &slog.HandlerOptions{
		Level: strToLevel(levelStr),
	}))
	slog.SetDefault(rootLogger)
	slog.Debug("logger initialized", "level", levelStr)
}

func strToLevel(levelStr string) slog.Level {
	if i, err := strconv.Atoi(levelStr); err == nil {
		return slog.Level(i)
	}
	switch strings.ToLower(levelStr) {
	case "debug":
		return slog.LevelDebug
	case "info":
		return slog.LevelInfo
	case "warn":
		return slog.LevelWarn
	case "error":
		return slog.LevelError
	default:
		return slog.LevelInfo
	}
}
