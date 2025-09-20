package config

import "context"

func InitServerConfig(ctx context.Context /*, flags Flags */) (*ServerConfig, context.Context, error) {

	cfgPath := locateServerConfig()
	err := loadConfigFromFile(cfgPath, &serverConfig)

	// TODO:  attach the server config to the context

	return &serverConfig, ctx, err
}
