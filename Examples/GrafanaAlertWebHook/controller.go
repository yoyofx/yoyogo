package main

import (
	"GrafanaAlertWebHook/WechatRequests"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"strconv"
)

func PostAlert(ctx *Context.HttpContext) {
	var request WechatRequests.GrafanaAlertRequest
	_ = ctx.Bind(&request)
	var message WechatRequests.MarkdownMessage
	if request.State == "alerting" && len(request.EvalMatches) > 0 {
		message = WechatRequests.MarkdownMessage{
			Markdown: struct {
				Content string `json:"content" gorm:"column:content"`
			}{
				Content: request.RuleName + ",请相关同事注意。\n" +
					" > [报警次数]:<font color=\"warning\">" + strconv.Itoa(request.EvalMatches[0].Value) + "次</font>" + "\n" +
					" > [报警明细](http://jcenter-main.easypass.cn/jiankong/d/trpHG7FGk/che-hou-ye-wu-ri-zhi-cha-xun?orgId=1&from=now-1h&to=now&var-app=jishi*&var-level=error&var-host=All&var-msg=*)\n",
			},
			Msgtype: "markdown",
		}
		//msg, _ := json.Marshal(message)
		//sendUrl := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=efaebe93-7b21-4bc3-888f-260744f397ac"
		//ctx.JSON(200, Context.H{
		//	"Message": postWechatMessage(sendUrl,string(msg)),
		//})
		ctx.JSON(200, message)
	}
	//msg, _ := json.Marshal(message)
	//
	//ctx.JSON(200,Context.H{
	//	"Message":postWechatMessage(string(msg)) ,
	//})
}
