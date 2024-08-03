package config

import (
	"os"
	"time"
)

type Config struct {
	Host       string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string
	ServerPort string
	Salt       string
	JWTTTL     time.Duration
	JWTKey     string
}

func Load() (Config, error) {
	jwtTTL, err := time.ParseDuration(os.Getenv("JWT_TTL"))
	if err != nil {
		return Config{}, err
	}

	return Config{
		Host:       os.Getenv("HOST"),
		DBPort:     os.Getenv("P_DB_PORT"),
		DBUser:     os.Getenv("P_DB_USER"),
		DBPassword: os.Getenv("P_DB_PASS"),
		DBName:     os.Getenv("P_DB_NAME"),
		DBSSLMode:  os.Getenv("P_DB_SSL_MODE"),
		ServerPort: os.Getenv("SERVER_PORT"),
		Salt:       os.Getenv("PW_SALT"),
		JWTTTL:     jwtTTL,
		JWTKey:     os.Getenv("JWT_KEY"),
	}, nil
}
