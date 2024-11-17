package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

var (
	APP_PORT   string
	APP_ENV    string
	APP_LOCALE string
	APP_SECRET string
	APP_URL    string

	APP_DB_HOST    string
	APP_DB_PORT    string
	APP_DB_NAME    string
	APP_DB_USER    string
	APP_DB_PASS    string
	APP_DB_SCHEMA  string
	APP_DB_SSLMODE string
)

func setup() {
	APP_PORT = GetEnv("APP_PORT", "4000")
	APP_ENV = GetEnv("APP_ENV", "development")
	APP_LOCALE = GetEnv("APP_LOCALE", "en")
	APP_SECRET = GetEnv("APP_SECRET", "secret")
	APP_URL = GetEnv("APP_URL", "localhost:4000")

	APP_DB_HOST = GetEnv("APP_DB_HOST", "localhost")
	APP_DB_PORT = GetEnv("APP_DB_PORT", "5432")
	APP_DB_NAME = GetEnv("APP_DB_NAME", "go-fiber-react")
	APP_DB_USER = GetEnv("APP_DB_USER", "postgres")
	APP_DB_PASS = GetEnv("APP_DB_PASS", "\"\"")
	APP_DB_SCHEMA = GetEnv("APP_DB_SCHEMA", "public")
	APP_DB_SSLMODE = GetEnv("APP_DB_SSLMODE", "disable")
}

func Env() {
	if err := godotenv.Load(".env"); err != nil {
		fmt.Printf("Error loading .env file: %v \n\n", err.Error())
		os.Exit(1)
	}
	setup()
}

func GetEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok && value != "" {
		return value
	}
	return fallback
}
