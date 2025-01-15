package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/internal/model"
	"github.com/tqlong1609/go_backend_ecommerce/internal/services"
	"github.com/tqlong1609/go_backend_ecommerce/pkg/response"
)

var UserLogin = new(UserLoginController)

type UserLoginController struct{}

func (ulc *UserLoginController) Register(c *gin.Context) {
	var params model.RegisterInput

	if err := c.ShouldBindJSON(&params); err != nil {
		response.FailResponse(c, response.BadRequestCode, err.Error())
		return
	}
	data, err := services.UserLogin().Register(c, params)
	if err != nil {
		response.FailResponse(c, response.FailCode, err.Error())
		return
	}
	response.SuccessResponse(c, response.SuccessCode, data)
}
