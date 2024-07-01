package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Database struct {
	Url               string        `required:"true"`
	SslMode           string        `default:"disable"`
	PoolMax           int32         `default:"5"`
	PoolMin           int32         `default:"0"`
	ConnectTimeout    time.Duration `default:"10s"`
	MaxConnLifetime   time.Duration `default:"60m"`
	MaxConnIdleTime   time.Duration `default:"30m"`
	HealthCheckPeriod time.Duration `default:"60s"`
}

func DBConfig() Database {
	var d Database
	envconfig.MustProcess("DB", &d)

	return d
}
