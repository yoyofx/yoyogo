package YoyoGo

import (
	"github.com/maxzhang1985/yoyogo/Abstractions"
	"github.com/maxzhang1985/yoyogo/WebFramework/Context"
)

type WebHostBuilder struct {
	Abstractions.HostBuilder
}

func NewWebHostBuilder() *WebHostBuilder {
	return &WebHostBuilder{
		Abstractions.HostBuilder{
			Context:   &Abstractions.HostBuildContext{HostingEnvironment: &Context.HostEnvironment{}},
			Decorator: NewWebHostBuilderDecorator(),
		},
	}
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
