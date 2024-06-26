package main

import (
	"context"
	"log"
	"os"

	"github.com/AxelTahmid/golang-starter/api"
	"github.com/AxelTahmid/golang-starter/config"
	"github.com/AxelTahmid/golang-starter/db"
)

func main() {
	ctx := context.Background()

	conf := config.New()

	dbconn, err := db.ConnectDB(context.Background(), conf.Database.Url)

	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	server := api.NewServer(conf, dbconn)
	server.Start(ctx)
}
