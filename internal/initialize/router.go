package initialize

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/tqlong1609/go_backend_ecommerce/global"
	"github.com/tqlong1609/go_backend_ecommerce/internal/routers"
)

func InitRouter() *gin.Engine {
	severConfig := global.Config.Server
	var r *gin.Engine
	global.Logger.Info(fmt.Sprintf("Server is running in %s mode", severConfig.Mode))
	if severConfig.Mode == "dev" {
		gin.SetMode(gin.DebugMode)
		gin.ForceConsoleColor()
		r = gin.Default()
	} else {
		gin.SetMode(gin.ReleaseMode)
		r = gin.New()
	}

	// ### middleware ###
	// logging
	// cross
	// limiter global
	// ### End ###

	adminRouter := routers.RouterGroupApp.Admin
	userRouter := routers.RouterGroupApp.User

	mainGroup := r.Group("/v1")
	{
		mainGroup.GET("/ping") // tracking monitor
	}
	{
		userRouter.InitUserRouter(mainGroup)
		userRouter.InitProductRouter(mainGroup)
	}
	{
		adminRouter.InitUserRouter(mainGroup)
		adminRouter.InitAdminRouter(mainGroup)
	}

	return r
}
