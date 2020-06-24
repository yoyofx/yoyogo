package Test

import (
	"github.com/magiconair/properties/assert"
	"github.com/yoyofx/yoyogo/DependencyInjection"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"testing"
)

func Test_DI_Register(t *testing.T) {
	he1 := &Context.HostEnvironment{ApplicationName: "h1"}
	he2 := &Context.HostEnvironment{ApplicationName: "h2"}
	services := DependencyInjection.NewServiceCollection()
	services.AddSingleton(func() *Context.HostEnvironment { return he1 })
	services.AddTransient(func() *Context.HostEnvironment { return he2 })

	serviceProvider := services.Build()

	var env *Context.HostEnvironment

	serviceProvider.GetService(&env)

	assert.Equal(t, env.ApplicationName, "h2")
}
