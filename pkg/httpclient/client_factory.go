package httpclient

import (
	"errors"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"regexp"
	"strings"
)

var httpExpr = "(https?)://[-A-Za-z0-9+&@#/%?=~_|!:,.;]+[-A-Za-z0-9+&@#/%=~_|]"

type IFactory interface {
	Create(baseUrl string) (*Client, error)
}

type IDiscoveryClientFactory interface {
	Create(baseUrl string) (*Client, error)
}

type Factory struct {
	selector     servicediscovery.ISelector
	hasDiscovery bool
}

func NewDiscoveryClientFactory(selector servicediscovery.ISelector) IDiscoveryClientFactory {
	return &Factory{selector: selector, hasDiscovery: true}
}

func NewFactory() IFactory {
	return &Factory{hasDiscovery: false}
}

func (cf *Factory) Create(baseUrl string) (*Client, error) {
	if !verificationBaseUrl(baseUrl) {
		return &Client{}, errors.New("Please enter the HTTP link for the specification")
	}
	client := NewClient()
	client.hasSelector = cf.hasDiscovery
	client.selector = cf.selector
	client.BaseUrl = baseUrl
	return client, nil
}

func verificationBaseUrl(baseUrl string) bool {
	if baseUrl == "" {
		return true
	}
	if !strings.HasPrefix(baseUrl, "http") {
		return false
	}
	reg := regexp.MustCompile(httpExpr)
	if !reg.MatchString(baseUrl) {
		//针对服务间调用的http调用进行特殊化处理
		if (strings.HasPrefix(baseUrl, "http://") || strings.HasPrefix(baseUrl, "https://")) && (strings.Contains(baseUrl, "[") && strings.Contains(baseUrl, "]")) {
			return true
		}
		return false
	}
	return true
}
