package internal

import (
	"log"
	"os"
)

// Config struct to hold configuration values
// Add your config fields here
type Config struct {
	// Example field
	Port string `json:"port"`
}

// Load function to load environment variables into Config struct
func Load() (*Config, error) {
	config := &Config{}
	
	// Replace with environment variable keys you want to load
	config.Port = os.Getenv("APP_PORT")

	if config.Port == "" {
		log.Println("Warning: APP_PORT is not set, using default value")
	}
	
	return config, nil
}