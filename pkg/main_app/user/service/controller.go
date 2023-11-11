package service

import (
	network "authTest/pkg/lib/net"
	"authTest/pkg/lib/security"
	"authTest/pkg/lib/util"
	"authTest/pkg/lib/validation"
	"authTest/pkg/storage/postgres"
	"log"

	domain "authTest/pkg/main_app/user/domain/model"
	db "authTest/pkg/main_app/user/repository"

	"context"

	"encoding/json"
	"net/http"
)

func RegisterUserController(w http.ResponseWriter, r *http.Request) {

	var user domain.User

	decoder := json.NewDecoder(r.Body)
	queries := db.New(postgres.DB)
	if err := decoder.Decode(&user); err != nil {
		network.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//!Validate here
	err := validation.UserValidator(&user)
	if err != nil {
		network.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}
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

	log.Printf("user: %+v", user)
	log.Printf("hashedPassword: %s", hashedPassword)
	log.Printf("otp: %s", otp)
	log.Printf("printing query %v", queries)

	_, dberr := queries.CreateUser(context.Background(), db.CreateUserParams{ //!
		Firstname:  user.Firstname,
		Middlename: user.Middlename,
		Lastname:   user.Lastname,
		Email:      user.Email,
		Password:   hashedPassword,
		Otp:        otp,
	})

	if dberr != nil {
		log.Fatal("Error occured creating user", dberr)
		network.RespondWithError(w, http.StatusInternalServerError, dberr.Error())
		return
	}

	go network.SendOtpByEmail(user.Email, otp)

	network.RespondWithJSON(w, http.StatusCreated, "User created successfully , OTP sent to Email Please Verify Your Account  ")

}

func LoginController(w http.ResponseWriter, r *http.Request) {
	var user domain.User

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		network.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//!Validate here
	validationErr := validation.UserValidator(&user)
	if validationErr != nil {
		network.RespondWithError(w, http.StatusBadRequest, validationErr.Error())
		return
	}
	//!
	q := db.New(postgres.DB)

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

	token, err := security.GenerateJWT(dbUser.Email, dbUser.ID)

	if err {
		network.RespondWithError(w, http.StatusInternalServerError, "Error While generating Token")
		return
	}

	network.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})

}

func VerifyOtpController(w http.ResponseWriter, r *http.Request) {
	var OtpRequest struct {
		Email string `json:"email"`
		OTP   string `json:"otp"`
	}
	q := db.New(postgres.DB)

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&OtpRequest); err != nil {
		network.RespondWithError(w, http.StatusBadRequest, err.Error())
		return
	}

	//!Validate here
	if !validation.IsValidEmail(OtpRequest.Email) {
		network.RespondWithError(w, http.StatusBadRequest, "Invalid Email")
		return
	}
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

	//!
	token, err := security.GenerateJWT(OtpRequest.Email, dbUser.ID)
	if !err {
		network.RespondWithError(w, http.StatusInternalServerError, "Internal Servor Error : Error While generating Token")
		return
	}
	//!

	network.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "OTP Verified", "token": token})
}
