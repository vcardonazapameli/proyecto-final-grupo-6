package config

import (
	"fmt"
	"os"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	ServerAddress string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPassword    string
	DBName        string
}

var (
	config *Config
	once   sync.Once
)

func LoadConfig() (*Config, error) {
	var err error
	once.Do(func() {
		if err = godotenv.Load(); err != nil {
			err = fmt.Errorf("error loading .env file: %w", err)
			return
		}
		config = &Config{
			ServerAddress: getEnv("SERVER_ADDRESS", ""),
			DBHost:        getEnv("DB_HOST", ""),
			DBPort:        getEnv("DB_PORT", ""),
			DBUser:        getEnv("DB_USER", ""),
			DBPassword:    getEnv("DB_PASSWORD", ""),
			DBName:        getEnv("DB_NAME", ""),
		}
	})
	return config, err
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
