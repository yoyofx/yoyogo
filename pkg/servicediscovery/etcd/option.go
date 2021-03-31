package etcd

import "github.com/yoyofx/yoyogo/abstractions"

type Config struct {
	ENV       *abstractions.HostEnvironment
	Address   []string `mapstructure:"address"`
	Namespace string   `mapstructure:"namespace"`
	Ttl       int64    `mapstructure:"ttl"`
	Auth      *Auth    `mapstructure:"auth"`
}

type Auth struct {
	Enable   bool   `mapstructure:"enable"`
	User     string `mapstructure:"username"`
	Password string `mapstructure:"password"`
}
