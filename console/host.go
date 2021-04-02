package console

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/platform/exithooksignals"
)

type Host struct {
	abstractions.ServiceHost
}

func NewHost(server abstractions.IServer, hostContext *abstractions.HostBuilderContext) Host {
	hostContext.HostingEnvironment.Server = "console"
	hostContext.HostingEnvironment.Port = "0"
	return Host{abstractions.NewServiceHost(server, hostContext)}
}

func (host Host) Run() {
	exithooksignals.HookSignals(host)
	host.ServiceHost.Run()
}
