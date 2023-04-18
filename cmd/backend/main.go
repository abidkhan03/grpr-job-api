package main

import (
	"context"
	"log"

	"github.com/grpr-job-api/internal/api"
	"github.com/grpr-job-api/internal/repo"
	"github.com/grpr-job-api/internal/server"
	"github.com/grpr-job-api/internal/service"
	"github.com/grpr-job-api/shared"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
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
	dbCfg, err := pgxpool.ParseConfig(shared.GetDBConnectionString())
	if err != nil {
		log.Fatal(err)
	}

	dbCfg.ConnConfig.LogLevel = pgx.LogLevelError

	conn, err := pgxpool.ConnectConfig(ctx, dbCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	db := repo.New(conn)

	// service
	svc := service.New(db)

	// apis
	pingApi := api.NewPing(svc)

	// server
	srv := server.New(cfg.HttpPort,
		pingApi,
	)

	err = srv.Start()
	if err != nil {
		log.Fatalf("start server err: %v", err)
	}
}
