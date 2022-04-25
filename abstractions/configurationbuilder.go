package abstractions

import (
	"github.com/spf13/viper"
	"github.com/yoyofx/yoyogo/abstractions/hostenv"
)

type ConfigurationContext struct {
	enableFlag          bool
	enableEnv           bool
	configDir           string
	ConfigType          string
	configName          string
	profile             string
	configFile          string
	EnableRemote        bool
	RemoteProvider      IConfigurationRemoteProvider
	decoderConfigOption viper.DecoderConfigOption
}

type ConfigurationBuilder struct {
	Context *ConfigurationContext
}

func NewConfigurationBuilder() *ConfigurationBuilder {
	return &ConfigurationBuilder{Context: &ConfigurationContext{}}
}

func (builder *ConfigurationBuilder) AddFlagArgs() *ConfigurationBuilder {
	builder.Context.enableFlag = true
	return builder
}

func (builder *ConfigurationBuilder) AddEnvironment() *ConfigurationBuilder {
	builder.Context.enableEnv = true
	return builder
}

func (builder *ConfigurationBuilder) AddYamlFile(name string) *ConfigurationBuilder {
	if builder.Context.ConfigType == "" {
		builder.Context.ConfigType = "yml"
		builder.Context.configName = name
	}
	return builder
}

func (builder *ConfigurationBuilder) AddJsonFile(name string) *ConfigurationBuilder {
	if builder.Context.ConfigType == "" {
		builder.Context.ConfigType = "json"
		builder.Context.configName = name
	}
	return builder
}

func (builder *ConfigurationBuilder) AddPropertiesFile(name string) *ConfigurationBuilder {
	if builder.Context.ConfigType == "" {
		builder.Context.ConfigType = "prop"
		builder.Context.configName = name
	}
	return builder
}

func (builder *ConfigurationBuilder) BuildEnv(env string) *Configuration {
	builder.Context.profile = env
	return NewConfiguration(builder.Context)
}

func (builder *ConfigurationBuilder) Build() *Configuration {
	builder.Context.profile = hostenv.Dev
	builder.Context.enableFlag = true
	return NewConfiguration(builder.Context)
}
