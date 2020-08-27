package Test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/Examples/SimpleWeb/contollers"
	"github.com/yoyofx/yoyogo/Examples/SimpleWeb/models"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofxteam/reflectx"
	"reflect"
	"strings"
	"testing"
)

//
//import (
//	"fmt"
//	"github.com/stretchr/testify/assert"
//	"github.com/yoyofx/yoyogo/Examples/SimpleWeb/contollers"
//	"github.com/yoyofx/yoyogo/Utils"
//	"github.com/yoyofx/yoyogo/Utils/Reflect"
//	"github.com/yoyofx/yoyogo/WebFramework/Context"
//	"github.com/yoyofx/yoyogo/WebFramework/Mvc"
//	"github.com/yoyofxteam/reflectx"
//	"testing"
//)
//
//type Person struct {
//	Name    string
//	Student *Student
//}
//
//type Student struct {
//	Name  string `json:"name"`
//	Age   int    `json:"age"`
//	Grade int    `json:"grade"`
//}
//
//func (typeInfo Student) Hello() string {
//	return "hello"
//}
//
//func (typeInfo Student) Say(hi string) string {
//	return "Hello " + hi
//}
//
//func Test_MethodCallerCall2(t *testing.T) {
//	utype := &UserInfo{}
//
//	methodInfo, _ := reflectx.GetObjectMethodInfoByName(utype, "Hello")
//	results := methodInfo.Invoke(&Context.HttpContext{}, "hello world!")
//
//	fmt.Println()
//	fmt.Printf("Result: %s", results)
//	fmt.Println()
//
//	assert.Equal(t, results[0].(string), "hello world!")
//}
//
//func Test_RecCreateStruct(t *testing.T) {
//	//yourtype := reflect.TypeOf(Mvc.RequestBody{})
//	//dd := Reflect.CreateInstance(yourtype)
//	//_ = dd
//	typeInfo, _ := Reflect.GetTypeInfo(Mvc.RequestBody{})
//	ins := typeInfo.CreateInstance()
//	assert.Equal(t, ins != nil, true)
//}
//
//func Test_GetCtorFuncTypeName(t *testing.T) {
//	ctorFunc := contollers.NewUserController
//	name, _ := Reflect.GetCtorFuncOutTypeName(ctorFunc)
//	name = Utils.LowercaseFirst(name)
//	assert.Equal(t, name, "userController")
//}
//
//func Test_ReflectStructFields(t *testing.T) {
//	student := &Student{
//		Name:  "json",
//		Age:   18,
//		Grade: 9,
//	}
//	p := Person{
//		Name:    "Json",
//		Student: student,
//	}
//
//	ptype, _ := Reflect.GetTypeInfo(p)
//	pf1 := ptype.GetFieldByName("Name")
//	assert.Equal(t, pf1.GetValue(), "Json")
//	pf2 := ptype.GetFieldByName("Student")
//	assert.Equal(t, pf2.GetValue(), student)
//	typeInfo, _ := pf2.AsTypeInfo()
//
//	typeInfo.GetFieldByName("Grade").SetValue(11)
//	assert.Equal(t, student.Grade, 11)
//	assert.Equal(t, typeInfo.HasMethods(), true)
//	sayRet := typeInfo.GetMethodByName("Say").Invoke("World!")[0].(string)
//	assert.Equal(t, sayRet, "Hello World!")
//
//}

type Person struct {
	Name    string
	Student *Student
}

type Student struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Grade int    `json:"grade"`
}

func (typeInfo Student) Hello() string {
	fmt.Println("hello ")
	return "hello"
}

func (typeInfo Student) Say(hi string) string {
	fmt.Println("Hello " + hi)
	return "Hello " + hi
}

func Test_GetStructMethodList(t *testing.T) {
	userInfo := &UserInfo{}
	userMethodList := reflectx.GetObjectMethodInfoList(userInfo)
	assert.Equal(t, len(userMethodList), 2)
	assert.Equal(t, getMehtodInfoByName(userMethodList, "Hello").Invoke(userInfo, &Context.HttpContext{}, "UserInfo Func Call:Hello,")[0],
		"UserInfo Func Call:Hello,")
	//---------------------------------------------------------------------------------------------
	student := Student{}
	studentMethodList := reflectx.GetObjectMethodInfoList(student)
	assert.Equal(t, len(studentMethodList), 2)
	assert.Equal(t, studentMethodList[0].Name, "Hello")
	assert.Equal(t, studentMethodList[1].Name, "Say")

	assert.Equal(t, getMehtodInfoByName(studentMethodList, "Hello").Invoke(student)[0], "hello")
	assert.Equal(t, getMehtodInfoByName(studentMethodList, "Say").Invoke(student, "Say: Student")[0], "Hello Say: Student")
}

func getMehtodInfoByName(infos []reflectx.MethodInfo, name string) reflectx.MethodInfo {
	for _, m := range infos {
		if m.Name == name {
			return m
		}
	}
	return infos[0]
}

func Test_UserController(t *testing.T) {
	controllerCtor := contollers.NewUserController
	controllerName, controllerType := reflectx.GetCtorFuncOutTypeName(controllerCtor)
	controllerName = strings.ToLower(controllerName)
	// Create Controller and Action descriptors

	instance := reflect.New(controllerType).Interface()
	actionList := reflectx.GetObjectMethodInfoList(instance)
	mi := getMehtodInfoByName(actionList, "GetUserName")
	_ = mi.Parameters[0].ParameterType.Elem().Name()
	rets := mi.Invoke(instance, &Context.HttpContext{}, &contollers.RegisterRequest{
		UserName: "he",
		Password: "123",
	})
	assert.Equal(t, len(rets), 1)

	instance1 := contollers.NewUserController(models.NewUserAction(), nil)
	_ = instance1
	//actionList = reflectx.GetObjectMethodInfoList(instance)
	//_ = actionList
	//mi = getMehtodInfoByName(actionList,"GetUserName")
	//mi.Invoke(instance,&Context.HttpContext{}, &contollers.RegisterRequest{
	//	 UserName: "he",
	//	 Password: "123",
	//})
}
