package apollo

import (
	"github.com/yoyofx/yoyogo/abstractions"
)

func AddRemoteWithApollo(builder *abstractions.ConfigurationBuilder) *abstractions.ConfigurationBuilder {
	if builder.Context.ConfigType == "" {
		builder.Context.ConfigType = "yml"
	}
	builder.Context.EnableRemote = true
	builder.Context.RemoteProvider = NewRemoteProvider(builder.Context.ConfigType)
	return builder
}

func RemoteConfig(configPath string) *abstractions.Configuration {
	return AddRemoteWithApollo(abstractions.NewConfigurationBuilder().AddEnvironment().AddYamlFile(configPath)).Build()
}
