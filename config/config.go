package config

type Config struct {
	NewServer
	Cors
}

func New() *Config {

	return &Config{
		Api:           NewServer(),
		Cors:          NewCors(),
	}
}