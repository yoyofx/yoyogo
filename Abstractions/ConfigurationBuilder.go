package Abstractions

type ConfigurationContext struct {
	enableFlag bool
	enableEnv  bool
	configType string
	configName string
	profile    string
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

func (builder *ConfigurationBuilder) Build(env string) *Configuration {
	builder.context.profile = env
	return NewConfiguration(builder.context)
}
