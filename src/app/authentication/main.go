package main

import (
	"Stock_broker_application/src/app/authentication/router"
	"Stock_broker_application/src/app/authentication/utils"
	"log"
)

func main() {

	utils.CreateConnection()
	// Set up router
	r := router.SetUpRouter()

	// Start the HTTP server
	err := r.Run(":8080")
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

}
