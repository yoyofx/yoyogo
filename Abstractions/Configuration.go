package Abstractions

import (
	"flag"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"github.com/yoyofx/yoyogo/Utils"
)

type Configuration struct {
	context *ConfigurationContext
	config  *viper.Viper
}

func NewConfiguration(configContext *ConfigurationContext) *Configuration {

	flag.String("profile", configContext.profile, "application profile")
	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()
	_ = viper.BindPFlags(pflag.CommandLine)

	configContext.profile = viper.GetString("profile")

	configName := configContext.configName + "_" + configContext.profile
	exists, _ := Utils.PathExists("./" + configName + "." + configContext.configType)
	if !exists {
		configName = configContext.configName
	}

	defaultConfig := viper.New()
	defaultConfig.AddConfigPath(".")
	defaultConfig.SetConfigName(configName)
	defaultConfig.SetConfigType(configContext.configType)
	if err := defaultConfig.ReadInConfig(); err != nil {
		return nil
	}

	return &Configuration{
		context: configContext,
		config:  defaultConfig,
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

func (c *Configuration) GetProfile() string {
	return c.context.profile
}
