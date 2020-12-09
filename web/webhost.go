package web

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"github.com/yoyofx/yoyogo/abstractions/platform/exithooksignals"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
)

type WebHost struct {
	HostContext *abstractions.HostBuilderContext
	webServer   abstractions.IServer
}

func NewWebHost(server abstractions.IServer, hostContext *abstractions.HostBuilderContext) WebHost {
	return WebHost{webServer: server, HostContext: hostContext}
}

func (host WebHost) Run() {
	hostEnv := host.HostContext.HostingEnvironment
	logger := xlog.GetXLogger("Application")
	logger.SetCustomLogFormat(nil)
	abstractions.RunningHostEnvironmentSetting(hostEnv)

	abstractions.PrintLogo(logger, hostEnv)

	exithooksignals.HookSignals(host)
	abstractions.HostRunning(logger, host.HostContext)
	//application running
	_ = host.webServer.Run(host.HostContext)
	//application ending
	abstractions.HostEnding(logger, host.HostContext)

}

func (host WebHost) StopApplicationNotify() {
	logger := xlog.GetXLogger("Application")
	abstractions.HostEnding(logger, host.HostContext)
	host.HostContext.ApplicationCycle.StopApplication()
}

// Shutdown is Graceful stop application
func (host WebHost) Shutdown() {
	host.webServer.Shutdown()
}

func (host WebHost) SetAppMode(mode string) {
	host.HostContext.HostingEnvironment.Profile = mode
}
