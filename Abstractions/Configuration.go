package Abstractions

import (
	"fmt"
	"github.com/spf13/viper"
)

type Configuration struct {
	context *ConfigurationContext
	config  *viper.Viper
}

func NewConfiguration(configContext *ConfigurationContext) *Configuration {
	v := viper.New()

	//设置配置文件的名字
	v.SetConfigName(configContext.configName)
	v.AddConfigPath("./")
	v.SetConfigType(configContext.configType)

	if err := v.ReadInConfig(); err != nil {

	}

	return &Configuration{
		context: configContext,
		config:  v,
	}
}

func (c *Configuration) Get(name string) interface{} {
	return c.config.Get(name)
}

func (c *Configuration) GetSection(name string) IConfiguration {
	section := c.config.Sub(name)
	if section != nil {
		return &Configuration{config: section}
	}
	return nil
}

func (c *Configuration) Unmarshal(obj interface{}) {
	err := c.config.Unmarshal(obj)
	if err != nil {
		fmt.Println("unmarshal config is failed, err:", err)
	}
}
