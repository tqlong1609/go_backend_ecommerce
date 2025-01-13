package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/global"
	"github.com/tqlong1609/go_backend_ecommerce/internal/services"
	"github.com/tqlong1609/go_backend_ecommerce/pkg/response"
	"go.uber.org/zap"
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
	var body struct {
		Account  string `json:"account"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&body); err != nil {
		response.FailResponse(c, response.BadRequestCode, err.Error())
		return
	}
	response.SuccessResponse(c, response.SuccessCode, uc.userService.Register(body.Account, body.Password))
}

func (uc *UserController) GetUserAccountByUserId(c *gin.Context) {
	userId, _ := strconv.ParseInt(c.Param("id"), 10, 32)
	global.Logger.Info("userID:::", zap.Int64("account", userId))

	response.SuccessResponse(c, response.SuccessCode, uc.userService.GetUserAccountByUserId(int32(userId)))
}
