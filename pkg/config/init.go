package config

import "context"

func InitServerConfig(ctx context.Context, flags *Flags) (*ServerConfig, context.Context, error) {

	serverConfig := &ServerConfig{
		Flags: flags,
	}
	cfgPath := locateServerConfig(flags)
	err := loadConfigFromFile(cfgPath, &serverConfig)
	serverConfig.Flags = flags
	ctx = setServerConfig(ctx, serverConfig)
	return serverConfig, ctx, err
}
