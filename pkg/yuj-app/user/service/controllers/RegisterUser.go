package controllers

import (
	"context"
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"net/mail"
	"net/smtp"
	"os"
	"strings"
)

func RegisterUserController(w http.ResponseWriter, r *http.Request, queries *db.Queries) {
	var user models.User
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&user); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid JSON")
		return
	}

	if user.Email == "" {
		respondWithError(w, http.StatusBadRequest, "Email is missing")
		return
	}

	_, err := mail.ParseAddress(user.Email)
	if err != nil {
		respondWithError(w, http.StatusBadRequest, "Please provide a valid email Adress")
		return
	}

	if user.Name == "" {
		respondWithError(w, http.StatusBadRequest, "Name is missing")
		return
	}
	if user.Password == "" {
		respondWithError(w, http.StatusBadRequest, "Password is missing")
		return
	}

	if err := user.HashPassword(user.Password); err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	otp, err := generateOTP()

	user.OTP = otp

	if err != nil {
		log.Fatal("error in generating OTP", err)
	}
	var userError error
	_, userError = queries.CreateUser(context.Background(), db.CreateUserParams{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
		Otp:      sql.NullString{String: user.OTP, Valid: true},
	})

	if userError != nil {
		// Log the error for debugging
		log.Println("Error creating user:", userError)

		if strings.Contains(userError.Error(), "unique constraint") {
			respondWithError(w, http.StatusConflict, "Email already in use")
		} else {
			respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		}
		return
	}

	//!sending otp
	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), "smtp.gmail.com")

	to := []string{user.Email}

	message := []byte("To: " + user.Email + "\r\n" +
		"Subject: OTP for Registration\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" +
		"<html><body>" +
		"<h1>Your OTP for registration is <strong>" + otp + "</strong></h1>" +
		"</body></html>")

	go func() {
		err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("EMAIL"), to, message)
		if err != nil {
			log.Println("Error in sending OTP:", err)
		}
	}()
	//!

	// log.Println(record)

	respondWithJSON(w, http.StatusCreated, map[string]string{"message": "User Registered Successfully, OTP sent to your email"})

}
