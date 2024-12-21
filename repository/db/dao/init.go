package dao

import (
	"context"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"strings"
	"time"
	"todo-list/config"
)

var _db *gorm.DB

func MysqlInit() {
	mysqlConfig := config.Config.MySql["default"]
	dsn := strings.Join([]string{mysqlConfig.UserName, ":", mysqlConfig.Password,
		"@tcp(", mysqlConfig.DbHost, ":", mysqlConfig.DbPort, ")/",
		mysqlConfig.DbName,
		"?charset=",
		mysqlConfig.Charset,
		"&parseTime=true"},
		"")

	var ormLogger logger.Interface

	if gin.Mode() == "debug" {
		ormLogger = logger.Default.LogMode(logger.Info)
	}

	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,  // string类型默认字段
		DisableDatetimePrecision:  true, // 禁用datetime精度
		DontSupportRenameIndex:    true, // 重命名索引的时候， 采用删除并重建的方式
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: false,
	}), &gorm.Config{
		Logger: ormLogger,
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, // 表名不使用复数
		},
	})

	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	sqlDB.SetMaxIdleConns(20)  // 连接池
	sqlDB.SetMaxOpenConns(100) // 最大打开数量
	sqlDB.SetConnMaxLifetime(time.Second * 30)
	_db = db
	//migrate()
}

func NewDBClient(ctx context.Context) *gorm.DB {
	db := _db
	return db.WithContext(ctx)
}

//func migrate() {
//	err := _db.Set("gorm:table_options", "charset=utf8mb4").
//		AutoMigrate(&model.UserModel{}, &model.TaskModel{})
//
//	if err != nil {
//		panic(err)
//	}
//}
