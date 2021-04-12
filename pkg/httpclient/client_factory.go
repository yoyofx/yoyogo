package httpclient

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
)

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
	client := NewClient()
	client.hasSelector = cf.hasDiscovery
	client.selector = cf.selector
	client.BaseUrl = baseUrl
	return client, nil
}
