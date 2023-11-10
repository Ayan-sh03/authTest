package service

import (
	network "authTest/pkg/lib/net"
	"authTest/pkg/lib/security"
	"authTest/pkg/lib/util"

	domain "authTest/pkg/main_app/user/domain/model"
	db "authTest/pkg/main_app/user/repository"

	"context"
	"database/sql"
	"encoding/json"
	"net/http"
)

func RegisterUserController(w http.ResponseWriter, r *http.Request, q *db.Queries) {

	var user domain.User

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		network.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//!Validate here

	//!

	hashedPassword, err := security.HashPassword(user.Password)

	if err != nil {
		network.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	otp, err := util.GenerateOTP()
	if err != nil {
		network.RespondWithError(w, http.StatusInternalServerError, "Error While generating OTP"+err.Error())
		return
	}

	_, dberr := q.CreateUser(context.Background(), db.CreateUserParams{
		Firstname:  user.Firstname,
		Middlename: sql.NullString{String: *user.Middlename, Valid: user.Middlename != nil},
		Lastname:   user.Lastname,
		Email:      user.Email,
		Password:   hashedPassword,
		Otp:        otp,
	})

	if dberr != nil {
		network.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	go network.SendOtpByEmail(user.Email, otp)

	network.RespondWithJSON(w, http.StatusCreated, "User created successfully , OTP sent to Email Please Verify Your Account  ")

}

func LoginController(w http.ResponseWriter, r *http.Request, q *db.Queries) {
	var user domain.User

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		network.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//!Validate here

	//!

	dbUser, userErr := q.GetUserByEmail(context.Background(), user.Email)

	if userErr != nil {
		network.RespondWithError(w, http.StatusInternalServerError, userErr.Error())
		return
	}

	// check if verified

	if !dbUser.IsVerified {

		network.RespondWithError(w, http.StatusUnauthorized, "Please Verify Your Account")
		go network.SendOtpByEmail(user.Email, dbUser.Otp)
		return

	}

	// check password
	securityErr := security.CheckPassword(user.Password, dbUser.Password)
	if securityErr != nil {
		network.RespondWithError(w, http.StatusUnauthorized, securityErr.Error())
		return
	}

	token, err := security.GenerateJWT(dbUser.Email)

	if err != nil {
		network.RespondWithError(w, http.StatusInternalServerError, "Error While generating Token"+err.Error())
		return
	}

	network.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})

}

func VerifyOtpController(w http.ResponseWriter, r *http.Request, q *db.Queries) {
	var OtpRequest struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&OtpRequest); err != nil {
		network.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//!Validate here

	//!

	dbUser, userErr := q.GetUserByEmail(context.Background(), OtpRequest.Email)

	if userErr != nil {
		if userErr.Error() == "sql: no rows in result set" {
			network.RespondWithError(w, http.StatusNotFound, "User Not Found")
			return
		}
		network.RespondWithError(w, http.StatusInternalServerError, userErr.Error())
		return
	}

	if dbUser.IsVerified {
		network.RespondWithError(w, http.StatusUnauthorized, "User Already Verified")
		return
	}
	if OtpRequest.OTP != dbUser.Otp {
		network.RespondWithError(w, http.StatusUnauthorized, "Invalid OTP")
		return
	}

	go func() {
		err := q.UpdateUserByEmail(context.Background(), dbUser.Email)

		if err != nil {
			network.RespondWithError(w, http.StatusInternalServerError, err.Error())
			return
		}

	}()

	token, err := security.GenerateJWT(OtpRequest.Email)

	if err != nil {
		network.RespondWithError(w, http.StatusInternalServerError, "Error While generating Token"+err.Error())
		return
	}

	network.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "OTP Verified", "token": token})
}
