package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	r := gin.Default()
	v1 := r.Group("/v1")
	{
		v1.GET("/ping/:name", Pong) // /v1/ping/hello
		v1.GET("/ping1", Pong1)     // /v1/ping1?id=123
	}
	return r
}

func Pong(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{
		"message": "GET: pong " + name,
	})
}

func Pong1(c *gin.Context) {
	id := c.DefaultQuery("id", "default123")
	c.JSON(http.StatusOK, gin.H{
		"message": "GET: ID " + id,
		"users":   []string{"user1", "user2", "user3"},
	})
}
