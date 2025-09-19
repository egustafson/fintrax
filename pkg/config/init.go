package config

import "context"

func InitServerConfig(ctx context.Context /*, flags Flags */) (*ServerConfig, context.Context, error) {

	serverConfig = ServerConfig{
		//Flags: flags,
		Port: defaultPort,
	}

	// TODO:  attach the server config to the context

	return &serverConfig, ctx, nil
}
