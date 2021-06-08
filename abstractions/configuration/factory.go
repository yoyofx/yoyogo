package configuration

import "github.com/yoyofx/yoyogo/abstractions"

// YAML config by yaml or yml file
func YAML(configPath string) *abstractions.Configuration {
	return abstractions.NewConfigurationBuilder().AddEnvironment().AddYamlFile(configPath).Build()
}

// JSON config by json file
func JSON(configPath string) *abstractions.Configuration {
	return abstractions.NewConfigurationBuilder().AddEnvironment().AddJsonFile(configPath).Build()
}

// PROPERITES config by properties file
func PROPERITES(configPath string) *abstractions.Configuration {
	return abstractions.NewConfigurationBuilder().AddEnvironment().AddPropertiesFile(configPath).Build()
}
