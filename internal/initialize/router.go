package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/internal/controllers"
	"github.com/tqlong1609/go_backend_ecommerce/internal/middlewares"
)

func AAA() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before AAA")
		c.Next()
		fmt.Println("After AAA")
	}
}

func BBB() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("Before BBB")
		c.Next()
		fmt.Println("After BBB")
	}
}

func CCC(c *gin.Context) {
	fmt.Println("Before CCC")
	c.Next()
	fmt.Println("After CCC")
}

func InitRouter() *gin.Engine {
	r := gin.Default()
	// r.Use(AAA(), BBB(), CCC)
	r.Use(middlewares.AuthMiddleware(), BBB(), CCC)
	v1 := r.Group("/v1")
	{
		v1.GET("/ping/:name", controllers.InitPongController().GetPong) // /v1/ping/hello
		v1.GET("/ping1", controllers.InitPong1Controller().GetPong1)    // /v1/ping1?id=123
	}
	return r
}
