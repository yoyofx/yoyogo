package configuration

import "github.com/yoyofx/yoyogo/abstractions"

func LocalConfig(configPath string) *abstractions.Configuration {
	return abstractions.NewConfigurationBuilder().AddEnvironment().AddYamlFile(configPath).Build()
}
