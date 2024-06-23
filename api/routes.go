package api

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func (s *Server) routes() {
	s.router.Use(render.SetContentType(render.ContentTypeJSON))

	s.router.Get("/health", s.handleGetHealth)

	s.router.Route("/api/v1", func(r chi.Router) {
		r.Get("/", s.handleGetHealth)
	})
}
