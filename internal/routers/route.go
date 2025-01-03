package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/internal/controllers"
	"github.com/tqlong1609/go_backend_ecommerce/internal/middlewares"
)

func Init() *gin.Engine {
	r := gin.Default()
	r.Use(middlewares.AuthMiddleware())
	v1 := r.Group("/v1")
	{
		v1.GET("/ping/:name", controllers.InitPongController().GetPong)
		v1.GET("/ping1", controllers.InitPong1Controller().GetPong1)
	}
	return r
}
