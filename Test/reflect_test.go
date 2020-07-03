package Test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/Examples/SimpleWeb/contollers"
	"github.com/yoyofx/yoyogo/Utils"
	"github.com/yoyofx/yoyogo/Utils/Reflect"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Mvc"
	"testing"
)

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

func Test_MethodCallerCall2(t *testing.T) {
	utype := &UserInfo{}

	methodInfo := Reflect.GetObjectMehtodInfoByName(utype, "Hello")
	results := methodInfo.Invoke(&Context.HttpContext{}, "hello world!")

	fmt.Println()
	fmt.Printf("Result: %s", results)
	fmt.Println()

	assert.Equal(t, results[0].(string), "hello world!")
}

func Test_RecCreateStruct(t *testing.T) {
	//yourtype := reflect.TypeOf(Mvc.RequestBody{})
	//dd := Reflect.CreateInstance(yourtype)
	//_ = dd
	typeInfo, _ := Reflect.GetTypeInfo(Mvc.RequestBody{})
	ins := typeInfo.CreateInstance()
	assert.Equal(t, ins != nil, true)
}

func Test_GetCtorFuncTypeName(t *testing.T) {
	ctorFunc := contollers.NewUserController
	name, _ := Reflect.GetCtorFuncOutTypeName(ctorFunc)
	name = Utils.LowercaseFirst(name)
	assert.Equal(t, name, "userController")
}

func Test_ReflectStructFields(t *testing.T) {
	//student := Student{
	//	Name:  "json",
	//	Age:   18,
	//	Grade: 9,
	//}

	//typeInfo,_ := Reflect.GetTypeInfo(student)

}
