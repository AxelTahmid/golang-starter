package config

type Config struct {
	Server
	Cors
	Secure
	Database
}

func New() *Config {

	return &Config{
		Server:   ServerConfig(),
		Cors:     CorsConfig(),
		Secure:   SecureConfig(),
		Database: DBConfig(),
	}
}
