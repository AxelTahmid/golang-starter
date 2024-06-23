package config

import "github.com/kelseyhightower/envconfig"

type Cors struct {
	AllowedOrigins   []string `split_words:"true"`
	AllowedMethods   []string `split_words:"true" default:"GET,POST,PUT,DELETE,PATCH,OPTIONS"`
	AllowedHeaders   []string `split_words:"true" default:"Origin,Content-Type,Accept,Authorization"`
	AllowCredentials bool     `split_words:"true" default:"true"`
	MaxAge           int      `split_words:"true" default:"300"`
}

func CorsConfig() Cors {
	var c Cors
	envconfig.MustProcess("CORS", &c)

	return c
}
