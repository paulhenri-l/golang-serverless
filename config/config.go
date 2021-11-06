package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"github.com/paulhenri-l/goenv"
	"strings"
)

var cfg Config

type Config struct {
	AppEnv  string
	AppHost string
	AppPort string

	LogLevel string
}

func init() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("error loading .env file: %s\n", err)
	}

	cfg = Config{
		AppEnv:  strings.ToLower(goenv.StringOr("APP_ENV", "production")),
		AppHost: goenv.StringOr("APP_HOST", "127.0.0.1"),
		AppPort: goenv.StringOr("APP_PORT", "3000"),

		LogLevel: goenv.StringOr("LOG_LEVEL", "info"),
	}
}

func Get() Config {
	return cfg
}

func EditConfig(f func(*Config)) {
	f(&cfg)
}
