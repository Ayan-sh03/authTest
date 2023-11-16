package service

import (
	network "authTest/pkg/lib/net"
	"authTest/pkg/lib/security"
	"authTest/pkg/lib/util"
	"authTest/pkg/lib/validation"
	"authTest/pkg/main_app/user/domain"
	db "authTest/pkg/main_app/user/repository"
	"authTest/pkg/storage/postgres"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strings"
)

// ^ Register :
//
//	@Summary		Register route
//	@Description	Allows users to create a new account.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			user	body		doc_model.Register		true	"User's firstname, lastname, middlename, email, password"
//	@Success		201		{object}	domain.User				"Successful response : User Model"
//	@Failure		400		{object}	doc_model.ErrorResponse	"Invalid JSON data, Invalid Email"
//	@Failure		409		{object}	doc_model.ErrorResponse	"User already exists"
//	@Failure		422		{object}	doc_model.ErrorResponse	"Please provide with sufficient credentials"
//	@Failure		500		{object}	doc_model.ErrorResponse	"Internal Server Error, Error in inserting the document, Error in hashing password, Error While generating OTP"
//	@Router			/user/register [post]
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
		network.RespondWithError(w, http.StatusInternalServerError, "Error While generating OTP "+err.Error())
		return
	}

	_, dbErr := queries.CreateUser(context.Background(), db.CreateUserParams{ //!
		Firstname:  user.Firstname,
		Middlename: user.Middlename,
		Lastname:   user.Lastname,
		Email:      user.Email,
		Password:   hashedPassword,
		Otp:        otp,
	})

	if dbErr != nil {
		log.Println("Error occured creating user", dbErr)
		if strings.Contains(dbErr.Error(), "\"users_email_key\"") {
			network.RespondWithError(w, http.StatusConflict, "User already exists")
			return
		}
		network.RespondWithError(w, http.StatusInternalServerError, dbErr.Error())
		return
	}

	go network.SendOtpByEmail(user.Email, otp)

	network.RespondWithJSON(w, http.StatusCreated, "User created successfully , OTP sent to Email Please Verify Your Account  ")

}

// ^ Login :
//
//	@Summary		Login route
//	@Description	Allows users to login into their account.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			Body	body		doc_model.Login				true	"User's email and password"
//	@Success		201		{object}	doc_model.SuccessResponse	"Successful response : Bearer \<token\>"
//	@Failure		400		{object}	doc_model.ErrorResponse		"Invalid JSON data, Invalid Email"
//	@Failure		401		{object}	doc_model.ErrorResponse		"Please Verify Your Account, Invalid Credentials"
//	@Failure		404		{object}	doc_model.ErrorResponse		"User is not registered"
//	@Failure		422		{object}	doc_model.ErrorResponse		"Please Verify Your Account"
//	@Failure		500		{object}	doc_model.ErrorResponse		"Internal server error"
//	@Router			/user/login [post]
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
		network.RespondWithError(w, http.StatusUnprocessableEntity, "Please Verify Your Account")
		go network.SendOtpByEmail(user.Email, dbUser.Otp)
		return

	}

	// check password
	securityErr := security.CheckPassword(user.Password, dbUser.Password)
	if securityErr != nil {
		network.RespondWithError(w, http.StatusUnauthorized, "Invalid Credentials : "+securityErr.Error())
		return
	}

	token, err := security.GenerateJWT(dbUser.Email, dbUser.ID) //! Changed

	if !err {
		network.RespondWithError(w, http.StatusInternalServerError, "Internal server error : Error While generating Token")
		return
	}

	network.RespondWithJSON(w, http.StatusOK, map[string]string{"token": token})
}

// ^ Validate Token :
//
//	@Summary		Validation route
//	@Description	Allows users to validate OTP and complete the registration process.
//	@Tags			user
//	@Accept			json
//	@Produce		json
//	@Param			Body	body		doc_model.OTP				true	"User's email address and otp"
//	@Success		200		{object}	doc_model.SuccessResponse	"Successful response : Bearer \<token\>"
//	@Failure		400		{object}	doc_model.ErrorResponse		"Invalid JSON data, Invalid Email"
//	@Failure		404		{object}	doc_model.ErrorResponse		"User Not Found"
//	@Failure		401		{object}	doc_model.ErrorResponse		"Invalid OTP, User Already Verified"
//	@Failure		500		{object}	doc_model.ErrorResponse		"Internal Server Error"
//	@Router			/user/otp [post]
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
	token, err := security.GenerateJWT(OtpRequest.Email, dbUser.ID) //! Changed
	if !err {
		network.RespondWithError(w, http.StatusInternalServerError, "Internal Servor Error : Error While generating Token")
		return
	}
	//!

	network.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "OTP Verified", "token": token})
}
