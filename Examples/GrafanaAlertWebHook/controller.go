package main

import (
	"GrafanaAlertWebHook/WechatRequests"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/web/context"
)

func PostAlert(ctx *context.HttpContext) {
	var request WechatRequests.GrafanaAlertRequest
	_ = ctx.Bind(&request)
	var config abstractions.IConfiguration
	_ = ctx.RequiredServices.GetService(&config)

	ctx.JSON(200, context.H{
		"Message": WechatRequests.SendTxtMessage(request, config),
	})
}
