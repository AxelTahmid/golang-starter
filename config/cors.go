package config

import "github.com/kelseyhightower/envconfig"


type Cors struct {
	AllowedOrigins []string `split_words:"true"`
	AllowedMethods []string `split_words:"true"`
	AllowedHeaders []string `split_words:"true"`
	AllowCredentials bool `yaml:"credentials"`
	MaxAge int `yaml:"max_age"`
}

func NewCors() Cors {
	var c Cors
	envconfig.MustProcess("CORS", &c)

	return c
}