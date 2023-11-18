package adapter

import (
	"authTest/pkg/main_app/user/domain"
	db "authTest/pkg/main_app/user/repository"
	helper "authTest/pkg/main_app/user/repository/helpers"
	"authTest/pkg/storage/postgres"
	"context"
)

func CreateUser(ctx context.Context, user *domain.User, otp string, hashedPassword string) (domain.User, error) {
	queries := db.New(postgres.DB)

	params := db.CreateUserParams{
		Firstname:  user.Firstname,
		Middlename: user.Middlename,
		Lastname:   user.Lastname,
		Email:      user.Email,
		Password:   hashedPassword,
		Otp:        otp,
	}

	repoUser, err := queries.CreateUser(ctx, params)

	if err != nil {
		return domain.User{}, err
	}

	domainUser := helper.ToDomainUser(&repoUser)

	return *domainUser, nil

}

func GetUserByEmail(ctx context.Context, email string) (db.User, error) {
	queries := db.New(postgres.DB)
	return queries.GetUserByEmail(ctx, email)

}

func UpdateUserByEmail(ctx context.Context, email string) error {
	queries := db.New(postgres.DB)
	return queries.UpdateUserByEmail(ctx, email)
}
