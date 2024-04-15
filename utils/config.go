package utils

import (
	"Stock_broker_application/src/app/authentication/models"
	"log"

	"github.com/spf13/viper"
)

// Config struct to hold the configuration values
type Config struct {
	Database models.DatabaseConfig
}

// LoadConfig loads configuration values from the provided YAML file using Viper.
// It reads the YAML file specified by 'filePath', unmarshals it into a Config struct,
// and returns the Config struct containing the loaded configuration values.
// @param filePath path string true "Path to the YAML configuration file"
// @return Config
func LoadConfig(filePath string) Config {
	// Set the path to the config file
	viper.SetConfigFile(filePath)

	// Read in the config file
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error reading config file: %s", err)
	}

	// Unmarshal the config values into a Config struct
	var config Config
	err = viper.Unmarshal(&config)
	if err != nil {
		log.Fatalf("Error unmarshaling config file: %s", err)
	}

	return config
}
