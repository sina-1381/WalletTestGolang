package config

import (
	"WalletTest/pkg/database"
	"fmt"
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Env struct {
	APP_PORT                   string
	APP_TIMEZONE               string
	APP_LOCATION               *time.Location
	LOG_FOLDER                 string
	POSTGRESQL_USER            string
	POSTGRESQL_PASSWORD        string
	POSTGRESQL_HOST            string
	POSTGRESQL_PORT            string
	POSTGRESQL_DATABASE        string
	POSTGRESQL_SSLMODE         string
	POSTGRESQL_MIGRATIONS_PATH string
	ADMINER_PORT               string
	REDIS_ADDRESS              string
	REDIS_PORT                 string
}

// env variables with default values if no env file is found

func NewEnv() *Env {
	_ = godotenv.Load()
	return &Env{
		APP_PORT:                   getEnv("APP_PORT", "50051"),
		APP_TIMEZONE:               getEnv("APP_TIMEZONE", "Asia/Tehran"),
		APP_LOCATION:               setLocation(),
		LOG_FOLDER:                 getEnv("LOG_FOLDER", "logs/"),
		POSTGRESQL_USER:            getEnv("POSTGRESQL_USER", "root"),
		POSTGRESQL_PASSWORD:        getEnv("POSTGRESQL_PASSWORD", "root"),
		POSTGRESQL_HOST:            getEnv("POSTGRESQL_HOST", "localhost"),
		POSTGRESQL_PORT:            getEnv("POSTGRESQL_PORT", "5432"),
		POSTGRESQL_DATABASE:        getEnv("POSTGRESQL_DATABASE", "WalletTest"),
		POSTGRESQL_SSLMODE:         getEnv("POSTGRESQL_SSLMODE", "disable"),
		POSTGRESQL_MIGRATIONS_PATH: getEnv("POSTGRESQL_MIGRATIONS_PATH", "file://migrations"),
		ADMINER_PORT:               getEnv("ADMINER_PORT", "8080"),
		REDIS_ADDRESS:              getEnv("REDIS_ADDRESS", "localhost"),
		REDIS_PORT:                 getEnv("REDIS_PORT", "6379"),
	}
}

func getEnv(key, defaultVal string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return defaultVal
}

func setLocation() *time.Location {
	location, err := time.LoadLocation(getEnv("APP_TIMEZONE", "Asia/Tehran"))
	if err != nil {
		location, _ = time.LoadLocation("UTC")
	}
	return location
}

func (e *Env) PostgresqlConfig() database.PostgresqlConfig {
	return database.PostgresqlConfig{
		User:           e.POSTGRESQL_USER,
		Password:       e.POSTGRESQL_PASSWORD,
		Host:           e.POSTGRESQL_HOST,
		Port:           e.POSTGRESQL_PORT,
		Database:       e.POSTGRESQL_DATABASE,
		SSLMode:        e.POSTGRESQL_SSLMODE,
		MigrationsPath: e.POSTGRESQL_MIGRATIONS_PATH,
	}
}

func (e *Env) RedisConfig() string {
	return fmt.Sprintf("%s:%s", e.REDIS_ADDRESS, e.REDIS_PORT)
}
