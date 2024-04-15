package models

// DatabaseConfig represents the configuration values for database connection.
type DatabaseConfig struct {
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
	Username string `yamls:"username"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}
