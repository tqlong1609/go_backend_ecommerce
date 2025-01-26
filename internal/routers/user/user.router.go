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
		userRouterPublic.POST("/verify-otp", controllers.UserLogin.VerifyOTP)
		userRouterPublic.POST("/login")
		userRouterPublic.POST("/complete-registration", controllers.UserLogin.CompleteRegistration)
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/:id")
	}
}
