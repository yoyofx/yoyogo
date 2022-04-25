package etcd

import "github.com/yoyofx/yoyogo/abstractions"

type Config struct {
	ENV       *abstractions.HostEnvironment
	Address   []string `mapstructure:"address" config:"address"`
	Namespace string   `mapstructure:"namespace" config:"namespace"`
	Ttl       int64    `mapstructure:"ttl" config:"ttl"`
	Auth      *Auth    `mapstructure:"auth" config:"auth"`
}

type Auth struct {
	Enable   bool   `mapstructure:"enable" config:"enable"`
	User     string `mapstructure:"username" config:"username"`
	Password string `mapstructure:"password" config:"password"`
}
