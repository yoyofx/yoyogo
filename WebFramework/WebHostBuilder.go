package YoyoGo

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
)

type WebHostBuilder struct {
	Abstractions.HostBuilder
}

func NewWebHostBuilder() *WebHostBuilder {
	builder := &WebHostBuilder{
		Abstractions.HostBuilder{
			Context:   &Abstractions.HostBuildContext{HostingEnvironment: &Context.HostEnvironment{}},
			Decorator: NewWebHostBuilderDecorator(),
		},
	}

	return builder
}

// SetEnvironment set value(Dev,Test,Prod) by environment
func (host *WebHostBuilder) SetEnvironment(mode string) *WebHostBuilder {
	host.HostBuilder.SetEnvironment(mode)
	return host
}

func (host *WebHostBuilder) UseFastHttpByAddr(addr string) *WebHostBuilder {
	host.Server = NewFastHttp(addr)
	return host
}

func (host *WebHostBuilder) UseFastHttp() *WebHostBuilder {
	host.Server = NewFastHttp("")
	return host
}

func (host *WebHostBuilder) UseHttpByAddr(addr string) *WebHostBuilder {
	host.Server = DefaultHttpServer(addr)
	return host
}

func (host *WebHostBuilder) UseHttp() *WebHostBuilder {
	host.Server = DefaultHttpServer("")
	return host
}
