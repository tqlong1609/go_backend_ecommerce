package global

import (
	"database/sql"

	"github.com/redis/go-redis/v9"
	"github.com/tqlong1609/go_backend_ecommerce/pkg/logger"
	"github.com/tqlong1609/go_backend_ecommerce/pkg/setting"
	"gorm.io/gorm"
)

var (
	Rdb    *redis.Client
	Mdb    *gorm.DB
	Mdbc   *sql.DB
	Config *setting.Config
	Logger *logger.LoggerZap
)
