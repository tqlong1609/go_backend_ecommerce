package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/internal/controllers"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router

	// userRepo := repo.InitUserRepository()
	// userService := services.InitUserService(userRepo)
	// userController := controllers.InitUserController(userService)
	// userController, _ := wire.InitUserRouterHandler()

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", controllers.UserLogin.Register)
		userRouterPublic.POST("/login")
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.GET("/:id")
	}
}
