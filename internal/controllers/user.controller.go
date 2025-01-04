package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/internal/services"
	"github.com/tqlong1609/go_backend_ecommerce/pkg/response"
)

type UserController struct {
	userService services.IUserService
}

func InitUserController(userService services.IUserService) *UserController {
	return &UserController{
		userService,
	}
}

func (uc *UserController) RegisterUser(c *gin.Context) {
	const email = "email"
	const password = "password"
	response.SuccessResponse(c, response.SuccessCode, uc.userService.Register(email, password))
}
