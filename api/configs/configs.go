package configs

import (
	"log"
	"os"
	"strconv"
)

// Config struct holds all required environment variables
type Config struct {
	DBAddress string
	APIPort   int
	APIHost   string
	RateLimit int
	APIDomain string
}

// LoadConfig validates and loads environment variables
func LoadConfig() *Config {
	requiredVars := []string{"DB_ADDRESS", "API_PORT", "API_HOST", "API_RATE_LIMIT_THRESHOLD", "API_DOMAIN"}

	// Check if required variables are set
	for _, v := range requiredVars {
		if os.Getenv(v) == "" {
			log.Fatalf("Error: Required environment variable %s is not set", v)
		}
	}

	// Convert numeric values
	apiPort, err := strconv.Atoi(os.Getenv("API_PORT"))
	if err != nil {
		log.Fatalf("Error: API_PORT must be a valid integer")
	}

	rateLimit, err := strconv.Atoi(os.Getenv("API_RATE_LIMIT_THRESHOLD"))
	if err != nil {
		log.Fatalf("Error: API_RATE_LIMIT_THRESHOLD must be a valid integer")
	}

	return &Config{
		DBAddress: os.Getenv("DB_ADDRESS"),
		APIPort:   apiPort,
		APIHost:   os.Getenv("API_HOST"),
		RateLimit: rateLimit,
		APIDomain: os.Getenv("API_DOMAIN"),
	}
}
