package api

import (
	"github.com/go-chi/chi/v5/middleware"
	"github.com/unrolled/secure"
	"github.com/rs/zerolog"
    "github.com/rs/zerolog/log"
)

func (s *Server) middlewares() {
	secureMiddleware := secure.New(secure.Options{
		FrameDeny:          s.conf.Secure.FrameDeny,
		ContentTypeNosniff: s.conf.Secure.ContentTypeNosniff,
		BrowserXssFilter:   s.conf.Secure.BrowserXssFilter,
	})

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger = 

	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)

	s.router.Use(secureMiddleware.Handler)
}
