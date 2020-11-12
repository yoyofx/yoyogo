package Abstractions

import (
	"encoding/base64"
	"fmt"
	"github.com/yoyofx/yoyogo"
	"github.com/yoyofx/yoyogo/Abstractions/Platform/ConsoleColors"
	"github.com/yoyofx/yoyogo/Abstractions/ServiceDiscovery"
	"github.com/yoyofx/yoyogo/Abstractions/XLog"
	"github.com/yoyofx/yoyogo/Utils"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"strconv"
)

type IServiceHost interface {
	Run()
	Shutdown()
	StopApplicationNotify()
	SetAppMode(mode string)
}

func HostRunning(log XLog.ILogger, context *HostBuildContext) {
	go startServerDiscovery(log, context)
}

func HostEnding(log XLog.ILogger, context *HostBuildContext) {
	endServerDiscovery(log, context)
}

func startServerDiscovery(log XLog.ILogger, context *HostBuildContext) {
	var sd ServiceDiscovery.IServiceDiscovery
	_ = context.HostServices.GetService(&sd)
	if sd != nil {
		_ = sd.Register()
	}
}

func endServerDiscovery(log XLog.ILogger, context *HostBuildContext) {
	var sd ServiceDiscovery.IServiceDiscovery
	_ = context.HostServices.GetService(&sd)
	if sd != nil {
		_ = sd.Destroy()
	}
}

func PrintLogo(l XLog.ILogger, env *Context.HostEnvironment) {
	logo, _ := base64.StdEncoding.DecodeString(YoyoGo.Logo)

	fmt.Println(ConsoleColors.Blue(string(logo)))
	fmt.Printf("%s                   (%s)", ConsoleColors.Green(":: YoyoGo ::"), ConsoleColors.Blue(env.Version))
	fmt.Println(" ")
	fmt.Println(" ")
	l.Debug(ConsoleColors.Green("Welcome to YoyoGo, starting application ..."))
	l.Debug("yoyogo framework version :  %s", ConsoleColors.Blue(env.Version))
	l.Debug("machine host ip          :  %s", ConsoleColors.Blue(env.Host))
	l.Debug("listening on port        :  %s", ConsoleColors.Blue(env.Port))
	l.Debug("application running pid  :  %s", ConsoleColors.Blue(strconv.Itoa(env.PID)))
	l.Debug("application name         :  %s", ConsoleColors.Blue(env.ApplicationName))
	l.Debug("application exec path    :  %s", ConsoleColors.Yellow(Utils.GetCurrentDirectory()))
	l.Debug("application config path  :  %s", ConsoleColors.Yellow(env.MetaData["config.path"]))
	l.Debug("application environment  :  %s", ConsoleColors.Yellow(ConsoleColors.Blue(env.Profile)))
	l.Debug("running in %s mode , change (Dev,Test,Prod) mode by HostBuilder.SetEnvironment .", ConsoleColors.Red(env.Profile))
	l.Debug(ConsoleColors.Green("Starting server..."))
	l.Debug("server setting map       :  %v", env.MetaData)

}
