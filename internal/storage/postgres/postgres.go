package postgres

import (
	"database/sql"
	"fmt"
	"io"
	"os"

	"github.com/Time-Tracker/internal/storage"
)

const (
	timersTable = "timers"
	usersTable  = "users"
)

func NewDB(cfg Config) (db *sql.DB, err error) {
	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode)

	db, err = sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func InitTables(schemaPath string, db *sql.DB) (err error) {
	schemaFile, err := os.Open(schemaPath)
	if err != nil {
		return err
	}

	defer schemaFile.Close()

	data, err := io.ReadAll(schemaFile)
	if err != nil {
		return err
	}

	_, err = db.Exec(string(data))
	if err != nil {
		return err
	}

	return nil
}

func NewStorage(db *sql.DB) *storage.Storage {
	return &storage.Storage{
		Timers: newTimers(db),
		Users:  newUsers(db),
	}
}
