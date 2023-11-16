package adapter

import (
	"authTest/pkg/main_app/user/domain"
	db "authTest/pkg/main_app/user/repository"
	"authTest/pkg/storage/postgres"
	"context"
)

func CreateUser(ctx context.Context, user *domain.User, otp string, hashedPassword string) (db.User, error) {
	queries := db.New(postgres.DB)

	params := db.CreateUserParams{
		Firstname:  user.Firstname,
		Middlename: user.Middlename,
		Lastname:   user.Lastname,
		Email:      user.Email,
		Password:   hashedPassword,
		Otp:        otp,
	}

	return queries.CreateUser(ctx, params)

}

func GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	queries := db.New(postgres.DB)
	return queries.GetUserByEmail(ctx, email)

}

func UpdateUserByEmail(ctx context.Context, email string) error {
	queries := db.New(postgres.DB)
	return queries.UpdateUserByEmail(ctx, email)
}
