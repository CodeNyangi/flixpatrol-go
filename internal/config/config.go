package config

import (
	"os"
	"time"

	"github.com/joho/godotenv"
)

type Config struct {
	APIKey      string
	BaseURL     string
	HTTPTimeout time.Duration
}

func Load() (*Config, error) {
	// .env 파일 로드 (옵션)
	if err := godotenv.Load(); err != nil {
		return nil, err
	}

	config := &Config{
		APIKey:      getEnv("FLIXPATROL_API_KEY", ""),
		BaseURL:     getEnv("FLIXPATROL_BASE_URL", "https://flixpatrol.com/api/v1.4/"),
		HTTPTimeout: getDuration("FLIXPATROL_HTTP_TIMEOUT", 10*time.Second),
	}

	return config, nil
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}

func getDuration(key string, defaultValue time.Duration) time.Duration {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	duration, err := time.ParseDuration(value)
	if err != nil {
		return defaultValue
	}
	return duration
}
