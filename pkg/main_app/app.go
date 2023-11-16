package main_app

import (
	_ "authTest/pkg/main_app/docs"
	"authTest/pkg/main_app/user/service"
	"authTest/pkg/storage"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	httpSwagger "github.com/swaggo/http-swagger"
)

//	@title			Registration API
//	@version		1.0
//	@description	This is a registration api for an application.
//	@BasePath		/api/v1

func Run() {
	err := godotenv.Load("../../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//* Run app
	storage.ConnectDB()

	//* Initialse router
	router := service.SetupRoutes()
	router.PathPrefix("/docs/").Handler(httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
	log.Fatal(http.ListenAndServe(":8080", router))
}
