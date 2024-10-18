package server

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

type Config struct {
	Host       string
	Port       string
	FilePath   string
	DataSource string
}

// getEnv obtains an environment variable or returns a default value provided.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

// LoadConfig initializes the server settings using the environment values provided.
func LoadConfig() Config {
	if err := godotenv.Load(); err != nil {
		log.Printf("Warning: Could not load .env file, using system environment variables instead.")
	}

	return Config{
		Host:       getEnv("HOST", "localhost"),
		Port:       getEnv("PORT", ":8080"),
		FilePath:   getEnv("FILEPATH", "internal/repository/adapters/memory/products.json"),
		DataSource: getEnv("DATA_SOURCE", "json"),
	}
}
