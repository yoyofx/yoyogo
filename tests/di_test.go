package tests

import (
	"github.com/magiconair/properties/assert"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"github.com/yoyofx/yoyogo/web/context"
	"testing"
)

func Test_DI_Register(t *testing.T) {
	he1 := &context.HostEnvironment{ApplicationName: "h1"}
	he2 := &context.HostEnvironment{ApplicationName: "h2"}
	services := dependencyinjection.NewServiceCollection()
	services.AddSingleton(func() *context.HostEnvironment { return he1 })
	services.AddTransient(func() *context.HostEnvironment { return he2 })

	serviceProvider := services.Build()

	var env *context.HostEnvironment

	serviceProvider.GetService(&env)

	assert.Equal(t, env.ApplicationName, "h2")
}
