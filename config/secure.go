package config

import "github.com/kelseyhightower/envconfig"

type Secure struct {
	AllowedHosts          []string          `split_words:"true" default:"localhost"`
	AllowedHostsAreRegex  bool              `split_words:"true" default:"false"`
	HostsProxyHeaders     []string          `split_words:"true" default:"X-Forwarded-Host"`
	SSLRedirect           bool              `split_words:"true" default:"true"`
	SSLHost               string            `split_words:"true" default:"localhost"`
	STSSeconds            int64             `split_words:"true" default:"31536000"`
	STSIncludeSubdomains  bool              `split_words:"true" default:"true"`
	STSPreload            bool              `split_words:"true" default:"true"`
	FrameDeny             bool              `split_words:"true" default:"true"`
	ContentTypeNosniff    bool              `split_words:"true" default:"true"`
	BrowserXssFilter      bool              `split_words:"true" default:"true"`
	ContentSecurityPolicy string            `split_words:"true" default:"script-src $NONCE"`
	SSLProxyHeaders       map[string]string `split_words:"true" default:"X-Forwarded-Proto:https"`
}

func secureConfig() Secure {
	var s Secure
	envconfig.MustProcess("SECURE", &s)

	return s
}
