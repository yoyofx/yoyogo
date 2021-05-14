package web

import (
	"github.com/yoyofx/yoyogo/abstractions"
)

type WebHostBuilder struct {
	abstractions.HostBuilder
}

func NewWebHostBuilder() *WebHostBuilder {
	builder := &WebHostBuilder{
		abstractions.HostBuilder{
			Context:   &abstractions.HostBuilderContext{HostingEnvironment: &abstractions.HostEnvironment{}},
			Decorator: NewWebHostBuilderDecorator(),
		},
	}

	return builder
}

//
//// SetEnvironment set value(Dev,tests,Prod) by environment
//func (host *WebHostBuilder) SetEnvironment(mode string) *WebHostBuilder {
//	host.HostBuilder.SetEnvironment(mode)
//	return host
//}
//
//func (host *WebHostBuilder) UseFastHttpByAddr(addr string) *WebHostBuilder {
//	host.Server = NewFastHttp(addr)
//	return host
//}
//
//func (host *WebHostBuilder) UseFastHttp() *WebHostBuilder {
//	host.Server = NewFastHttp("")
//	return host
//}
//
//func (host *WebHostBuilder) UseHttpByAddr(addr string) *WebHostBuilder {
//	host.Server = DefaultHttpServer(addr)
//	return host
//}
//
//func (host *WebHostBuilder) UseHttp() *WebHostBuilder {
//	host.Server = DefaultHttpServer("")
//	return host
//}
