package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	SecretKey string
	DBUrl     string
}

var AppConfig *Config

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	config := Config{
		SecretKey: os.Getenv("SECRET_KEY"),
		DBUrl:     os.Getenv("DB_URL"),
	}
	AppConfig = &config

	return &config, nil
}
