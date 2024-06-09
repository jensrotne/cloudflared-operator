package config

import (
	"os"

	_ "github.com/joho/godotenv/autoload"
)

func Get(key string) string {
	envVar := os.Getenv(key)

	return envVar
}
