package config

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"time"

	"github.com/joho/godotenv"
)

type Config interface {
	APIKey() string
	BaseURL() string
	HTTPTimeout() time.Duration
}

type configImpl struct {
	apiKey      string
	baseURL     string
	httpTimeout time.Duration
}

func (c *configImpl) APIKey() string             { return c.apiKey }
func (c *configImpl) BaseURL() string            { return c.baseURL }
func (c *configImpl) HTTPTimeout() time.Duration { return c.httpTimeout }

func findProjectRoot() (string, error) {
	_, b, _, _ := runtime.Caller(0)
	currentDir := filepath.Dir(b)
	for {
		if _, err := os.Stat(filepath.Join(currentDir, "go.mod")); err == nil {
			return currentDir, nil
		}
		if currentDir == filepath.Dir(currentDir) {
			return "", fmt.Errorf("could not find project root")
		}
		currentDir = filepath.Dir(currentDir)
	}
}

func Load() (Config, error) {
	projectRoot, err := findProjectRoot()
	if err != nil {
		return nil, fmt.Errorf("finding project root: %w", err)
	}

	envPath := filepath.Join(projectRoot, ".env")
	log.Printf("Trying to load .env from: %s", envPath)

	err = godotenv.Load(envPath)
	if err != nil {
		log.Printf("Failed to load .env file: %v", err)
		return nil, fmt.Errorf("loading .env file: %w", err)
	}
	log.Println("Successfully loaded .env file")

	apiKey := getEnv("FLIXPATROL_API_KEY", "")
	if apiKey == "" {
		log.Println("Warning: FLIXPATROL_API_KEY is not set")
	} else {
		log.Println("FLIXPATROL_API_KEY is set")
	}

	return &configImpl{
		apiKey:      apiKey,
		baseURL:     getEnv("FLIXPATROL_BASE_URL", "https://flixpatrol.com/api/v1.4/"),
		httpTimeout: getDuration("FLIXPATROL_HTTP_TIMEOUT", 10*time.Second),
	}, nil
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
