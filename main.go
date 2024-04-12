package main

import (
	"Stock_broker_application/src/app/authentication/constants"
	_ "Stock_broker_application/src/app/authentication/docs"
	"Stock_broker_application/src/app/authentication/router"
	"Stock_broker_application/src/app/authentication/utils/db"
	"log"
)

//@title Stock Broker Application
//@version 1.0
//@description This is a Stock Broker Application API
//@host localhost:8080
//@BasePath /

func main() {

	// Create a database connection
	db.CreateConnection()

	// Defer closing the database connection when the function exits
	defer func() {
		sqlDB, err := db.GormDb.DB()
		if err != nil {
			panic("Failed to get underlying DB connection")
		}
		if err := sqlDB.Close(); err != nil {
			panic("Failed to close database connection")
		}
	}()

	// Setting up the router
	r := router.SetUpRouter()

	// Start the HTTP server
	serverErr := r.Run(":8080")
	if serverErr != nil {
		log.Fatalf("Failed to start server: %v", constants.ErrStartingServer)
	}

}
