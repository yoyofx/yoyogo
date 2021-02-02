package mysql

import (
	"github.com/yoyofx/yoyogo/abstractions/health"
	"gorm.io/gorm"
)

type MySQLHealthIndicator struct {
	db *gorm.DB
}

func NewMysqlHealthIndicator(db *gorm.DB) *MySQLHealthIndicator {
	return &MySQLHealthIndicator{db: db}
}

func (h *MySQLHealthIndicator) Health() health.ComponentStatus {
	status := health.Up("mysqlHealth")
	var field0 int64
	sql := "select 1024"
	row := h.db.Raw(sql).Row()
	_ = row.Scan(&field0)
	if field0 <= 0 {
		status.SetStatus("down")
	}
	return status.WithDetail("sql", sql)
}
