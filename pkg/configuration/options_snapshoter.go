package configuration

import "github.com/yoyofx/yoyogo/abstractions"

type OptionsSnapshot[T any] struct {
	config      abstractions.IConfiguration
	sectionName string
	value       T
}

func (options OptionsSnapshot[T]) CurrentValue() T {
	var configObject T
	options.config.GetConfigObject(options.sectionName, &configObject)
	return configObject
}
