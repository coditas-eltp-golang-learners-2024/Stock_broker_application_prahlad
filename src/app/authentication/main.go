package main

import (
	"Stock_broker_application/src/app/authentication/constants"
	"Stock_broker_application/src/app/authentication/router"
	utils "Stock_broker_application/src/app/authentication/utils/db"
	"log"
)

func main() {

	utils.CreateConnection()
	defer utils.Db.Close() //Closing the database connection after main function is executed

	// Setting up router
	r := router.SetUpRouter()

	// Start the HTTP server
	err := r.Run(":8081")
	if err != nil {
		log.Fatalf("Found Error: %v", constants.ErrStartingServer)
	}

}
