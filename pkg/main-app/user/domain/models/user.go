package domain

import (
	db "authTest/pkg/main-app/user/repository"
	"database/sql"
)

type User struct {
	Firstname  string  `json:"firstname"`
	Middlename *string `json:"middlename"`
	Lastname   string  `json:"lastname"`
	Email      string  `json:"email"`
	Password   string  `json:"password"`
}

func ToRepository(user *User) *db.CreateUserParams {
	return &db.CreateUserParams{
		Firstname:  user.Firstname,
		Middlename: sql.NullString{String: *user.Middlename, Valid: user.Middlename != nil},
		Lastname:   user.Lastname,
		Email:      user.Email,
		Password:   user.Password,
	}
}
