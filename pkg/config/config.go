package config

import (
	"fmt"
	"github.com/joho/godotenv"
	"os"
)

type Config struct {
	SecretKey  string
	DBDriver   string
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

var APP_CONFIG *Config

func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil {
		return nil, fmt.Errorf("failed to load .env file: %v", err)
	}

	config := Config{
		SecretKey:  os.Getenv("SECRET_KEY"),
		DBDriver:   os.Getenv("DB_DRIVER"),
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     os.Getenv("DB_PORT"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
	}
	APP_CONFIG = &config

	return &config, nil
}
