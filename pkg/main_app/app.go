package main_app

import (
	"authTest/pkg/main_app/user/service"
	"authTest/pkg/storage"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

func Run() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//* Run app
	storage.ConnectDB()

	//* Initialse router
	router := service.SetupRoutes()

	log.Fatal(http.ListenAndServe(":8080", router))
}
