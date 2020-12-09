package consul

import (
	consul "github.com/hashicorp/consul/api"
	"github.com/yoyofx/yoyogo/web/context"
	"log"
)

type Option struct {
	ENV     *context.HostEnvironment
	Address string   `mapstructure:"address"`
	Token   string   `mapstructure:"token"`
	Tags    []string `mapstructure:"tags"`
}

type Client struct {
	consul *consul.Client
}

// NewClient returns an implementation of the Client interface, wrapping a
// concrete consul client.
func NewClient(op Option) *Client {
	config := consul.DefaultConfig()
	config.Address = op.Address
	if op.Token != "" {
		config.Token = op.Token
	}
	client, err := consul.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	return &Client{consul: client}
}

func (c *Client) Register(r *consul.AgentServiceRegistration) error {
	return c.consul.Agent().ServiceRegister(r)
}

func (c *Client) Deregister(r *consul.AgentServiceRegistration) error {
	return c.consul.Agent().ServiceDeregister(r.ID)
}

func (c *Client) GetService(service, tag string, passingOnly bool, queryOpts *consul.QueryOptions) ([]*consul.ServiceEntry, *consul.QueryMeta, error) {
	return c.consul.Health().Service(service, tag, passingOnly, queryOpts)
}
