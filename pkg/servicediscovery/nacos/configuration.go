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
	Url         string `mapstructure:"url" config:"url"`
	Port        uint64 `mapstructure:"port" config:"port"`
	NamespaceId string `mapstructure:"namespace" config:"namespace"`
	GroupName   string `mapstructure:"group" config:"group"`
	Cluster     string `mapstructure:"cluster" config:"cluster"`
	Auth        *Auth  `mapstructure:"auth" config:"auth"`
}

type Auth struct {
	Enable   bool   `mapstructure:"enable" config:"enable"`
	User     string `mapstructure:"username" config:"username"`
	Password string `mapstructure:"password" config:"password"`
	// ACM Endpoint
	Endpoint string `mapstructure:"endpoint" config:"endpoint"`
	// ACM RegionId
	RegionId string `mapstructure:"regionId" config:"regionId"`
	// ACM AccessKey
	AccessKey string `mapstructure:"accessKey" config:"accessKey"`
	// ACM SecretKey
	SecretKey string `mapstructure:"secretKey" config:"secretKey"`
	// ACM OpenKMS
	OpenKMS bool `mapstructure:"openKMS" config:"openKMS"`
}
