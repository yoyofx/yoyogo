package tests

import (
	"encoding/json"
	"gopkg.in/go-playground/assert.v1"
	"grafanaalertwebhook/wechatrequests"
	"testing"
)

func TestMessage(t *testing.T) {

	var message *wechatrequests.MarkdownMessage
	message = &wechatrequests.MarkdownMessage{
		Markdown: struct {
			Content string `json:"content" gorm:"column:content"`
		}{
			Content: "## " + "test" + ",请相关同事注意。\n" +
				" > [报警信息] : " + "test" + "\n" +
				" > [报警次数] : <font color=\"warning\">" + "0" + "次</font>" + "\n" +
				" > [报警明细] : (" + "app" + ")\n",
		},
		Msgtype: "markdown",
	}

	msg, _ := json.Marshal(message)
	msgStr := string(msg)

	//return sendUrl + msgStr
	dd := wechatrequests.PostWechatMessage("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=93fb0726-9794-49f5-b10f-dac3c85396da", msgStr)

	assert.Equal(t, dd != "", true)
}
