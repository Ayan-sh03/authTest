package model

import (
	db "authTest/pkg/main_app/user/repository"
)

type User struct {
	Firstname  string `json:"firstname"`
	Middlename string `json:"middlename"`
	Lastname   string `json:"lastname"`
	Email      string `json:"email"`
	Password   string `json:"password"`
}

func ToRepository(user *User) *db.CreateUserParams {
	return &db.CreateUserParams{
		Firstname:  user.Firstname,
		Middlename: user.Middlename,
		Lastname:   user.Lastname,
		Email:      user.Email,
		Password:   user.Password,
	}
}
