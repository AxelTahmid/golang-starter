package main

import (
	"context"

	"github.com/AxelTahmid/golang-starter/api"
	"github.com/AxelTahmid/golang-starter/config"
)

func main() {
	ctx := context.Background()

	cfg := config.New()

	server := api.NewServer(cfg.Api)
	server.Start(ctx)
}
