package Test

import (
	"fmt"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
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

func Test_StructGetFieldTag(t *testing.T) {
	user := &UserInfo{"John Doe The Fourth", 20}

	value := reflect.TypeOf(user).Elem()
	for i := 0; i < value.NumField(); i++ {
		f := value.Field(i)
		fmt.Printf("%d: %s %s %s \n", i,
			f.Name, f.Type, f.Tag.Get("json"))
	}
}
