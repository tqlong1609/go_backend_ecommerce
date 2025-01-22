package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/internal/controllers"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router
	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register-email", controllers.UserLogin.RegisterWithEmail)
		userRouterPublic.POST("/login")
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/:id")
	}
}
