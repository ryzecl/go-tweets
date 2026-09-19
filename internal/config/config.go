package config

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port           string
	DBUrlMigration string
	SecretJWT      string

	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig() (*Config, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, fmt.Errorf("Failed to load .env file : %v", err)
	}

	log.Println("Config loaded successfully")

	return &Config{
		Port:           os.Getenv("PORT"),
		DBUrlMigration: os.Getenv("DATABASE_URL"),
		SecretJWT:      os.Getenv("SECRET_JWT"),
		DBHost:         os.Getenv("DB_HOST"),
		DBPort:         os.Getenv("DB_PORT"),
		DBUser:         os.Getenv("DB_USER"),
		DBPassword:     os.Getenv("DB_PASSWORD"),
		DBName:         os.Getenv("DB_NAME"),
	}, nil
}
