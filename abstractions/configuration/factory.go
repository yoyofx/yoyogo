package configuration

import "github.com/yoyofx/yoyogo/abstractions"

func YAML(configPath string) *abstractions.Configuration {
	return abstractions.NewConfigurationBuilder().AddEnvironment().AddYamlFile(configPath).Build()
}

func NACOS(configPath string) *abstractions.Configuration {
	return abstractions.NewConfigurationBuilder().AddEnvironment().AddYamlFile(configPath).AddRemoteWithNacos().Build()
}
