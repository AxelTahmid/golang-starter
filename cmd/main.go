package main

import (
	"context"

	"github.com/AxelTahmid/golang-starter/api"
	"github.com/AxelTahmid/golang-starter/config"
)

func main() {
	ctx := context.Background()

	conf := config.New()

	server := api.NewServer(conf)
	server.Start(ctx)
}
