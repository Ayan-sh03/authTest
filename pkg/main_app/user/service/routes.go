package service

import (
	"net/http"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()
	basePath := "/api/v1/users"

	userRouter := router.PathPrefix(basePath).Subrouter()

	userRouter.HandleFunc("/register", func(w http.ResponseWriter, r *http.Request) {
		RegisterUserController(w, r)
	}).Methods("POST")

	userRouter.HandleFunc("/login", func(w http.ResponseWriter, r *http.Request) {
		LoginController(w, r)
	}).Methods("POST")

	userRouter.HandleFunc("/otp", func(w http.ResponseWriter, r *http.Request) {
		VerifyOtpController(w, r)
	}).Methods("POST")

	return router
}
