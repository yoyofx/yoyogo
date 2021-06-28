package apollo

import (
	"fmt"
	"github.com/spf13/viper"
	remote "github.com/yoyofxteam/agollo/viper-remote"
)

type ViperRemoteProvider struct {
	configType string
	configSet  string
}

func NewRemoteProvider(configType string) *ViperRemoteProvider {
	return &ViperRemoteProvider{
		configType: configType,
		configSet:  "yoyogo.cloud.configcenter.apollo"}
}

func (provider *ViperRemoteProvider) GetProvider(runtime_viper *viper.Viper) *viper.Viper {
	var option *Option
	err := runtime_viper.Sub(provider.configSet).Unmarshal(&option)
	if err != nil {
		panic(err)
		return nil
	}
	remote.SetAppID(option.AppID)

	remote_viper := viper.New()
	err = remote_viper.AddRemoteProvider("apollo", option.Endpoint, option.Namespace)
	if provider.configType == "" {
		provider.configType = "yaml"
	}
	remote_viper.SetConfigType(provider.configType)
	err = remote_viper.ReadRemoteConfig()

	if err == nil && len(remote_viper.AllSettings()) > 0 {
		err = remote_viper.WatchRemoteConfigOnChannel()
		if err == nil {
			fmt.Println("config center ..........")
			fmt.Println("used remote viper by apollo")
			fmt.Printf("apollo config: endpoint %s , namespace: %s , app_id: %s", option.Endpoint, option.Namespace, option.AppID)
			return remote_viper
		}
	}
	return runtime_viper
}
