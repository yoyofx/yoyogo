package configuration

import (
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofxteam/dependencyinjection"
	"github.com/yoyofxteam/reflectx"
	"reflect"
)

// AddConfiguration 注入函数 用户API
func AddConfiguration(sc *dependencyinjection.ServiceCollection, objType interface{}) {
	_, objectType := reflectx.GetCtorFuncOutTypeName(objType)
	configObject := reflect.New(objectType).Interface().(abstractions.IConfigurationProperties)
	sectionName := configObject.GetSection()
	fmt.Println(sectionName)
	sc.AddTransient(objType)
}
