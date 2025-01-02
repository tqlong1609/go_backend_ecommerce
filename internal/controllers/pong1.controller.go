package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/internal/services"
	"github.com/tqlong1609/go_backend_ecommerce/pkg/response"
)

type Pong1Controller struct {
	pong1Service *services.Pong1Service
}

func InitPong1Controller() *Pong1Controller {
	return &Pong1Controller{
		pong1Service: services.InitPong1Service(),
	}
}

func (pc *Pong1Controller) GetPong1(c *gin.Context) {
	fmt.Println("My Handler ----")
	response.SuccessResponse(c, response.SuccessCode, []string{"hello1", "hello2"})
}
