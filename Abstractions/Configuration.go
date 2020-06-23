package Abstractions

import (
	"github.com/spf13/viper"
)

type Configuration struct {
	context *ConfigurationContext
	config  *viper.Viper
}

func NewConfiguration(configContext *ConfigurationContext) *Configuration {
	v := viper.New()

	return &Configuration{
		context: configContext,
		config:  v,
	}
}

func (c *Configuration) Get(name string) interface{} {
	return c.config.Get(name)
}

func (c *Configuration) GetSection(name string) *Configuration {
	section := c.config.Sub(name)
	if section != nil {
		return &Configuration{config: section}
	}
	return nil
}
