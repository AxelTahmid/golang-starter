package config

import (
	"log"

	"github.com/joho/godotenv"
)

type Config struct {
	Api
	Cors
	Secure
}

func New() *Config {

	err := godotenv.Load()
	if err != nil {
		log.Println(err)
	}

	return &Config{
		Api:  ServerConfig(),
		Cors: CorsConfig(),
		Secure: SecureConfig(),
	}
}
