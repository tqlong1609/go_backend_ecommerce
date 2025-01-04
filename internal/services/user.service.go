package services

import (
	"github.com/tqlong1609/go_backend_ecommerce/internal/repo"
	"github.com/tqlong1609/go_backend_ecommerce/pkg/response"
)

type IUserService interface {
	Register(email string, password string) int
}

type userService struct {
	userRepo repo.IUserRepository
}

func InitUserService(userRepo repo.IUserRepository) IUserService {
	return &userService{
		userRepo: userRepo,
	}
}

func (us *userService) Register(email string, password string) int {
	_, err := us.userRepo.GetUserByEmail(email)
	if err != nil {
		return response.FailCode
	}
	return response.SuccessCode
}
