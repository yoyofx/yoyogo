package WechatRequests

import (
	"bytes"
	"fmt"
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
