package global

import (
	"github.com/tqlong1609/go_backend_ecommerce/pkg/logger"
	"github.com/tqlong1609/go_backend_ecommerce/pkg/setting"
	"gorm.io/gorm"
)

var (
	Mdb    *gorm.DB
	Config *setting.Config
	Logger *logger.LoggerZap
)
