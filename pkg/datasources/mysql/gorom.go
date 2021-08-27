package mysql

import (
	"database/sql"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

func NewGormDb(source *MySqlDataSource) *gorm.DB {
	timestr := time.Now().Format("2006/01/02 - 15:04:05.00")
	logPrefix := fmt.Sprintf("%s - [YOYOGO] - [DEBUG] ", timestr)
	dbLogger := logger.New(
		log.New(os.Stdout, logPrefix, log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  false,         // 禁用彩色打印
		},
	)

	conn, _, _ := source.Open()
	sqlDB := conn.(*sql.DB)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{
		Logger: dbLogger,
	})

	if err != nil {
		panic(err)
	}

	if source.isDebug {
		return gormDB.Debug()
	}

	return gormDB
}
