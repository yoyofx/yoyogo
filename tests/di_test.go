package tests

import (
	"github.com/magiconair/properties/assert"
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofxteam/dependencyinjection"
	"testing"
)

func Test_DI_Register(t *testing.T) {
	he1 := &abstractions.HostEnvironment{ApplicationName: "h1"}
	he2 := &abstractions.HostEnvironment{ApplicationName: "h2"}
	services := dependencyinjection.NewServiceCollection()
	services.AddSingleton(func() *abstractions.HostEnvironment { return he1 })
	services.AddTransient(func() *abstractions.HostEnvironment { return he2 })

	serviceProvider := services.Build()

	var env *abstractions.HostEnvironment

	serviceProvider.GetService(&env)

	assert.Equal(t, env.ApplicationName, "h2")
}
