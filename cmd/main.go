package main

import (
	"context"
	"crypto/tls"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/AxelTahmid/golang-starter/config"
	"github.com/AxelTahmid/golang-starter/db"
	"github.com/AxelTahmid/golang-starter/internal/server"
	"github.com/lmittmann/tint"
)

func main() {
	ctx := context.Background()

	conf := config.New()

	var logger *slog.Logger

	if conf.AppEnv != "production" {
		logger = slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		)
		slog.SetDefault(logger)
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
		slog.SetDefault(logger)
	}

	serverTLSCert, err := tls.LoadX509KeyPair(conf.Server.TLSCertPath, conf.Server.TLSKeyPath)
	if err != nil {
		log.Fatalf("Error loading certificate and key file: %v", err)
	}

	tlsConfig := &tls.Config{
		Certificates: []tls.Certificate{serverTLSCert},
	}

	dbconn, err := db.CreatePool(ctx, conf.Database)
	if err != nil {
		log.Fatalf("Db Connection Failed: %v", err)
	}

	err = dbconn.Ping(ctx)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	server := server.NewServer(conf, dbconn, tlsConfig, logger)
	server.Start(ctx)
}
