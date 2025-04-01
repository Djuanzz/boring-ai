package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	OpenAIKey string
	GMapsKey  string
}

func LoadEnv() Config {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}

	config := Config{
		Port:      getEnv("SERVER_PORT", "5000"),
		OpenAIKey: getEnv("OPENAI_KEY", ""),
		GMapsKey:  getEnv("GMAPS_KEY", ""),
	}

	return config
}

func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
