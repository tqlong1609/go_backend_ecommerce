package initialize

import (
	"fmt"

	"github.com/tqlong1609/go_backend_ecommerce/global"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func InitMySql() {
	m := global.Config.Mysql
	dsn := "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Database)
	db, err := gorm.Open(mysql.Open(s), &gorm.Config{
		SkipDefaultTransaction: false,
	})

	checkErrorPanic(err, "mysql init error")
	global.Logger.Info("mysql init success")
	global.Mdb = db

	// Set Pool

}
