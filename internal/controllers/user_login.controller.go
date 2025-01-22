package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/internal/model"
	"github.com/tqlong1609/go_backend_ecommerce/internal/services"
	"github.com/tqlong1609/go_backend_ecommerce/pkg/response"
)

var UserLogin = new(UserLoginController)

type UserLoginController struct{}

// register with email
func (ulc *UserLoginController) RegisterWithEmail(c *gin.Context) {
	var params model.RegisterWithEmailInput
	if err := c.ShouldBindJSON(&params); err != nil {
		response.FailResponse(c, response.BadRequestCode, err.Error())
		return
	}
	err := services.UserLogin().RegisterWithEmail(c, params)
	if err != nil {
		response.FailResponse(c, response.FailCode, err.Error())
		return
	}
	response.SuccessResponse(c, response.SuccessCode, nil)
	return
}
