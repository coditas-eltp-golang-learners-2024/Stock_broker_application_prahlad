package db

import (
	"Stock_broker_application/src/app/authentication/utils"
	"database/sql"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *sql.DB
var GormDb *gorm.DB

// CreateConnection establishes a connection to the database using GORM.
func CreateConnection() {
	// Load configuration values from the YAML file
	cfg := utils.LoadConfig("resources/application.yml")

	// Construct the DSN (Data Source Name) for the database connection
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?parseTime=true", // Add `?parseTime=true` for time parsing
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name)

	// Open a new GORM connection to the database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	// Assign the GORM DB instance to the global variable
	GormDb = db

	fmt.Println("Successfully created a connection to the database using GORM!")
}
