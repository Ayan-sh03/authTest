// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.22.0

package db

import (
	"time"
)

type User struct {
	ID         int64
	Firstname  string
	Middlename string
	Lastname   string
	Email      string
	Password   string
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsVerified bool
	Otp        string
}
