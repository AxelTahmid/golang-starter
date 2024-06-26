package api

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"

	"github.com/AxelTahmid/golang-starter/config"
)

type Server struct {
	conf   config.Config
	router *chi.Mux
}

func NewServer(conf *config.Config) *Server {
	server := &Server{
		conf:   *conf,
		router: chi.NewRouter(),
	}

	server.routes()

	return server
}

func (s *Server) Start(ctx context.Context) {
	server := http.Server{
		Addr:         fmt.Sprintf(":%d", s.conf.Server.Port),
		Handler:      s.router,
		IdleTimeout:  s.conf.Server.IdleTimeout,
		ReadTimeout:  s.conf.Server.ReadTimeout,
		WriteTimeout: s.conf.Server.WriteTimeout,
	}

	shutdownComplete := handleShutdown(func() {
		if err := server.Shutdown(ctx); err != nil {
			log.Printf("server.Shutdown failed: %v\n", err)
		}
	})

	log.Printf("Server started on port %d\n", s.conf.Server.Port)

	if err := server.ListenAndServe(); err == http.ErrServerClosed {
		<-shutdownComplete
	} else {
		log.Printf("http.ListenAndServe failed: %v\n", err)
	}

	log.Println("Shutdown gracefully")
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
