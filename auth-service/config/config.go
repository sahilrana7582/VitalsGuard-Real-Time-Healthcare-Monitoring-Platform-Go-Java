package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	DB_URL string
	PORT   string
}

func LoadAuthConfig(path string) *Config {
	if err := godotenv.Load(path); err != nil {
		log.Fatalf("❌ Failed to load env file from path %s: %v", path, err)
	}

	dbURL := os.Getenv("DB_URL")
	port := os.Getenv("PORT")

	if dbURL == "" {
		panic("❌ Missing required environment variable: DB_URL")
	}
	if port == "" {
		panic("❌ Missing required environment variable: PORT")
	}

	log.Println("✅ Tenants configuration loaded successfully.")

	return &Config{
		DB_URL: dbURL,
		PORT:   port,
	}
}
