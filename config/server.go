package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Server struct {
	Env          string        `default:"development"`
	Name         string        `default:"go-api"`
	Host         string        `default:"0.0.0.0"`
	Port         int           `default:"3000"`
	Logging      bool          `default:"true"`
	IdleTimeout  time.Duration `split_words:"true" default:"60s"`
	ReadTimeout  time.Duration `split_words:"true" default:"5s"`
	WriteTimeout time.Duration `split_words:"true" default:"10s"`
}

func ServerConfig() Server {
	var server Server

	envconfig.MustProcess("", &server)

	return server
}
