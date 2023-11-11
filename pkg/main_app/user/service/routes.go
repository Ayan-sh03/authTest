package service

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func SetupRoutes() *httprouter.Router {
	router := httprouter.New()
	basePath := "/api/v1/users"

	router.POST(basePath+"/register", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		RegisterUserController(w, r)
	})

	router.POST(basePath+"/login", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		LoginController(w, r)
	})

	router.POST(basePath+"/otp", func(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
		VerifyOtpController(w, r)
	})

	return router
}
