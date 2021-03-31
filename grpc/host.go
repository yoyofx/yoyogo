package grpc

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/platform/exithooksignals"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
)

type Host struct {
	HostContext *abstractions.HostBuilderContext
	Server      abstractions.IServer
	logger      xlog.ILogger
}

func NewHost(server abstractions.IServer, hostContext *abstractions.HostBuilderContext) Host {
	return Host{Server: server, HostContext: hostContext, logger: xlog.GetXLogger("Grpc Application")}
}

func (host Host) Run() {
	hostEnv := host.HostContext.HostingEnvironment

	host.logger.SetCustomLogFormat(nil)
	abstractions.RunningHostEnvironmentSetting(hostEnv)

	abstractions.PrintLogo(host.logger, hostEnv)

	exithooksignals.HookSignals(host)
	abstractions.HostRunning(host.logger, host.HostContext)
	//application running
	_ = host.Server.Run(host.HostContext)
	//application ending
	abstractions.HostEnding(host.logger, host.HostContext)

}

func (host Host) StopApplicationNotify() {
	abstractions.HostEnding(host.logger, host.HostContext)
	host.HostContext.ApplicationCycle.StopApplication()
}

// Shutdown is Graceful stop application
func (host Host) Shutdown() {
	host.Server.Shutdown()
}

func (host Host) SetAppMode(mode string) {
	host.HostContext.HostingEnvironment.Profile = mode
}
