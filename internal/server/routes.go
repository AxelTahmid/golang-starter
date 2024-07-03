package server

import (
	"log/slog"
	"os"
	"time"

	"github.com/go-chi/chi/v5"
	chiMiddleware "github.com/go-chi/chi/v5/middleware"
	"github.com/lmittmann/tint"

	"github.com/AxelTahmid/golang-starter/internal/middlewares"
	"github.com/AxelTahmid/golang-starter/internal/modules/auth"
)

func (s *Server) routes() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	if s.conf.AppEnv != "production" {
		logger = slog.New(tint.NewHandler(os.Stderr, nil))

		slog.SetDefault(slog.New(
			tint.NewHandler(os.Stderr, &tint.Options{
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		))
	}

	s.router.Use(chiMiddleware.RealIP)
	s.router.Use(chiMiddleware.RequestID)
	s.router.Use(middlewares.Logger(logger))
	s.router.Use(middlewares.Recovery)
	s.router.Use(middlewares.Helmet(s.conf.Secure).Handler)
	s.router.Use(middlewares.Json)
	s.router.Use(chiMiddleware.Heartbeat("/ping"))

	// routes
	s.router.Route("/api/v1", func(r chi.Router) {
		r.Mount("/auth", auth.Routes(s.db))
	})
}
