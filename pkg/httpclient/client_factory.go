package httpclient

import (
	"errors"
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"net/http"
	"strings"
)

var baseUrlMap = make(map[string]string)

type ClientOptions struct {
	clientName string
}

type ClientFactory struct {
	Selector servicediscovery.Selector
}

func AddHttpClient(clientName string, action func(options *ClientOptions)) {
	action(&ClientOptions{
		clientName: clientName,
	})
}

func (options *ClientOptions) AddClientBaseUrl(baseUrl string) {
	baseUrlMap[options.clientName] = baseUrl
}

func (cf *ClientFactory) CreateClient(clientName string) (*Client, error) {
	client := NewClient()
	if clientName != "" {
		baseUrl := baseUrlMap[clientName]
		if baseUrl == "" {
			return nil, errors.New("can't find this clientName")
		}
		client.BaseUrl = baseUrl
		return client, nil
	}
	return client, nil
}

func (cf *ClientFactory) BuilderSelectorRequest(url string, method string) (*Request, error) {
	//return errors.New("url is empty"),Client{};
	if url == "" {
		return &Request{}, errors.New("url is empty")
	}
	//获取当前服务名称
	serverName := strings.Split(strings.Split(url, "[")[1], "]")[0]
	if serverName == "" {
		return &Request{}, errors.New("url don't contans serveName")
	}
	//获取服务实例
	serverInstance, err := cf.Selector.Select(serverName)
	if err != nil {
		return &Request{}, err
	}
	//根据服务名称进行url转化
	parser := servicediscovery.NewUriParser(url)
	tagUrl := parser.Generate(fmt.Sprintf("%s:%v", serverInstance.GetHost(), serverInstance.GetPort()))
	//创建http 客户端
	req := &Request{
		url:     tagUrl,
		method:  method,
		header:  http.Header{},
		timeout: 5,
	}
	return req, nil
}
