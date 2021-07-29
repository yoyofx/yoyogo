package models

import "github.com/yoyofx/yoyogo/abstractions"

const DbConfigTag = "yoyogo.datasource.db"

type DbConfig struct {
	Name     string `mapstructure:"name" config:"name"`
	Url      string `mapstructure:"url" config:"url"`
	UserName string `mapstructure:"username" config:"user_name"`
	Password string `mapstructure:"password" config:"password"`
	Debug    bool   `mapstructure:"debug" config:"debug"`
}

func (db *DbConfig) GetSection() string {
	return DbConfigTag
}

func NewDbConfig(configuration abstractions.IConfiguration) DbConfig {
	var config DbConfig
	configuration.GetConfigObject(DbConfigTag, &config) //configuration.GetConfigObject(configuration,DbConfigTag,config)
	return config
}
