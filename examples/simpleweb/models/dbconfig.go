package models

import "github.com/yoyofx/yoyogo/abstractions"

type DbConfig struct {
	Name     string `mapstructure:"name" config:"name"`
	Url      string `mapstructure:"url" config:"url"`
	UserName string `mapstructure:"username" config:"user_name"`
	Password string `mapstructure:"password" config:"password"`
	Debug    bool   `mapstructure:"debug" config:"debug"`
}

// NewDbConfig 托管到IConfiguration和依赖注入中管理
func NewDbConfig(configuration abstractions.IConfiguration) (config DbConfig) {
	configuration.GetConfigObject("yoyogo.datasource.db", &config)
	return config
}
