package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func CreateConnection() {

	// Load configuration values from the YAML file
	cfg := LoadConfig("resources/application.yml")

	// Opening a connection from Go application to the Database
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%d)/%s",
		cfg.Database.Username, cfg.Database.Password, cfg.Database.Host, cfg.Database.Port, cfg.Database.Name))
	if err != nil {
		log.Println("Error in opening the sql connection")
		panic(err.Error())
	}

	Db = db

	//defer Db.Close()
	//Verifying the Connection
	error1 := Db.Ping()
	if error1 != nil {
		fmt.Println("There is some error in verifying the connection with the database")
		return // Exit function if there's an error
	}

	fmt.Println("Successfully created a connection!")
}