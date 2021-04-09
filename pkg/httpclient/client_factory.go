package httpclient

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
)

type ClientFactory struct {
}

func (cf *ClientFactory) CreateClient(baseUrl string) (*Client, error) {
	client := NewClient()
	client.hasSelector = false
	client.BaseUrl = baseUrl
	return client, nil
}

func (cf *ClientFactory) CreateServiceDiscoveryCleint(baseUrl string, selector servicediscovery.Selector) (*Client, error) {
	client := NewClient()
	client.hasSelector = true
	client.selector = selector
	client.BaseUrl = baseUrl
	return client, nil
}
