package config

import (
	"fmt"
	"os"
	"time"
)

type Config struct {
	Host         string
	DBPort       string
	DBUser       string
	DBPassword   string
	DBName       string
	DBSSLMode    string
	ServerPort   string
	Salt         string
	JWTTTL       time.Duration
	JWTKey       string
	WriteTimeout time.Duration
	ReadTimeout  time.Duration
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
		Host:         os.Getenv("HOST"),
		DBPort:       os.Getenv("P_DB_PORT"),
		DBUser:       os.Getenv("P_DB_USER"),
		DBPassword:   os.Getenv("P_DB_PASS"),
		DBName:       os.Getenv("P_DB_NAME"),
		DBSSLMode:    os.Getenv("P_DB_SSL_MODE"),
		ServerPort:   os.Getenv("SERVER_PORT"),
		Salt:         os.Getenv("PW_SALT"),
		JWTTTL:       jwtTTL,
		JWTKey:       os.Getenv("JWT_KEY"),
		WriteTimeout: writeTimeout,
		ReadTimeout:  readTimeout,
	}, nil
}
