package wechatrequests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"io/ioutil"
	"net/http"
)

func PostWechatMessage(sendUrl, msg string) string {
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

func SendTxtMessage(request GrafanaAlertRequest, config abstractions.IConfiguration) string {
	tag := request.GetTag()
	logger := xlog.GetXLogger("wechat")
	js, _ := json.Marshal(request)
	logger.Info("Request json: %s", string(js))
	if tag == "" {
		logger.Info("no send")
		return ""
	}
	sendUrl := config.Get(fmt.Sprintf("alert.%s.webhook_url", tag)).(string)
	linkUrl := config.Get(fmt.Sprintf("alert.%s.link_url", tag)).(string)
	logger.Info("request tag:%s", tag)
	logger.Info(sendUrl)
	logger.Info(linkUrl)

	var message *MarkdownMessage
	if request.State == "alerting" && len(request.EvalMatches) > 0 {
		message = &MarkdownMessage{
			Markdown: struct {
				Content string `json:"content" gorm:"column:content"`
			}{
				Content: "## " + request.RuleName + ",请相关同事注意。\n" +
					" > [报警信息] : " + request.Message + "\n" +
					" > [报警次数] : <font color=\"warning\">" + request.GetMetricValue() + "次</font>" + "\n" +
					" > [报警明细] : (" + linkUrl + ")\n",
			},
			Msgtype: "markdown",
		}
	}
	msg, _ := json.Marshal(message)
	msgStr := string(msg)
	logger.Info("send message:%s", msgStr)

	//return sendUrl + msgStr
	return PostWechatMessage(sendUrl, msgStr)
}
