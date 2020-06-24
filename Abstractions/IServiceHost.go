package Abstractions

import (
	"encoding/base64"
	"fmt"
	"github.com/maxzhang1985/yoyogo"
	"github.com/maxzhang1985/yoyogo/Abstractions/Platform/ConsoleColors"
	"github.com/maxzhang1985/yoyogo/WebFramework/Context"
	"log"
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
	fmt.Println("")
	l.Println("Welcome to YoyoGo, starting application ...")
	l.Printf("version: %s", env.Version)
	l.Printf("listening on %s", env.Port)
	l.Printf("application is running , pid: %d", env.PID)
	l.Printf("running in %s mode , change mode by app.SetEnvironment .", env.Profile)
	l.Println("Starting HTTP server...")
}
