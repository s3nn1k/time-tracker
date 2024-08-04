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

func NewDB(host, port, user, password, dbName, sslMode string) (db *sql.DB, err error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", user, password, host, port, dbName, sslMode)

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
