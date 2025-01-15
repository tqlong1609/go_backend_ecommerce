package implements

import (
	"context"

	"github.com/tqlong1609/go_backend_ecommerce/internal/database"
	"github.com/tqlong1609/go_backend_ecommerce/internal/model"
)

type UserLoginImplement struct {
	dbQueries *database.Queries
}

func InitUserLoginImplement(dbQueries *database.Queries) *UserLoginImplement {
	return &UserLoginImplement{
		dbQueries: dbQueries,
	}
}

func (ul *UserLoginImplement) Login(ctx context.Context) error {
	return nil
}

func (ul *UserLoginImplement) Register(ctx context.Context, params model.RegisterInput) (response model.RegisterOutput, err error) {
	salt := "test"
	userId, err := ul.dbQueries.CreateNewUser(ctx, database.CreateNewUserParams{
		UserAccount:  params.Account,
		UserPassword: params.Password,
		UserSalt:     salt,
	})
	if err != nil {
		return response, err
	}
	return model.RegisterOutput{
		UserID: userId,
	}, nil
}

func (ul *UserLoginImplement) Logout(ctx context.Context) error {
	return nil
}

func (ul *UserLoginImplement) VerifyOTP(ctx context.Context) error {
	return nil
}

func (ul *UserLoginImplement) UpdatePassword(ctx context.Context) error {
	return nil
}
