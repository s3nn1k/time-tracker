package main

import (
	"log"

	"github.com/Time-Tracker/internal/service"
	"github.com/Time-Tracker/internal/storage/postgres"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

const (
	schemaPath = "./schema.sql"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	cfg := postgres.LoadConfig()

	db, err := postgres.NewDB(cfg)
	if err != nil {
		log.Fatal(err)
	}

	err = postgres.InitTables(schemaPath, db)
	if err != nil {
		log.Println(err)
	}

	strg := postgres.NewStorage(db)

	srvc := service.New(strg)

	token, err := srvc.Auth.GenerateToken(5)
	if err != nil {
		log.Fatal(err)
	}

	log.Println(token)

	// TODO: init handler
	// TODO: run server
	// TODO: add graceful shutdown
	// TODO: make Dockerfile for application
}
