package utils

import (
	"Stock_broker_application/src/app/authentication/models"
	"log"

	"github.com/spf13/viper"
)

//From here we will create the mapping of the yml file which is readable to go lang

// Config struct to hold the configuration values
type Config struct {
	Database models.DatabaseConfig
}

// LoadConfig loads configuration values from the provided YAML file
func LoadConfig(filePath string) Config {
	// Set the path to the config file: BAsically telling viper where to read data from.
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
