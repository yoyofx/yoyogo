package apollo

import (
	"bytes"
	"fmt"
	remote "github.com/shima-park/agollo/viper-remote"
	"github.com/spf13/viper"
)

type ViperRemoteProvider struct {
	configType    string
	configSet     string
	viperProvider *RemoteProvider
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
	provider.viperProvider = DefaultRemoteProvider()
	provider.viperProvider.provider = "apollo"
	provider.viperProvider.endpoint = option.Endpoint
	provider.viperProvider.path = option.Namespace
	err = remote_viper.AddRemoteProvider("apollo", option.Endpoint, option.Namespace)
	if provider.configType == "" {
		provider.configType = "yaml"
	}
	remote_viper.SetConfigType(provider.configType)
	err = remote_viper.ReadRemoteConfig()

	if err == nil && len(remote_viper.AllSettings()) > 0 {
		//err = remote_viper.WatchRemoteConfigOnChannel()
		if err == nil {
			fmt.Println("config center ..........")
			fmt.Println("used remote viper by apollo")
			fmt.Printf("apollo config: endpoint %s , namespace: %s , app_id: %s", option.Endpoint, option.Namespace, option.AppID)
			return remote_viper
		}
	}
	return runtime_viper
}

func (provider *ViperRemoteProvider) WatchRemoteConfigOnChannel(remoteViper *viper.Viper) <-chan bool {
	updater := make(chan bool)

	respChan, _ := viper.RemoteConfig.WatchChannel(provider.viperProvider)
	go func(rc <-chan *viper.RemoteResponse) {
		for {
			b := <-rc
			reader := bytes.NewReader(b.Value)
			_ = remoteViper.ReadConfig(reader)
			// configuration on changed
			updater <- true
		}
	}(respChan)

	return updater
}
