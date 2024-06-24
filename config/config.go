package config

type Config struct {
	Api
	Cors
	Secure
}

func New() *Config {

	// load .env either by Makefile or Docker Compose

	return &Config{
		Api:    ServerConfig(),
		Cors:   CorsConfig(),
		Secure: SecureConfig(),
	}
}
