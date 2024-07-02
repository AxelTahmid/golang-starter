package config

type Config struct {
	Server
	Cors
	Secure
	Database
}

func New() *Config {

	return &Config{
		Server:   serverConfig(),
		Cors:     corsConfig(),
		Secure:   secureConfig(),
		Database: dBConfig(),
	}
}
