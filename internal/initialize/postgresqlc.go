package initialize

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/tqlong1609/go_backend_ecommerce/global"
)

func InitPostgresqlC() {

	pg := global.Config.PostgreSQL

	// sslmode=disable: Vô hiệu hóa SSL (phù hợp cho development). Trong production, bạn nên cấu hình chế độ SSL an toàn
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pg.Host, pg.Port, pg.Username, pg.Password, pg.Database)

	db, err := sql.Open("postgres", psqlInfo)

	checkErrorPanic(err, "PostgreSQL init error (SqlC)")
	global.Logger.Info("PostgreSQL init successfully (SqlC)")
	global.Mdbc = db
	// defer db.Close()
}
