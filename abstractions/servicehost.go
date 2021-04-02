package abstractions

import (
	"encoding/base64"
	"fmt"
	"github.com/yoyofx/yoyogo"
	"github.com/yoyofx/yoyogo/abstractions/platform/consolecolors"
	"github.com/yoyofx/yoyogo/abstractions/servicediscovery"
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/utils"
	"strconv"
)

type IServiceHost interface {
	Run()
	Shutdown()
	StopApplicationNotify()
	SetAppMode(mode string)
}

func HostRunning(log xlog.ILogger, context *HostBuilderContext) {
	go startServerDiscovery(log, context)
}

func HostEnding(log xlog.ILogger, context *HostBuilderContext) {
	endServerDiscovery(log, context)
}

func startServerDiscovery(log xlog.ILogger, context *HostBuilderContext) {
	var sd servicediscovery.IServiceDiscovery
	_ = context.HostServices.GetService(&sd)
	if sd != nil {
		_ = sd.Register()
	}
}

func endServerDiscovery(log xlog.ILogger, context *HostBuilderContext) {
	var sd servicediscovery.IServiceDiscovery
	var sdcache servicediscovery.Cache
	err := context.HostServices.GetService(&sdcache)
	if err == nil {
		sdcache.Stop()
	}
	err = context.HostServices.GetService(&sd)
	if err == nil && sd != nil {
		_ = sd.Destroy()
	}
}

func PrintLogo(l xlog.ILogger, env *HostEnvironment) {
	logo, _ := base64.StdEncoding.DecodeString(yoyogo.Logo)

	fmt.Println(consolecolors.Blue(string(logo)))
	fmt.Println(" ")
	fmt.Printf("%s   (version:  %s)", consolecolors.Green(":: YoyoGo ::"), consolecolors.Blue(env.Version))

	fmt.Print(consolecolors.Blue(`
light and fast , dependency injection based micro-service framework written in Go.
`))

	fmt.Println(" ")
	l.Info(consolecolors.Green("Welcome to YoyoGo, starting application ..."))
	l.Info("yoyogo framework version :  %s", consolecolors.Blue(env.Version))
	l.Info("server & protocol        :  %s", consolecolors.Green(env.Server))
	l.Info("machine host ip          :  %s", consolecolors.Blue(env.Host))
	l.Info("listening on port        :  %s", consolecolors.Blue(env.Port))
	l.Info("application running pid  :  %s", consolecolors.Blue(strconv.Itoa(env.PID)))
	l.Info("application name         :  %s", consolecolors.Blue(env.ApplicationName))
	l.Info("application exec path    :  %s", consolecolors.Yellow(utils.GetCurrentDirectory()))
	l.Info("application config path  :  %s", consolecolors.Yellow(env.MetaData["config.path"]))
	l.Info("application environment  :  %s", consolecolors.Yellow(consolecolors.Blue(env.Profile)))
	l.Info("running in %s mode , change (Dev,tests,Prod) mode by HostBuilder.SetEnvironment .", consolecolors.Red(env.Profile))
	l.Info(consolecolors.Green("Starting server..."))
	l.Info("server setting map       :  %v", env.MetaData)
	l.Info(consolecolors.Green("Server is Started."))
}
