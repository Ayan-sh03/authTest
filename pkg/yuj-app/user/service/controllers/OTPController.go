package controllers

import (
	"context"
	"encoding/json"
	"log"

	"net/http"
)

func CheckOtpController(w http.ResponseWriter, r *http.Request, q *db.Queries) {
	var otpRequest OTPRequest
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&otpRequest); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request")
		return
	}

	user, err := q.GetUserByEmail(context.Background(), otpRequest.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}
	userotp := user.Otp
	log.Println(userotp.String)
	if otpRequest.OTP != (userotp.String) {
		respondWithError(w, http.StatusUnauthorized, "Invalid OTP")
		return
	}

	tokenString, err := authorization.GenerateJWT(user.Email)
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
		return
	}

	go func() {
		if err := q.UpdateUserByEmail(context.Background(), otpRequest.Email); err != nil {
			respondWithError(w, http.StatusInternalServerError, "Internal Server Error")
			return
		}
	}()

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "OTP Verified", "token": tokenString})
}
