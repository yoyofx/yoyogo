package Test

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/yoyofx/yoyogo/Utils/Reflect"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"testing"
)

func Test_MethodCallerCall2(t *testing.T) {
	utype := &UserInfo{}

	methodInfo := Reflect.GetObjectMehtodInfoByName(utype, "Hello")
	results := methodInfo.Invoke(&Context.HttpContext{}, "hello world!")

	fmt.Println()
	fmt.Printf("Result: %s", results)
	fmt.Println()

	assert.Equal(t, results[0].(string), "hello world!")
}
