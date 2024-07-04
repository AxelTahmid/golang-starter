package server

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"

	"github.com/AxelTahmid/golang-starter/config"
	"github.com/AxelTahmid/golang-starter/db"
)

type Server struct {
	conf   *config.Config
	router *chi.Mux
	db     *db.Postgres
	tls    *tls.Config
	log    *slog.Logger
}

func NewServer(c *config.Config, db *db.Postgres, t *tls.Config, l *slog.Logger) *Server {
	server := &Server{
		conf:   c,
		router: chi.NewRouter(),
		db:     db,
		tls:    t,
		log:    l,
	}

	server.routes()

	return server
}

func (s *Server) Start(ctx context.Context) {
	logger := slog.NewLogLogger(slog.NewJSONHandler(os.Stdout, nil), slog.LevelError)

	server := http.Server{
		Addr:         fmt.Sprintf(":%d", s.conf.Server.Port),
		Handler:      s.router,
		TLSConfig:    s.tls,
		IdleTimeout:  s.conf.Server.IdleTimeout,
		ReadTimeout:  s.conf.Server.ReadTimeout,
		WriteTimeout: s.conf.Server.WriteTimeout,
		ErrorLog:     logger,
	}

	shutdownComplete := handleShutdown(func() {

		log.Println("Starting server shutdown ...")

		s.db.Close()

		if err := server.Shutdown(ctx); err != nil {
			log.Printf("server.Shutdown failed: %v\n", err)
		}
	})

	log.Printf("Server started on port %d\n", s.conf.Server.Port)

	if err := server.ListenAndServeTLS("", ""); err == http.ErrServerClosed {
		<-shutdownComplete
	} else {
		log.Printf("http.ListenAndServe failed: %v\n", err)
	}

	log.Println("Server shutdown gracefully")
}

func handleShutdown(onShutdownSignal func()) <-chan struct{} {
	shutdown := make(chan struct{})

	go func() {
		shutdownSignal := make(chan os.Signal, 1)
		signal.Notify(shutdownSignal, os.Interrupt, syscall.SIGTERM)

		<-shutdownSignal

		onShutdownSignal()
		close(shutdown)
	}()

	return shutdown
}
