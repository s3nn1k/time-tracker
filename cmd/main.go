package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/Time-Tracker/cmd/config"
	"github.com/Time-Tracker/internal/auth/jwt"
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
	} else {
		log.Println("Tables created successfully")
	}

	strg := postgres.NewStorage(db)

	srvc := service.New(strg)

	auth := jwt.NewAuth(cfg.Salt, cfg.JWTTTL, cfg.JWTKey)

	h := handler.New(srvc, auth)

	r := h.InitRoutes()

	addr := fmt.Sprintf("%s:%s", cfg.Host, cfg.ServerPort)

	srvr := http.Server{
		Addr:           addr,
		MaxHeaderBytes: 1 << 20,
		Handler:        r,
		WriteTimeout:   cfg.WriteTimeout,
		ReadTimeout:    cfg.ReadTimeout,
	}

	go func() {
		if err := srvr.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	log.Println("Service started successfully")

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down service...")

	err = srvr.Shutdown(context.Background())
	if err != nil {
		log.Println(err)
	}

	err = db.Close()
	if err != nil {
		log.Println(err)
	}

	log.Println("Service shutdown successfully")
}
