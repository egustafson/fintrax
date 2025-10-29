package config

import (
	"context"
	"fmt"
	"reflect"
)

type Flags struct {
	Verbose bool
}

// DBConfig struct --> db.go

type ServerConfig struct {
	Flags *Flags    `yaml:"-" json:"-"`
	Port  int       `yaml:"port" json:"port"`
	DB    *DBConfig `yaml:"db,omitempty" json:"db,omitempty"`
}

var serverConfigFullyQualifiedName = fmt.Sprintf("%s.%s",
	reflect.TypeOf(ServerConfig{}).PkgPath(),
	reflect.TypeOf(ServerConfig{}).Name(),
)

func GetServerConfig(ctx context.Context) *ServerConfig {
	return ctx.Value(serverConfigFullyQualifiedName).(*ServerConfig)
}

func setServerConfig(ctx context.Context, cfg *ServerConfig) context.Context {
	return context.WithValue(ctx, serverConfigFullyQualifiedName, cfg)
}
