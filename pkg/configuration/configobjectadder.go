package configuration

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofxteam/dependencyinjection"
)

// AddConfiguration 注入函数 用户API
func AddConfiguration(sc *dependencyinjection.ServiceCollection, objType interface{}) {
	//_, objectType := reflectx.GetCtorFuncOutTypeName(objType)
	//configObject := reflect.New(objectType).Interface().(abstractions.IConfigurationProperties)
	//sectionName := configObject.GetSection()
	//fmt.Println(sectionName)
	sc.AddTransient(objType)
}

func Local(configName string) *abstractions.Configuration {
	config := abstractions.NewConfigurationBuilder().
		AddEnvironment().
		AddYamlFile(configName).Build()
	return config
}
