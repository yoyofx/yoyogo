package Test

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/maxzhang1985/yoyogo/Examples/SimpleWeb/contollers"
	"github.com/maxzhang1985/yoyogo/Utils"
	"github.com/maxzhang1985/yoyogo/WebFramework/Context"
	_ "github.com/maxzhang1985/yoyogo/WebFramework/Context"
	"github.com/maxzhang1985/yoyogo/WebFramework/Mvc"
	"reflect"
	"testing"
)

type UserInfo struct {
	Name string `json:"name" w1:"12"`
	Age  int
}

func (user *UserInfo) Say(hi string) {
	fmt.Print(hi)
}

func (user *UserInfo) Hello(context *Context.HttpContext, hi string) string {
	return hi
}

func Test_reflectCall(t *testing.T) {
	utype := new(UserInfo)
	result := reflectCall(utype, "Hello", &Context.HttpContext{}, "hello world!")

	fmt.Println()
	fmt.Printf("Result: %s", result)
	fmt.Println()
}

func reflectCall(ctype interface{}, funcName string, params ...interface{}) interface{} {
	t := reflect.ValueOf(ctype)
	methodInfo := t.MethodByName(funcName)
	methodType := methodInfo.Type()

	methodParamsNum := methodType.NumIn()
	paramTypes := make([]reflect.Type, methodParamsNum)
	paramValues := make([]reflect.Value, methodParamsNum)
	for idx := 0; idx < methodParamsNum; idx++ {
		paramTypes[idx] = methodType.In(idx)
		paramValues[idx] = reflect.ValueOf(params[idx])
	}

	fmt.Printf("Type: %s ,Call Method: %s", t.Type().Name(), funcName)
	fmt.Printf("%s", paramTypes)

	rets := t.MethodByName(funcName).Call(paramValues)

	if len(rets) > 0 {
		return rets[0].Interface()
	}
	return nil
}

func Test_MethodCallerCall(t *testing.T) {
	utype := &UserInfo{}
	method := Utils.NewMethodCaller(utype, "Hello")
	results := method.Invoke(&Context.HttpContext{}, "hello world!")

	fmt.Println()
	fmt.Printf("Result: %s", results)
	fmt.Println()
}

func Test_StructGetFieldTag(t *testing.T) {
	user := &UserInfo{"John Doe The Fourth", 20}

	value := reflect.TypeOf(user).Elem()
	for i := 0; i < value.NumField(); i++ {
		f := value.Field(i)
		fmt.Printf("%d: %s %s %s \n", i,
			f.Name, f.Type, f.Tag.Get("json"))
	}
}

func Test_RecCreateStruct(t *testing.T) {
	yourtype := reflect.TypeOf(Mvc.RequestBody{})
	dd := reflect.New(yourtype).Elem().Interface()
	_ = dd
}

func Test_GetCtorFuncTypeName(t *testing.T) {
	ctorFunc := contollers.NewUserController
	name := Utils.GetCtorFuncName(ctorFunc)
	name = Utils.LowercaseFirst(name)
	assert.Equal(t, name, "userController")
}
