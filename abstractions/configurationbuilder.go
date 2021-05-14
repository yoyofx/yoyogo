package abstractions

import (
	"github.com/yoyofx/yoyogo/abstractions/hostenv"
	nacos_viper_remote "github.com/yoyofxteam/nacos-viper-remote"
)

type ConfigurationContext struct {
	enableFlag     bool
	enableEnv      bool
	configDir      string
	configType     string
	configName     string
	profile        string
	configFile     string
	enableRemote   bool
	remoteProvider IConfigurationRemoteProvider
}

type ConfigurationBuilder struct {
	context *ConfigurationContext
}

func NewConfigurationBuilder() *ConfigurationBuilder {
	return &ConfigurationBuilder{context: &ConfigurationContext{}}
}

func (builder *ConfigurationBuilder) AddFlagArgs() *ConfigurationBuilder {
	builder.context.enableFlag = true
	return builder
}

func (builder *ConfigurationBuilder) AddEnvironment() *ConfigurationBuilder {
	builder.context.enableEnv = true
	return builder
}

func (builder *ConfigurationBuilder) AddYamlFile(name string) *ConfigurationBuilder {
	if builder.context.configType == "" {
		builder.context.configType = "yml"
		builder.context.configName = name
	}
	return builder
}

func (builder *ConfigurationBuilder) AddJsonFile(name string) *ConfigurationBuilder {
	if builder.context.configType == "" {
		builder.context.configType = "json"
		builder.context.configName = name
	}
	return builder
}

func (builder *ConfigurationBuilder) AddRemoteWithNacos() *ConfigurationBuilder {
	if builder.context.configType == "" {
		builder.context.configType = "yml"
	}
	builder.context.enableRemote = true
	builder.context.remoteProvider = nacos_viper_remote.NewRemoteProvider(builder.context.configType)
	return builder
}

func (builder *ConfigurationBuilder) BuildEnv(env string) *Configuration {
	builder.context.profile = env
	return NewConfiguration(builder.context)
}

func (builder *ConfigurationBuilder) Build() *Configuration {
	builder.context.profile = hostenv.Dev
	builder.context.enableFlag = true
	return NewConfiguration(builder.context)
}
