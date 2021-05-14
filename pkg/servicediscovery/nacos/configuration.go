package nacos

import (
	"github.com/yoyofx/yoyogo/abstractions"
)

const (
	GroupName = "DEFAULT_GROUP"
	Cluster   = "DEFAULT"
)

type Config struct {
	ENV         *abstractions.HostEnvironment
	Url         string `mapstructure:"url"`
	Port        uint64 `mapstructure:"port"`
	NamespaceId string `mapstructure:"namespace"`
	GroupName   string `mapstructure:"group"`
	Cluster     string `mapstructure:"cluster"`
	Auth        *Auth  `mapstructure:"auth"`
}

type Auth struct {
	Enable   bool   `mapstructure:"enable"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
