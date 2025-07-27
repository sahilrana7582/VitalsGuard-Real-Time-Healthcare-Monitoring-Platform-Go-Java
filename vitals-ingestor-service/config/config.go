package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	PORT   string
	DB_URL string
	kafka  KafkaConfig
}

type KafkaConfig struct {
	Brokers []string
	Topic   string
}

func NewConfig() *Config {
	godotenv.Load(".env")

	port := os.Getenv("PORT")
	db_url := os.Getenv("DB_URL")

	kafkaBrokersStr := "localhost:9092"
	kafkaBrokers := []string{kafkaBrokersStr}
	kafkaTopic := "vitals"

	return &Config{
		kafka: KafkaConfig{
			Brokers: kafkaBrokers,
			Topic:   kafkaTopic,
		},
		PORT:   port,
		DB_URL: db_url,
	}
}
