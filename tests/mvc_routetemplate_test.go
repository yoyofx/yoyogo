package tests

import (
	"github.com/magiconair/properties/assert"
	"github.com/yoyofx/yoyogo/web/mvc"
	"strings"
	"testing"
)

func Test_MvcTemplate(t *testing.T) {
	url := "v1/usercontroller/register"
	template := mvc.NewRouteTemplate("v1/{controller}/{action}")
	assert.Equal(t, template.GetControllerIndex(), 1)
	assert.Equal(t, template.GetActionIndex(), 2)

	matchInfo := &mvc.MatchMvcInfo{}
	assert.Equal(t, template.Match(strings.Split(url, "/"), matchInfo), true)
	assert.Equal(t, matchInfo.ControllerName, "usercontroller")
	assert.Equal(t, matchInfo.ActionName, "register")

	template1 := mvc.NewRouteTemplate("api/v1/{controller}/{action}")
	assert.Equal(t, template1.GetControllerIndex(), 2)
	assert.Equal(t, template1.GetActionIndex(), 3)

	assert.Equal(t, template1.Match(strings.Split(url, "/"), matchInfo), false)
	assert.Equal(t, template1.Match(strings.Split("api/"+url, "/"), matchInfo), true)
	assert.Equal(t, matchInfo.ControllerName, "usercontroller")
	assert.Equal(t, matchInfo.ActionName, "register")

}
