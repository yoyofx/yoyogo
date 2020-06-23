package YoyoGo

import (
	"encoding/base64"
	"fmt"
	"github.com/maxzhang1985/yoyogo/Abstractions"
	"github.com/maxzhang1985/yoyogo/Abstractions/Platform"
	"github.com/maxzhang1985/yoyogo/WebFramework/Context"
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
	vlog := log.New(os.Stdout, "[yoyogo] ", 0)

	Abstractions.RunningHostEnvironmentSetting(hostEnv)

	printLogo(vlog, hostEnv)

	Platform.HookSignals(host)
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

func printLogo(l *log.Logger, env *Context.HostEnvironment) {
	logo, _ := base64.StdEncoding.DecodeString("IF8gICAgIF8gICAgICAgICAgICAgICAgICAgIF9fXyAgICAgICAgICAKKCApICAgKCApICAgICAgICAgICAgICAgICAgKCAgX2BcICAgICAgICAKYFxgXF8vJy8nXyAgICBfICAgXyAgICBfICAgfCAoIChfKSAgIF8gICAKICBgXCAvJy8nX2BcICggKSAoICkgLydfYFwgfCB8X19fICAvJ19gXCAKICAgfCB8KCAoXykgKXwgKF8pIHwoIChfKSApfCAoXywgKSggKF8pICkKICAgKF8pYFxfX18vJ2BcX18sIHxgXF9fXy8nKF9fX18vJ2BcX19fLycKICAgICAgICAgICAgICggKV98IHwgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgIGBcX19fLycgICAgICAgICAgICBMaWdodCBhbmQgZmFzdC4gIA==")
	fmt.Println(string(logo))

	l.Printf("version: %s", env.Version)
	l.Printf("listening on %s", env.Port)
	l.Printf("application is running , pid: %d", env.PID)
	l.Printf("running in %s mode , switch on 'Dev' or 'Test' or 'Prod' mode in production.", env.Profile)
	l.Print("- use Prod by app.SetEnvironment(Context.Prod) ")
	l.Print("Starting HTTP server...")
}
