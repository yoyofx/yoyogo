package nacos

import (
	"github.com/yoyofx/yoyogo/abstractions"
	nacos_viper_remote "github.com/yoyofxteam/nacos-viper-remote"
)

func AddRemoteWithNacos(builder *abstractions.ConfigurationBuilder) *abstractions.ConfigurationBuilder {
	if builder.Context.ConfigType == "" {
		builder.Context.ConfigType = "yml"
	}
	builder.Context.EnableRemote = true
	builder.Context.RemoteProvider = nacos_viper_remote.NewRemoteProvider(builder.Context.ConfigType)
	return builder
}

func RemoteConfig(configPath string) *abstractions.Configuration {
	return AddRemoteWithNacos(abstractions.NewConfigurationBuilder().AddEnvironment().AddYamlFile(configPath)).Build()
}
