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
	defaultConfig := viper.New()
	defaultConfig.AddConfigPath(".")
	defaultConfig.SetConfigName(configContext.configName)
	defaultConfig.SetConfigType(configContext.configType)
	if err := defaultConfig.ReadInConfig(); err != nil {
		return nil
	}

	profile := defaultConfig.Get("application.profile")
	var profileConfig *viper.Viper
	if profile != nil {
		profileConfig = viper.New()
		profileConfig.AddConfigPath(".")
		profileConfig.SetConfigName(configContext.configName + "_" + profile.(string))
		profileConfig.SetConfigType(configContext.configType)
		configs := defaultConfig.AllSettings()
		// 将default中的配置全部以默认配置写入
		for k, v := range configs {
			profileConfig.Set(k, v)
		}

		if err := profileConfig.ReadInConfig(); err != nil {
			profileConfig = defaultConfig
		}
	}

	return &Configuration{
		context: configContext,
		config:  profileConfig,
	}
}

func (c *Configuration) Get(name string) interface{} {
	return c.config.Get(name)
}

func (c *Configuration) GetSection(name string) IConfiguration {
	section := c.config.Sub(name)

	configs := c.config.AllSettings()
	// 将default中的配置全部以默认配置写入
	for k, v := range configs {
		section.Set(k, v)
	}

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
