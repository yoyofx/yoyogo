package WechatRequests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"io/ioutil"
	"net/http"
)

func postWechatMessage(sendUrl, msg string) string {
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

func SendTxtMessage(request GrafanaAlertRequest, config Abstractions.IConfiguration) string {
	tag := request.GetTag()
	logger := XLog.GetXLogger("wechat")
	js, _ := json.Marshal(request)
	logger.Info(string(js))
	if tag == "" {
		logger.Info("no send")
		return ""
	}
	sendUrl := config.Get("alert.webhook_url").(string)
	linkUrl := config.Get(fmt.Sprintf("alert.%s.link_url", tag)).(string)
	var message *MarkdownMessage
	if request.State == "alerting" && len(request.EvalMatches) > 0 {
		message = &MarkdownMessage{
			Markdown: struct {
				Content string `json:"content" gorm:"column:content"`
			}{
				Content: request.RuleName + ",请相关同事注意。\n" +
					" > [报警次数]:<font color=\"warning\">" + request.GetMetricValue() + "次</font>" + "\n" +
					" > [报警明细](" + linkUrl + ")\n",
			},
			Msgtype: "markdown",
		}
	}
	msg, _ := json.Marshal(message)
	msgStr := string(msg)
	logger.Info("send message:%s", msgStr)

	return postWechatMessage(sendUrl, msgStr)
}
