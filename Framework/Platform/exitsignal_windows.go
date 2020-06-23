package Platform

import (
	"fmt"
	"github.com/maxzhang1985/yoyogo/Abstract"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func HookSignals(host Abstract.IServiceHost) {
	quitSig := make(chan os.Signal)
	signal.Notify(
		quitSig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGKILL,
	)

	go func() {
		var sig os.Signal
		for {
			sig = <-quitSig
			fmt.Println()
			switch sig {
			case syscall.SIGQUIT:
				host.StopApplicationNotify()
				host.Shutdown()
				os.Exit(0)
				// graceful stop
			case syscall.SIGHUP:
				host.StopApplicationNotify()
				host.Shutdown()
			case syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM:
				host.StopApplicationNotify()
				host.Shutdown()
				os.Exit(0)
				// terminate now
			}
			time.Sleep(time.Second * 3)
		}
	}()
}
