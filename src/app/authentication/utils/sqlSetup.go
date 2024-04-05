package utils

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Client *sql.DB

func NewConnection() error {

	// Open up our database connection.
	// I've set up a database on my local machine using Docker.

	db, err := sql.Open("mysql", "root:Root@123@tcp(localhost:3306)/hospital")

	// if there is an error opening the connection, handle it
	if err != nil {
		// panic(err.Error())
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}

	Client = db
	return nil
}
