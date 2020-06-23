package Abstract

type ConfigurationContext struct {
	enableFlag bool
	enableEnv  bool
	configType string
	configName string
}

type ConfigurationBuilder struct {
	context *ConfigurationContext
}

func NewConfigurationBuilder() *ConfigurationBuilder {
	return &ConfigurationBuilder{context: &ConfigurationContext{}}
}

func (builder *ConfigurationBuilder) AddFlagArgs() {
	builder.context.enableFlag = true
}

func (builder *ConfigurationBuilder) AddEnvironment() {
	builder.context.enableEnv = true
}

func (builder *ConfigurationBuilder) AddYamlFile(name string) {
	if builder.context.configType == "" {
		builder.context.configType = "yaml"
		builder.context.configName = name
	}
}

func (builder *ConfigurationBuilder) AddJsonFile(name string) {
	if builder.context.configType == "" {
		builder.context.configType = "json"
		builder.context.configName = name
	}
}

func (builder *ConfigurationBuilder) Build() *Configuration {
	return NewConfiguration(builder.context)
}
