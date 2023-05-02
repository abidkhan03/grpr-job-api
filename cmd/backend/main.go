package main

import (
	"context"
	"log"

	"github.com/grpr-job-api/internal/api"
	"github.com/grpr-job-api/internal/repo"
	"github.com/grpr-job-api/internal/server"
	"github.com/grpr-job-api/internal/service"
	"github.com/grpr-job-api/shared"
)

func main() {
	_, cancel := context.WithCancel(context.Background())
	defer cancel()

	// config
	err := shared.LoadConfig(".env")
	if err != nil {
		log.Fatalf("error loading config %v", err)
	}
	cfg, err := readConfig()
	if err != nil {
		log.Fatal(err)
	}

	// database
	db, err := repo.New(shared.GetDBConnectionString())
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// service
	svc := service.New(db)

	// apis
	pingApi := api.NewPing(svc)
	fetcherApi := api.NewFetcher(svc)
	grouperApi := api.NewGrouper(svc)
	combinedApi := api.NewCombinedJob(svc)

	// server
	srv := server.New(cfg.HttpPort,
		pingApi,
		fetcherApi,
		grouperApi,
		combinedApi,
	)

	err = srv.Start()
	if err != nil {
		log.Fatalf("start server err: %v", err)
	}
}
