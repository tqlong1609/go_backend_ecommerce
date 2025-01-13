package repo

import (
	"context"

	"github.com/tqlong1609/go_backend_ecommerce/global"
	"github.com/tqlong1609/go_backend_ecommerce/internal/database"
)

type IUserRepository interface {
	CreateNewUser(account string, password string, salt string) (int32, error)
	GetUserAccountByUserId(userId int32) (string, error)
}

type userRepository struct {
	sqlc *database.Queries
}

func InitUserRepository() IUserRepository {
	return &userRepository{
		sqlc: database.New(global.Mdbc),
	}
}

func (ur *userRepository) CreateNewUser(account string, password string, salt string) (int32, error) {
	userId, err := ur.sqlc.CreateNewUser(context.Background(), database.CreateNewUserParams{
		UserAccount:  account,
		UserPassword: password,
		UserSalt:     salt,
	})
	return userId, err
}

func (ur *userRepository) GetUserAccountByUserId(userId int32) (string, error) {
	userAccount, err := ur.sqlc.GetUserAccountByUserId(context.Background(), userId)
	return userAccount, err
}
