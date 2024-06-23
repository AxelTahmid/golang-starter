package api

import (
	"os"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
	"github.com/rs/zerolog"
	"github.com/unrolled/secure"

	"github.com/AxelTahmid/golang-starter/api/middlewares"
)

func (s *Server) routes() {

	// global middlewares
	secureMiddleware := secure.New(secure.Options{
		FrameDeny:          s.conf.Secure.FrameDeny,
		ContentTypeNosniff: s.conf.Secure.ContentTypeNosniff,
		BrowserXssFilter:   s.conf.Secure.BrowserXssFilter,
	})

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	loggerOutput := zerolog.ConsoleWriter{Out: os.Stderr}
	log := zerolog.New(loggerOutput)

	if s.conf.Env != "development" {
		log = zerolog.New(os.Stderr)
	}

	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.RequestID)
	s.router.Use(middlewares.Logger(log))
	s.router.Use(middleware.Recoverer)

	s.router.Use(secureMiddleware.Handler)

	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	// routes
	s.router.Get("/health", s.handleGetHealth)

}
