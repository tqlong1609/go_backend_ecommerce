package services

import (
	"github.com/tqlong1609/go_backend_ecommerce/global"
	"github.com/tqlong1609/go_backend_ecommerce/internal/repo"
	"github.com/tqlong1609/go_backend_ecommerce/pkg/response"
	"go.uber.org/zap"
)

type IUserService interface {
	Register(account string, password string) int32
	GetUserAccountByUserId(userId int32) string
}

type userService struct {
	userRepo repo.IUserRepository
}

func (us *userService) GetUserAccountByUserId(userId int32) string {
	userAccount, err := us.userRepo.GetUserAccountByUserId(userId)
	global.Logger.Info("userID:::", zap.String("userAccount", userAccount))
	if err != nil {
		return "Error"
	}
	return userAccount
}

func InitUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(account string, password string) int32 {
	salt := "test"
	global.Logger.Info("userID:::", zap.String("account", account), zap.String("password", password))
	userId, err := us.userRepo.CreateNewUser(account, password, salt)
	global.Logger.Info("userID:::", zap.Int32("userId", userId))
	if err != nil {
		return response.FailCode
	}
	return response.SuccessCode
}
