package config

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DATABASE_URL string
	PORT         string
}

func Load() *Config {
	if err := godotenv.Load(".env"); err != nil {
		slog.Warn("No env file found, using system env")
	}
	return &Config{
		DATABASE_URL: GetEnv("DATABASE_URL", "postgresql://postgres:abraham@localhost:5432/back-to-go"),
		PORT:         GetEnv("PORT", ":4000"),
	}

}

func GetEnv(key string, fallback string) string {
	val, exits := os.LookupEnv(key)

	if !exits {
		slog.WarnContext(context.Background(), fmt.Sprintf("Value of %v in env not found", key))
		return fallback
	}

	return val
}
