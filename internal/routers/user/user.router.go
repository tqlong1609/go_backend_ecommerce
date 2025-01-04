package user

import (
	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/internal/controllers"
	"github.com/tqlong1609/go_backend_ecommerce/internal/repo"
	"github.com/tqlong1609/go_backend_ecommerce/internal/services"
)

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// public router

	userRepo := repo.InitUserRepository()
	userService := services.InitUserService(userRepo)
	userController := controllers.InitUserController(userService)

	userRouterPublic := Router.Group("/user")
	{
		userRouterPublic.POST("/register", userController.RegisterUser)
		userRouterPublic.POST("/login")
	}

	// private router
	userRouterPrivate := Router.Group("/user")
	{
		userRouterPrivate.POST("/get_info")
	}
}
