package config

import "sync"

type Config struct {
	Server
	Cors
	Secure
	Database
}

var (
	once sync.Once
	conf *Config
)

func New() *Config {

	once.Do(func() {
		conf = &Config{
			Server:   serverConfig(),
			Cors:     corsConfig(),
			Secure:   secureConfig(),
			Database: dBConfig(),
		}
	},
	)

	return conf
}
