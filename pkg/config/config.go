package config

var (
	serverConfig ServerConfig
)

func GetServerConfig() ServerConfig {
	return serverConfig
}

// DBConfig struct --> db.go

type ServerConfig struct {
	Flags Flags     `yaml:"-" json:"-"`
	Port  int       `yaml:"port" json:"port"`
	DB    *DBConfig `yaml:"db,omitempty" json:"db,omitempty"`
}

type Flags struct {
	Verbose bool
}
