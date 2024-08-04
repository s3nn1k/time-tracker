package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	DBHost     string //msut be the same as postgres container name in docker-compose file
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	ServerHost   string
	ServerPort   string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration

	Salt   string
	JWTTTL time.Duration
	JWTKey string
}

func Load() (Config, error) {
	jwtTTL, err := time.ParseDuration(os.Getenv("JWT_TTL"))
	if err != nil {
		return Config{}, fmt.Errorf("wrong JWT_TTL value format: %s", err)
	}

	writeTimeout, err := time.ParseDuration(os.Getenv("WRITE_TIMEOUT"))
	if err != nil {
		return Config{}, fmt.Errorf("wrong WRITE_TIMEOUT value format: %s", err)
	}

	readTimeout, err := time.ParseDuration(os.Getenv("READ_TIMEOUT"))
	if err != nil {
		return Config{}, fmt.Errorf("wrong READ_TIMEOUT value format: %s", err)
	}

	return Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("P_DB_PORT"),
		DBUser:     os.Getenv("P_DB_USER"),
		DBPassword: os.Getenv("P_DB_PASS"),
		DBName:     os.Getenv("P_DB_NAME"),
		DBSSLMode:  os.Getenv("P_DB_SSL_MODE"),

		ServerHost:   os.Getenv("SERVER_HOST"),
		ServerPort:   os.Getenv("SERVER_PORT"),
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,

		Salt:   os.Getenv("PW_SALT"),
		JWTTTL: jwtTTL,
		JWTKey: os.Getenv("JWT_KEY"),
	}, nil
}
