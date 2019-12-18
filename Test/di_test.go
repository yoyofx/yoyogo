package Test

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/DependencyInjection"
	"testing"
)

func Test_DI_Register(t *testing.T) {
	he1 := &Context.HostEnvironment{ApplicationName: "h1"}
	he2 := &Context.HostEnvironment{ApplicationName: "h2"}
	services := DependencyInjection.NewServiceCollection()
	services.AddSingleton(func() *Context.HostEnvironment { return he1 })
	services.AddSingleton(func() *Context.HostEnvironment { return he2 })

	serviceProvier := services.Build()

	_ = serviceProvier
}
