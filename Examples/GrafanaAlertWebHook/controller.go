package main

import (
	"GrafanaAlertWebHook/WechatRequests"
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
)

func PostAlert(ctx *Context.HttpContext) {
	var request WechatRequests.GrafanaAlertRequest
	_ = ctx.Bind(&request)
	var config Abstractions.IConfiguration
	_ = ctx.RequiredServices.GetService(&config)

	ctx.JSON(200, Context.H{
		"Message": WechatRequests.SendTxtMessage(request, config),
	})
}
