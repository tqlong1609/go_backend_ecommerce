package initialize

import (
	"github.com/tqlong1609/go_backend_ecommerce/global"
	"github.com/tqlong1609/go_backend_ecommerce/pkg/logger"
)

func InitLogger() {
	global.Logger = logger.InitLogger(global.Config.Logger)
}
