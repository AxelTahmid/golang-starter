package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Server struct {
	AppEnv       string        `default:"development"`
	Name         string        `default:"go-api"`
	Host         string        `default:"0.0.0.0"`
	Port         int           `default:"3000"`
	Logging      bool          `default:"true"`
	IdleTimeout  time.Duration `split_words:"true" default:"60s"`
	ReadTimeout  time.Duration `split_words:"true" default:"5s"`
	WriteTimeout time.Duration `split_words:"true" default:"10s"`
	TLSCertPath  string        `split_words:"true" default:"./cert/tls.crt"`
	TLSKeyPath   string        `split_words:"true" default:"./cert/tls.key"`
}

func serverConfig() Server {
	var server Server

	envconfig.MustProcess("", &server)

	return server
}
