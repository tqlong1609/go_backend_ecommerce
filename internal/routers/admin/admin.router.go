package admin

import "github.com/gin-gonic/gin"

type AdminRouter struct{}

func (a *AdminRouter) InitAdminRouter(Router *gin.RouterGroup) {
	// public router
	adminRouterPublic := Router.Group("/admin")
	{
		adminRouterPublic.POST("/login")
	}
	// private router
	adminRouterPrivate := Router.Group("/admin")
	{
		adminRouterPrivate.POST("/get_info")
	}
}
