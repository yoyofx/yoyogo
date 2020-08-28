package Nacos

import "github.com/yoyofx/yoyogo/WebFramework/Context"

const (
	GroupName = "DEFAULT_GROUP"
	Cluster   = "DEFAULT"
)

type Config struct {
	ENV         *Context.HostEnvironment
	Url         string `mapstructure:"url"`
	Port        uint64 `mapstructure:"port"`
	NamespaceId string `mapstructure:"namespace"`
	GroupName   string `mapstructure:"group_name"`
	ClusterName string `mapstructure:"cluster_name"`
}
