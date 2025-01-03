package initialize

import (
	"github.com/tqlong1609/go_backend_ecommerce/global"
	"go.uber.org/zap"
)

func Run() {
	InitConfig()
	InitLogger()
	global.Logger.Info("logger init success", zap.String("hello", "world"))
	InitMySql()
	InitRedis()

	r := InitRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") - change port: r.Run(":8002")
}
