package postgres

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

const connectMsg string = "---------------------------------------------------------------------------------------------\nConnected to DB\n---------------------------------------------------------------------------------------------"

func Postgres() *sql.DB {

	uri := os.Getenv("SQLURI")

	// Open a connection to the database
	db, err := sql.Open("postgres", uri)
	if err != nil {
		log.Println(err)
		return nil
	}

	// Ping the database to check if the connection is valid
	err = db.Ping()
	if err != nil {
		log.Println(err)
		return nil
	}

	DB = db

	fmt.Println(connectMsg)
	return db
}
