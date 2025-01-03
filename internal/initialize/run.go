package initialize

import (
	"fmt"

	"github.com/tqlong1609/go_backend_ecommerce/global"
)

func Run() {
	InitConfig()
	fmt.Println("Init config", global.Config.Mysql.Username)
	InitLogger()
	InitMySql()
	InitRedis()

	r := InitRouter()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080") - change port: r.Run(":8002")
}
