package server

import (
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"

	"github.com/AxelTahmid/golang-starter/internal/middlewares"
	"github.com/AxelTahmid/golang-starter/internal/modules/auth"
)

func (s *Server) routes() {

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	log := zerolog.New(os.Stderr).With().Timestamp().Logger()

	if s.conf.Env == "development" {
		log = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr}).With().Timestamp().Logger()
	}

	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.RequestID)
	s.router.Use(middlewares.Logger(log))
	s.router.Use(middleware.Recoverer)
	s.router.Use(middlewares.Helmet(s.conf.Secure).Handler)
	s.router.Use(middleware.Heartbeat("/ping"))

	// routes
	s.router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/auth", auth.Routes(s.db))
	})
}
