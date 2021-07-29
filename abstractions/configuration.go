package abstractions

import "github.com/spf13/viper"

type IConfiguration interface {
	Get(name string) interface{}
	GetString(name string) string
	GetBool(name string) bool
	GetInt(name string) int
	GetSection(name string) IConfiguration
	Unmarshal(interface{})
	GetProfile() string
	GetConfDir() string
	GetConfigObject(configTag string, configObject interface{})
	RefreshAll()
	RefreshBy(name string)
}

type IConfigurationRemoteProvider interface {
	GetProvider(*viper.Viper) *viper.Viper
}

type ConfigurationProperties struct {
}
