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
	if err == nil {
		_ = sd.Destroy()
	}
}

func PrintLogo(l xlog.ILogger, env *HostEnvironment) {
	logo, _ := base64.StdEncoding.DecodeString(yoyogo.Logo)

	fmt.Println(consolecolors.Blue(string(logo)))
	fmt.Printf("%s                   (%s)", consolecolors.Green(":: YoyoGo ::"), consolecolors.Blue(env.Version))
	fmt.Println(" ")
	fmt.Println(" ")
	l.Debug(consolecolors.Green("Welcome to YoyoGo, starting application ..."))
	l.Debug("yoyogo framework version :  %s", consolecolors.Blue(env.Version))
	l.Debug("machine host ip          :  %s", consolecolors.Blue(env.Host))
	l.Debug("listening on port        :  %s", consolecolors.Blue(env.Port))
	l.Debug("application running pid  :  %s", consolecolors.Blue(strconv.Itoa(env.PID)))
	l.Debug("application name         :  %s", consolecolors.Blue(env.ApplicationName))
	l.Debug("application exec path    :  %s", consolecolors.Yellow(utils.GetCurrentDirectory()))
	l.Debug("application config path  :  %s", consolecolors.Yellow(env.MetaData["config.path"]))
	l.Debug("application environment  :  %s", consolecolors.Yellow(consolecolors.Blue(env.Profile)))
	l.Debug("running in %s mode , change (Dev,tests,Prod) mode by HostBuilder.SetEnvironment .", consolecolors.Red(env.Profile))
	l.Debug(consolecolors.Green("Starting server..."))
	l.Debug("server setting map       :  %v", env.MetaData)

}
