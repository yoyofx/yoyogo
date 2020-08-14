package YoyoGo

import (
	"fmt"
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Abstractions/Platform/ConsoleColors"
	"github.com/yoyofx/yoyogo/Abstractions/Platform/ExitHookSignals"
	"github.com/yoyofx/yoyogo/Abstractions/xlog"
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
	logger := xlog.GetXLogger("Application")
	logger.SetCustomLogFormat(logFormater)
	Abstractions.RunningHostEnvironmentSetting(hostEnv)

	Abstractions.PrintLogo(logger, hostEnv)

	ExitHookSignals.HookSignals(host)
	_ = host.webServer.Run(host.HostContext)

}

func logFormater(logInfo xlog.LogInfo) string {
	outLog := fmt.Sprintf(ConsoleColors.Yellow("[yoyogo] ")+"[%s] %s",
		logInfo.StartTime, logInfo.Message)
	return outLog
}

func (host WebHost) StopApplicationNotify() {
	host.HostContext.ApplicationCycle.StopApplication()
}

// Shutdown is Graceful stop application
func (host WebHost) Shutdown() {
	host.webServer.Shutdown()
}

func (host WebHost) SetAppMode(mode string) {
	host.HostContext.HostingEnvironment.Profile = mode
}
