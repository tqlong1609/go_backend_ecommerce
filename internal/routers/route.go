package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/internal/controllers"
)

func Init() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping/:name", controllers.InitPongController().GetPong) // /v1/ping/hello
		v1.GET("/ping1", controllers.InitPong1Controller().GetPong1)    // /v1/ping1?id=123
	}
	return r
}
