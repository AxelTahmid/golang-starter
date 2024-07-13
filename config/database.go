package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Database struct {
	Url               string        `required:"true"`
	SslMode           string        `default:"disable"`
	PoolMax           int32         `split_words:"true" default:"5"`
	PoolMin           int32         `split_words:"true" default:"0"`
	ConnectTimeout    time.Duration `split_words:"true" default:"10s"`
	MaxConnLifetime   time.Duration `split_words:"true" default:"60m"`
	MaxConnIdleTime   time.Duration `split_words:"true" default:"30m"`
	HealthCheckPeriod time.Duration `split_words:"true" default:"60s"`
	TimeZone          string        `split_words:"true" default:"UTC"`
}

func dBConfig() Database {
	var d Database
	envconfig.MustProcess("DB", &d)

	return d
}
