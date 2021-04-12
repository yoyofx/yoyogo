package httpclient

import (
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
)

type Factory struct {
	selector     servicediscovery.Selector
	hasDiscovery bool
}

func NewDiscoveryClientFactory(selector servicediscovery.Selector) *Factory {
	return &Factory{selector: selector, hasDiscovery: true}
}

func NewFactory() *Factory {
	return &Factory{hasDiscovery: false}
}

func (cf *Factory) Create(baseUrl string) (*Client, error) {
	client := NewClient()
	client.hasSelector = cf.hasDiscovery
	client.selector = cf.selector
	client.BaseUrl = baseUrl
	return client, nil
}
