package YoyoGo

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Abstractions/Platform/ExitHookSignals"
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
)

type WebHost struct {
	HostContext *Abstractions.HostBuildContext
	webServer   Abstractions.IServer
}

func NewWebHost(server Abstractions.IServer, hostContext *Abstractions.HostBuildContext) WebHost {
	return WebHost{webServer: server, HostContext: hostContext}
}

func (host WebHost) Run() {
	hostEnv := host.HostContext.HostingEnvironment
	logger := XLog.GetXLogger("Application")
	logger.SetCustomLogFormat(nil)
	Abstractions.RunningHostEnvironmentSetting(hostEnv)

	Abstractions.PrintLogo(logger, hostEnv)

	ExitHookSignals.HookSignals(host)
	Abstractions.HostRunning(logger, host.HostContext)
	//application running
	_ = host.webServer.Run(host.HostContext)
	//application ending
	Abstractions.HostEnding(logger, host.HostContext)

}

func (host WebHost) StopApplicationNotify() {
	logger := XLog.GetXLogger("Application")
	Abstractions.HostEnding(logger, host.HostContext)
	host.HostContext.ApplicationCycle.StopApplication()
}

// Shutdown is Graceful stop application
func (host WebHost) Shutdown() {
	host.webServer.Shutdown()
}

func (host WebHost) SetAppMode(mode string) {
	host.HostContext.HostingEnvironment.Profile = mode
}
