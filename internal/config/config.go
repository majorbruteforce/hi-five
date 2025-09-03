package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	log "github.com/majorbruteforce/hifive/pkg/logger"
)

type Config struct {
	Env string
}

func Load() Config {
	_ = godotenv.Load()

	return Config{
		Env: getEnv("ENV", "development"),
	}
}

func getEnv(key, fallback string) string {
	if val := os.Getenv(key); val != "" {
		return val
	}
	return fallback
}

func getEnvAsInt(key string, fallback int) int {
	if val := os.Getenv(key); val != "" {
		var i int
		_, err := fmt.Sscanf(val, "%d", &i)
		if err == nil {
			return i
		}
		log.Log.Errorf("Invalid int for %s: %s", key, val)
	}
	return fallback
}

func getEnvAsFloat(key string, fallback float64) float64 {
	if val := os.Getenv(key); val != "" {
		var f float64
		_, err := fmt.Sscanf(val, "%f", &f)
		if err == nil {
			return f
		}
		log.Log.Errorf("Invalid float for %s: %s", key, val)
	}
	return fallback
}
