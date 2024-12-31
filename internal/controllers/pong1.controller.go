package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/internal/services"
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
	c.JSON(http.StatusOK, gin.H{
		"message": pc.pong1Service.GetPongService(),
		"users":   []string{"user1", "user2", "user3"},
	})
}
