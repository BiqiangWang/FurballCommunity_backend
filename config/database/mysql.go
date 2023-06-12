package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var (
	DB    *gorm.DB
	DbErr error
)

func InitMySQL() {
	// ----------------------- 日志设置 -----------------------
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			Colorful:                  false,         // Disable color
		},
	)
	// ----------------------- 连接数据库 -----------------------
	dsn := "root:Wbq123456!@tcp(47.113.230.181:3306)/fc?charset=utf8mb4&parseTime=True&loc=Local"
	DB, DbErr = gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,
		DefaultStringSize:         256,
		DisableDatetimePrecision:  true,
		DontSupportRenameIndex:    true,
		DontSupportRenameColumn:   true,
		SkipInitializeWithVersion: true,
	}), &gorm.Config{
		Logger: newLogger,
	})
	if DbErr != nil {
		log.Fatalf("mysql connect error %v", DbErr)
	}
	//DB = db
	//if err != nil {
	//	fmt.Printf("mysql connect error %v", err)
	//	panic(err)
	//}
	if DB.Error != nil {
		fmt.Printf("database error %v", DB.Error)
	}
}
