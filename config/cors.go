package config

import "github.com/kelseyhightower/envconfig"

type Cors struct {
	AllowedOrigins   []string `split_words:"true" required:"true"`
	AllowedMethods   []string `split_words:"true" default:"GET,POST,PUT,DELETE,PATCH,OPTIONS"`
	AllowedHeaders   []string `split_words:"true" default:"Origin,Authorization,User-Agent,Content-Type,Accept,Accept-Encoding,Accept-Language,Cache-Control,Connection,DNT,Host,Origin,Pragma,Referer"`
	AllowCredentials bool     `split_words:"true" default:"true"`
	MaxAge           int      `split_words:"true" default:"300"`
}

func corsConfig() Cors {
	var c Cors
	envconfig.MustProcess("CORS", &c)

	return c
}
