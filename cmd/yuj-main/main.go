package main

import (
	db_scripts "yuj/pkg/storage/db-scripts"
)

func main() {
	//* Connecting to database
	db_scripts.ConnectDB()
}
