package server

import (
	"log/slog"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	"github.com/AxelTahmid/golang-starter/internal/middlewares"
	"github.com/AxelTahmid/golang-starter/internal/modules/auth"
)

func (s *Server) routes() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.RequestID)
	s.router.Use(middlewares.Logger(logger))
	s.router.Use(middleware.Recoverer)
	s.router.Use(middlewares.Helmet(s.conf.Secure).Handler)
	s.router.Use(middleware.Heartbeat("/ping"))

	// routes
	s.router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/auth", auth.Routes(s.db))
	})
}
