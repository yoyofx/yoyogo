package mysql

import (
	"database/sql"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewGormDb(source *MySqlDataSource) *gorm.DB {
	conn, _, _ := source.Open()
	sqlDB := conn.(*sql.DB)
	gormDB, err := gorm.Open(mysql.New(mysql.Config{
		Conn: sqlDB,
	}), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	return gormDB
}
