package main

import (
	main_app "authTest/pkg/main-app"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	main_app.Run()
}
