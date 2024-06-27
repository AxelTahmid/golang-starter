package config

import "github.com/kelseyhightower/envconfig"

type Database struct {
	Url     string `required:"true"`
	PoolMax int32    `default:"5"`
	PoolMin int32    `default:"1"`
}

func DBConfig() Database {
	var d Database
	envconfig.MustProcess("DB", &d)

	return d
}
