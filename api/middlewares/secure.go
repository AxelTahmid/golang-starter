package middlewares

import (
	"github.com/AxelTahmid/golang-starter/config"
	"github.com/unrolled/secure"
)

func Helmet(conf config.Secure) *secure.Secure {
	secureMiddleware := secure.New(secure.Options{
		// FrameDeny:          	  conf.FrameDeny,
		// ContentTypeNosniff: 	  conf.ContentTypeNosniff,
		// BrowserXssFilter:   	  conf.BrowserXssFilter,
		// AllowedHosts:          conf.AllowedHosts,
		// AllowedHostsAreRegex:  conf.AllowedHostsAreRegex,
		// HostsProxyHeaders:     conf.HostsProxyHeaders,
		// SSLRedirect:           conf.SSLRedirect,
		// SSLHost:               conf.SSLHost,
		// SSLProxyHeaders:       conf.SSLProxyHeaders,
		STSSeconds:            conf.STSSeconds,
		STSIncludeSubdomains:  conf.STSIncludeSubdomains,
		STSPreload:            conf.STSPreload,
		FrameDeny:             conf.FrameDeny,
		ContentTypeNosniff:    conf.ContentTypeNosniff,
		BrowserXssFilter:      conf.BrowserXssFilter,
		ContentSecurityPolicy: conf.ContentSecurityPolicy,
	})

	return secureMiddleware
}
