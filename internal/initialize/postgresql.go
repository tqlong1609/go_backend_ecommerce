package initialize

import (
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/tqlong1609/go_backend_ecommerce/global"
	"github.com/tqlong1609/go_backend_ecommerce/internal/model"
	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func checkErrorPanic(err error, msg string) {
	if err != nil {
		global.Logger.Error(msg, zap.Error(err))
		panic(err)
	}
}

func InitPostgresql() {

	pg := global.Config.PostgreSQL

	// sslmode=disable: Vô hiệu hóa SSL (phù hợp cho development). Trong production, bạn nên cấu hình chế độ SSL an toàn
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		pg.Host, pg.Port, pg.Username, pg.Password, pg.Database)

	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  psqlInfo,
		PreferSimpleProtocol: true, // yêu cầu driver PostgreSQL sử dụng giao thức đơn giản, giảm tải một số tính năng nâng cao nhưng tăng hiệu suất trong một số trường hợp
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		// Nếu đặt là true, GORM sẽ bỏ qua việc tạo giao dịch mặc định (default transaction) khi thực thi các truy vấn

		// Nếu đặt là false, nghĩa là GORM sẽ tự động tạo giao dịch mặc định cho các thao tác như Create, Update, hoặc Delete. Điều này đảm bảo an toàn trong các trường hợp xảy ra lỗi, nhưng có thể ảnh hưởng đến hiệu suất
	})

	checkErrorPanic(err, "PostgreSQL init error")
	global.Logger.Info("PostgreSQL init successfully")
	global.Mdb = db
	// defer db.Close()

	// Connection Pool
	setConnectionPool()

	// Migrate
	migrateTables()
	genTableDAO()
}

func genTableDAO() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db
	// g.GenerateAllTable()
	// g.GenerateModel("users")
	// g.GenerateModel("categories")

	// Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})

	// Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()
}

// được sử dụng để tối ưu hóa việc quản lý pool kết nối với cơ sở dữ liệu
func setConnectionPool() {
	pg := global.Config.PostgreSQL

	// pgDB: Là đối tượng *sql.DB, cung cấp các phương thức quản lý pool kết nối cơ sở dữ liệu
	pgDB, err := global.Mdb.DB()

	checkErrorPanic(err, "PostgreSQL set connection pool error")

	/*
		MaxIdleConns: Số lượng kết nối nhàn rỗi tối đa
		MaxOpenConns: Số lượng kết nối tối đa được mở cùng một lúc
		MaxLifeTime: Thời gian tối đa mà một kết nối có thể được sử dụng trong pool
	*/

	// Thiết lập thời gian tối đa mà một kết nối có thể nằm trong pool ở trạng thái nhàn rỗi trước khi nó bị đóng
	pgDB.SetConnMaxIdleTime(time.Duration(pg.MaxIdleConns))
	// Thiết lập số lượng tối đa các kết nối có thể được mở đồng thời
	pgDB.SetMaxOpenConns(pg.MaxOpenConns)
	// Thiết lập thời gian tối đa mà một kết nối có thể được sử dụng trong pool
	// Nếu thời gian này đã trôi qua, kết nối sẽ được đóng và loại bỏ khỏi pool
	pgDB.SetConnMaxLifetime(time.Duration(pg.MaxLifeTime))
}

func migrateTables() {
	// err := global.Mdb.AutoMigrate(&po.User{}, &po.Role{})
	err := global.Mdb.AutoMigrate(&model.UsersV2{})
	checkErrorPanic(err, "PostgreSQL migrate tables error")
	global.Logger.Info("PostgreSQL migrate tables successfully")
}
