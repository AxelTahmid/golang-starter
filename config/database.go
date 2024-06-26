package config

import "github.com/kelseyhightower/envconfig"

type Database struct {
	Url string `required:"true"`
}

func DBConfig() Database {
	var d Database
	envconfig.MustProcess("DB", &d)

	return d
}
