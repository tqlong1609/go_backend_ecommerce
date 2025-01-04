package admin

import "github.com/gin-gonic/gin"

type UserRouter struct{}

func (u *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	// private router
	userRouterPrivate := Router.Group("/admin/user")
	{
		userRouterPrivate.POST("/create_user")
		userRouterPrivate.POST("/update_user")
	}
}
