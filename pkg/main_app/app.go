package main_app

import (
	"authTest/pkg/storage"
	"log"

	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//* Run app
	storage.ConnectDB()
}
