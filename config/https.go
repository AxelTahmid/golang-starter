package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

const envPrefix = "API"

type HTTPServer struct {
	Name         string        `default:"go-api"`
	Host         string        `default:"0.0.0.0"`
	Port         int           `default:"8080"`
	Logging      bool          `default:"true"`
	IdleTimeout  time.Duration `split_words:"true" default:"60s"`
	ReadTimeout  time.Duration `split_words:"true" default:"5s"`
	WriteTimeout time.Duration `split_words:"true" default:"10s"`
}

func NewServer() HTTPServer {
	var server HTTPServer

	envconfig.MustProcess("API", &server)

	return server
}
