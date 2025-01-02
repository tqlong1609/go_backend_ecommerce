package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/pkg/response"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token != "token" {
			response.FailResponse(c, response.UnauthorizedCode)
			c.Abort()
			return
		}
		c.Next()
	}
}
