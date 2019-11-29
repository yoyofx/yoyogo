package YoyoGo

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

type WebHost struct {
	hostEnv         *HostEnv
	webServer       IServer
	requestDelegate IRequestDelegate
}

func NewWebHost(server IServer, request IRequestDelegate, hostContext *HostBuildContext) WebHost {

	return WebHost{webServer: server, requestDelegate: request, hostEnv: hostContext.hostingEnvironment}
}

func (host WebHost) Run() {
	vlog := log.New(os.Stdout, "[yoyogo] ", 0)
	host.hostEnv.Args = os.Args
	host.hostEnv.Addr = host.webServer.GetAddr()
	host.hostEnv.Port = detectAddress(host.hostEnv.Addr)
	host.hostEnv.PID = os.Getpid()
	host.hostEnv.Version = Version
	printLogo(vlog, host.hostEnv)

	vlog.Fatal(host.webServer.Run(host.requestDelegate))

}

func (host WebHost) SetAppMode(mode string) {
	host.hostEnv.AppMode = mode
}

func printLogo(l *log.Logger, env *HostEnv) {
	logo, _ := base64.StdEncoding.DecodeString("IF8gICAgIF8gICAgICAgICAgICAgICAgICAgIF9fXyAgICAgICAgICAKKCApICAgKCApICAgICAgICAgICAgICAgICAgKCAgX2BcICAgICAgICAKYFxgXF8vJy8nXyAgICBfICAgXyAgICBfICAgfCAoIChfKSAgIF8gICAKICBgXCAvJy8nX2BcICggKSAoICkgLydfYFwgfCB8X19fICAvJ19gXCAKICAgfCB8KCAoXykgKXwgKF8pIHwoIChfKSApfCAoXywgKSggKF8pICkKICAgKF8pYFxfX18vJ2BcX18sIHxgXF9fXy8nKF9fX18vJ2BcX19fLycKICAgICAgICAgICAgICggKV98IHwgICAgICAgICAgICAgICAgICAgICAKICAgICAgICAgICAgIGBcX19fLycgICAgICAgICAgICBMaWdodCBhbmQgZmFzdC4gIA==")
	fmt.Println(string(logo))

	l.Printf("version: %s", env.Version)
	l.Printf("listening on %s", env.Port)
	l.Printf("application is runing , pid: %d", env.PID)
	l.Printf("runing in %s mode , switch on 'Prod' mode in production.", env.AppMode)
	l.Print(" - use Prod app.SetMode(Prod) ")
	l.Print("Starting HTTP server...")
}
