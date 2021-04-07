package httpclient

import (
	"errors"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"net/http"
	"strings"
)

type ClientFactory struct {
	Selector servicediscovery.Selector
}

func (cf *ClientFactory) CreatHttpClient(url string) (*Client, error) {
	//return errors.New("url is empty"),Client{};
	if url == "" {
		return &Client{}, errors.New("url is empty")
	}
	//获取当前服务名称
	serverName := strings.Split(strings.Split(url, "[")[1], "]")[0]
	if serverName == "" {
		return &Client{}, errors.New("url don't contans serveName")
	}
	//获取服务实例
	serverInstance, err := cf.Selector.Select(serverName)
	if err != nil {
		return &Client{}, err
	}
	//根据服务名称进行url转化
	parser := servicediscovery.NewUriParser(url)
	targeUrl := parser.Generate(serverInstance.GetHost())
	//创建http 客户端
	client := NewClient()
	client.Request = &Request{
		url:    targeUrl,
		method: "GET",
		header: http.Header{},
	}
	return client, nil
}
