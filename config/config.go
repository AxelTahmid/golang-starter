package config

type Config struct {
	Api
	Cors
}

func New() *Config {

	return &Config{
		Api:           ServerConfig(),
		Cors:          CorsConfig(),
	}
}