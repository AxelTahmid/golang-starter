package config

import "github.com/kelseyhightower/envconfig"

type Jwt struct {
	JwtPubKeyPath string `split_words:"true" required:"true"`
	JwtPvtKeyPath string `split_words:"true" required:"true"`
}

func jwtConfig() Jwt {
	var j Jwt
	envconfig.MustProcess("", &j)
	return j
}
