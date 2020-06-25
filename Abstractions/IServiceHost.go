package Abstractions

import (
	"encoding/base64"
	"fmt"
	"github.com/yoyofx/yoyogo"
	"github.com/yoyofx/yoyogo/Abstractions/Platform/ConsoleColors"
	"github.com/yoyofx/yoyogo/Utils"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"log"
	"strconv"
)

type IServiceHost interface {
	Run()
	Shutdown()
	StopApplicationNotify()
	SetAppMode(mode string)
}

func PrintLogo(l *log.Logger, env *Context.HostEnvironment) {
	logo, _ := base64.StdEncoding.DecodeString(YoyoGo.Logo)

	fmt.Println(ConsoleColors.Blue(string(logo)))
	fmt.Println(" ")
	l.Println(ConsoleColors.Green("Welcome to YoyoGo, starting application ..."))
	l.Printf("yoyogo framework version :  %s", ConsoleColors.Blue(env.Version))
	l.Printf("machine host ip          :  %s", ConsoleColors.Blue(env.Host))
	l.Printf("listening on port        :  %s", ConsoleColors.Blue(env.Port))
	l.Printf("application running pid  :  %s", ConsoleColors.Blue(strconv.Itoa(env.PID)))
	l.Printf("application environment  :  %s", ConsoleColors.Blue(env.Profile))
	l.Printf("application exec path    :  %s", ConsoleColors.Yellow(Utils.GetCurrentDirectory()))
	l.Printf("running in %s mode , change (Dev,Test,Prod) mode by HostBuilder.SetEnvironment .", ConsoleColors.Blue(env.Profile))
	l.Println(ConsoleColors.Green("Starting HTTP server..."))

}
