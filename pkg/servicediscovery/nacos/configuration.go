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
	GroupName   string `mapstructure:"group_name"`
	ClusterName string `mapstructure:"cluster_name"`
}
