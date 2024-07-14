package main

import (
	"context"
	"log"
	"log/slog"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"github.com/AxelTahmid/golang-starter/config"
	"github.com/AxelTahmid/golang-starter/db"
	"github.com/AxelTahmid/golang-starter/internal/server"
	"github.com/AxelTahmid/golang-starter/internal/utils/jwt"
)

func main() {
	ctx := context.Background()

	// Load configuration
	conf := config.New()

	// Setup logger
	var logger *slog.Logger

	if conf.AppEnv != "production" {
		logger = slog.New(
			tint.NewHandler(os.Stdout, &tint.Options{
				Level:      slog.LevelDebug,
				TimeFormat: time.Kitchen,
			}),
		)
	} else {
		logger = slog.New(slog.NewJSONHandler(os.Stdout, nil))
	}
	// Set default logger
	slog.SetDefault(logger)

	// Connect to database
	dbconn, err := db.CreatePool(ctx, conf.Database, logger)
	if err != nil {
		log.Fatalf("Db Connection Failed: %v", err)
	}

	// Check if database connection is successful
	err = dbconn.Ping(ctx)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Set default jwt configuration
	jwt.SetDefaults(conf.Jwt)

	// Start server
	server := server.NewServer(conf, dbconn, logger)
	server.Start(ctx)
}
