package servicediscovery

// Config 服务发现配置
type Config struct {
	RegisterWithSelf bool
}

func NewConfig(regWithSelf bool) *Config {
	return &Config{RegisterWithSelf: regWithSelf}
}
