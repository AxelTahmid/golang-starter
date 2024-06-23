package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Api struct {
	Env          string        `default:"development"`
	Name         string        `default:"go-api"`
	Host         string        `default:"0.0.0.0"`
	Port         int           `default:"3000"`
	Logging      bool          `default:"true"`
	IdleTimeout  time.Duration `split_words:"true" default:"60s"`
	ReadTimeout  time.Duration `split_words:"true" default:"5s"`
	WriteTimeout time.Duration `split_words:"true" default:"10s"`
}

func ServerConfig() Api {
	var server Api

	envconfig.MustProcess("", &server)

	return server
}
