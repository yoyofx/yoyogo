package YoyoGo

import (
	"encoding/base64"
	"fmt"
	"log"
	"os"
)

type HostEnv struct {
	ApplicationName string
	AppMode         string
	Args            []string
	Addr            string
	Port            string
	PID             int
}

type WebHost struct {
	hostEnv         HostEnv
	webServer       IServer
	requestDelegate IRequestDelegate
}

func NewWebHost(server IServer, request IRequestDelegate) WebHost {
	env := HostEnv{
		ApplicationName: "host",
		AppMode:         Dev,
	}
	return WebHost{webServer: server, requestDelegate: request, hostEnv: env}
}

func (host WebHost) Run() {
	l := log.New(os.Stdout, "[yoyogo] ", 0)
	host.hostEnv.Args = os.Args
	host.hostEnv.Addr = host.webServer.GetAddr()
	host.hostEnv.Port = detectAddress(host.hostEnv.Addr)
	host.hostEnv.PID = os.Getpid()

	printLogo(l, host.hostEnv)

	l.Fatal(host.webServer.Run(host.requestDelegate))
}

func (host WebHost) SetAppMode(mode string) {
	host.hostEnv.AppMode = mode
}

func printLogo(l *log.Logger, env HostEnv) {
	logo, _ := base64.StdEncoding.DecodeString("CiBfICAgICBfICAgICAgICAgICAgICAgICAgICBfX18gICAgICAgICAgCiggKSAgICggKSAgICAgICAgICAgICAgICAgICggIF9gXCAgICAgICAgCmBcYFxfLycvJ18gICAgXyAgIF8gICAgXyAgIHwgKCAoXykgICBfICAgCiAgYFwgLycvJ19gXCAoICkgKCApIC8nX2BcIHwgfF9fXyAgLydfYFwgCiAgIHwgfCggKF8pICl8IChfKSB8KCAoXykgKXwgKF8sICkoIChfKSApCiAgIChfKWBcX19fLydgXF9fLCB8YFxfX18vJyhfX19fLydgXF9fXy8nCiAgICAgICAgICAgICAoIClffCB8ICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICBgXF9fXy8nICAgICAgICAgICAgICAgICAgICAgCg==")
	fmt.Println(string(logo))

	l.Printf("listening on %s", env.Port)
	l.Printf("application is runing pid: %d", env.PID)
	l.Printf("runing in %s mode , switch on 'Prod' mode in production.", env.AppMode)
	l.Println(" - use Prod app.SetMode(Prod) ")
}
