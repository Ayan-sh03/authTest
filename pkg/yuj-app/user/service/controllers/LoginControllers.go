package controllers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

func LoginController(w http.ResponseWriter, r *http.Request, q *db.Queries) {
	var request TokenRequest
	var user db.User
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&request); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	var err error
	user, err = q.GetUserByEmail(context.Background(), request.Email) // Use the generated function
	if err != nil {
		respondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
	// log.Println(user)

	if !user.IsVerified {
		auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), "smtp.gmail.com")
		otp := user.Otp

		to := []string{user.Email}

		message := []byte("To: " + user.Email + "\r\n" +
			"Subject: OTP for Registration\r\n" +
			"MIME-Version: 1.0\r\n" +
			"Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" +
			"<html><body>" +
			"<h1>Your OTP for registration is <strong>" + otp.String + "</strong></h1>" +
			"</body></html>")

		go func() {
			err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("EMAIL"), to, message)
			if err != nil {
				log.Println("Error in sending OTP:", err)
			}
		}()
		//!

		respondWithError(w, http.StatusUnauthorized, "User is not verified ,  please verify Otp Sent ")
		return
	}

	credentialsError := models.CheckPassword(request.Password, user.Password)
	if credentialsError != nil {
		respondWithError(w, http.StatusUnauthorized, "Invalid Credentials")
		return
	}

	tokenString, err := authorization.GenerateJWT(user.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"token": tokenString})
}
