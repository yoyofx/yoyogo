package contollers

import (
	"github.com/yoyofx/yoyogo/pkg/cache/redis"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"gorm.io/gorm"
)

type DbController struct {
	mvc.ApiController
}

func NewDbController() *DbController {
	return &DbController{}
}

func (controller DbController) GetMysql(ctx *context.HttpContext) mvc.ApiResult {
	var db *gorm.DB
	_ = ctx.RequiredServices.GetService(&db)
	// 原生 SQL
	var field0 int64
	row := db.Raw("select 1024").Row()
	_ = row.Scan(&field0)

	return controller.OK(context.H{"select": field0})
}

func (controller DbController) GetRedis(ctx *context.HttpContext) mvc.ApiResult {
	var client redis.IClient
	_ = ctx.RequiredServices.GetService(&client)

	strv, err := client.GetKVOps().GetString("dcctor1")
	if err == nil {
		return controller.OK(context.H{"redis key: dcctor1": strv})
	}
	return controller.Fail(err.Error())
}
