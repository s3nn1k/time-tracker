package postgres

import "os"

type Config struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
	SSLMode  string
}

func LoadConfig() Config {
	return Config{
		Host:     os.Getenv("HOST"),
		Port:     os.Getenv("P_DB_PORT"),
		User:     os.Getenv("P_DB_USER"),
		Password: os.Getenv("P_DB_PASS"),
		DBName:   os.Getenv("P_DB_NAME"),
		SSLMode:  os.Getenv("P_DB_SSL_MODE"),
	}
}
