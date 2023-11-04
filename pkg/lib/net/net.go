package net

import (
	"encoding/json"
	"log"
	"net/http"
	"net/smtp"
	"os"
)

func SendOtpByEmail(email, otp string) {
	auth := smtp.PlainAuth("", os.Getenv("EMAIL"), os.Getenv("PASSWORD"), "smtp.gmail.com")

	to := []string{email}

	message := []byte("To: " + email + "\r\n" +
		"Subject: OTP for Verification\r\n" +
		"MIME-Version: 1.0\r\n" +
		"Content-Type: text/html; charset=\"utf-8\"\r\n\r\n" +
		"<html><body>" +
		"<h1>Your OTP for Verification is <strong>" + otp + "</strong></h1>" +
		"</body></html>")

	go func() {
		err := smtp.SendMail("smtp.gmail.com:587", auth, os.Getenv("EMAIL"), to, message)
		if err != nil {
			log.Println("Error in sending OTP:", err)
		}
	}()

}

func RespondWithJSON(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Println("Error encoding JSON response:", err)
	}
}

func RespondWithError(w http.ResponseWriter, status int, message string) {
	response := map[string]string{"error": message}
	RespondWithJSON(w, status, response)
}
