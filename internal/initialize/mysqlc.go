package initialize

import (
	"database/sql"
	"example.com/go-ecommerce-backend-api/global"
	"fmt"
	"go.uber.org/zap"
	"gorm.io/gen"
	"time"
)

func checkErrorPanicC(err error, errString string) {
	if err != nil {
		global.Logger.Error(errString, zap.Error(err))
		panic(err)
	}
}

func InitMysqlC() {
	m := global.Config.MySQL
	dsn := "%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	var s = fmt.Sprintf(dsn, m.Username, m.Password, m.Host, m.Port, m.Dbname)
	db, err := sql.Open("mysql", s)
	checkErrorPanicC(err, "Failed to connect to MySQL database")
	global.Logger.Info("initialize MySQL database success")
	global.Mdbc = db
	// set Pool
	SetPoolC()
	//genTableDAO()
	//MigrateTablesC()

}

func SetPoolC() {
	m := global.Config.MySQL
	sqlDb, err := global.Mdb.DB()
	if err != nil {
		fmt.Println("mysql error: %s::", err)
	}
	sqlDb.SetMaxIdleConns(int(time.Duration(m.MaxIdleConns)))
	sqlDb.SetMaxOpenConns(m.MaxOpenConns)
	sqlDb.SetConnMaxLifetime(time.Duration(m.ConnMaxLifetime))

}

func migrateTablesC() {
	err := global.Mdb.AutoMigrate(
	//&po.User{},
	//&po.Role{},
	//model.GoCrmUserV2{},
	)
	if err != nil {
		fmt.Println("migrate tables error: %s::", err)
	}
}

func genTableDAOC() {
	g := gen.NewGenerator(gen.Config{
		OutPath: "./internal/model",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	// // gormdb, _ := gorm.Open(mysql.Open("root:@(127.0.0.1:3306)/demo?charset=utf8mb4&parseTime=True&loc=Local"))
	g.UseDB(global.Mdb) // reuse your gorm db
	// g.GenerateAllTable()
	g.GenerateModel("go_crm_user")
	// // Generate basic type-safe DAO API for struct `model.User` following conventions
	// g.ApplyBasic(model.User{})

	// // Generate Type Safe API with Dynamic SQL defined on Querier interface for `model.User` and `model.Company`
	// g.ApplyInterface(func(Querier) {}, model.User{}, model.Company{})

	// Generate the code
	g.Execute()

}
