package main

import (
	"log"
	"net/http"

	jwtoken "github.com/Time-Tracker/internal/auth/jwt"
	"github.com/Time-Tracker/internal/config"
	"github.com/Time-Tracker/internal/handler"
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
	cfg, err := config.Load()
	if err != nil {
		log.Fatal(err)
	}

	db, err := postgres.NewDB(cfg.Host, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName, cfg.DBSSLMode)
	if err != nil {
		log.Fatal(err)
	}

	err = postgres.InitTables(schemaPath, db)
	if err != nil {
		log.Println(err)
	}

	strg := postgres.NewStorage(db)

	srvc := service.New(strg)

	auth := jwtoken.NewAuth(cfg.Salt, cfg.JWTTTL, cfg.JWTKey)

	h := handler.New(srvc, auth)

	r := h.InitRoutes()

	s := http.Server{
		Addr:    "localhost:8080",
		Handler: r,
	}

	s.ListenAndServe()

	// TODO: add graceful shutdown
	// TODO: make Dockerfile for application
}
