package YoyoGo

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/Abstractions/Platform/ConsoleColors"
	"github.com/yoyofx/yoyogo/Abstractions/Platform/ExitHookSignals"
	"log"
	"os"
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
	xlog := log.New(os.Stdout, ConsoleColors.Yellow("[yoyogo] "), 1)

	Abstractions.RunningHostEnvironmentSetting(hostEnv)

	Abstractions.PrintLogo(xlog, hostEnv)

	ExitHookSignals.HookSignals(host)
	_ = host.webServer.Run(host.HostContext)

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
