package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/internal/services"
)

type PongController struct {
	pongService *services.PongService
}

func InitPongController() *PongController {
	return &PongController{
		pongService: services.InitPongService(),
	}
}

func (pc *PongController) GetPong(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": pc.pongService.GetPongService(),
	})
}
