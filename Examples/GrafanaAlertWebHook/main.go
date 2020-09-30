package main

import (
	"GrafanaAlertWebHook/WechatRequests"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yoyofx/yoyogo/WebFramework"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
	"io/ioutil"
	"net/http"
	"strconv"
)

func main() {
	YoyoGo.CreateDefaultBuilder(func(router Router.IRouterBuilder) {
		router.POST("/alert", func(ctx *Context.HttpContext) { // 支持Group方式
			var request AlertRequest
			_ = ctx.Bind(&request)
			var message WechatRequests.MarkdownMessage
			if len(request.EvalMatches) > 0 {
				message = WechatRequests.MarkdownMessage{
					Markdown: struct {
						Content string `json:"content" gorm:"column:content"`
					}{
						Content: request.RuleName + ",请相关同事注意。\n" +
							" > [报警次数]:<font color=\"warning\">" + strconv.Itoa(request.EvalMatches[0].Value) + "次</font>" + "\n" +
							" > [报警明细](http://jcenter-main.easypass.cn/jiankong/d/trpHG7FGk/che-hou-ye-wu-ri-zhi-cha-xun?orgId=1&from=now-30m&to=now&var-app=jishi-jishiapi&var-level=error&var-host=All&var-msg=*)\n",
					},
					Msgtype: "markdown",
				}
				msg, _ := json.Marshal(message)

				ctx.JSON(200, Context.H{
					"Message": postWechatMessage(string(msg)),
				})
				//ctx.JSON(200, message)
			}
			//msg, _ := json.Marshal(message)
			//
			//ctx.JSON(200,Context.H{
			//	"Message":postWechatMessage(string(msg)) ,
			//})
		})
	}).Build().Run() //默认端口号 :8080
}

func postWechatMessage(msg string) string {
	sendUrl := "https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=efaebe93-7b21-4bc3-888f-260744f397ac"
	client := &http.Client{}
	req, _ := http.NewRequest("POST", sendUrl, bytes.NewBuffer([]byte(msg)))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("charset", "UTF-8")
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response Status:", resp.Status)
	body, _ := ioutil.ReadAll(resp.Body)
	strBody := string(body)
	fmt.Println("response Body:", strBody)
	return strBody
}
