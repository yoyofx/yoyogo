package tests

import (
	"fmt"
	"github.com/magiconair/properties/assert"
	"github.com/yoyofx/yoyogo/web/mvc"
	"testing"
)

type TestActionFilter struct {
}

func (f *TestActionFilter) OnActionExecuting(context mvc.ActionFilterContext) bool {
	fmt.Println("TestActionFilter OnActionExecuted")
	return true
}

func (f *TestActionFilter) OnActionExecuted(context mvc.ActionFilterContext) {
	fmt.Println("TestActionFilter OnActionExecuted")
}

func Test_Filter(t *testing.T) {

	chain := mvc.NewActionFilterChain("u*/get*", &TestActionFilter{})

	assert.Equal(t, chain.MatchPath("user/getuser"), true)

	assert.Equal(t, chain.MatchPath("user/get/1"), true)

	assert.Equal(t, chain.MatchPath("/user/get/1"), false)

	assert.Equal(t, chain.MatchPath("v1/user/get/1"), false)

	filter := chain.MatchFilter("user/get/1")
	assert.Equal(t, filter != nil, true)
	c := mvc.ActionFilterContext{}
	assert.Equal(t, filter.OnActionExecuting(c), true)
	filter.OnActionExecuted(c)

	chain1 := mvc.NewActionFilterChain("v1/user/info", &TestActionFilter{})
	assert.Equal(t, chain1.MatchPath("v1/user/info"), true)

}
