package initialize

import (
	"github.com/tqlong1609/go_backend_ecommerce/global"
	"github.com/tqlong1609/go_backend_ecommerce/internal/database"
	"github.com/tqlong1609/go_backend_ecommerce/internal/services"
	"github.com/tqlong1609/go_backend_ecommerce/internal/services/implements"
)

func InitServiceInterface() {
	queries := database.New(global.Mdbc)
	services.InitUserLogin(implements.InitUserLoginImplement(queries))
}
