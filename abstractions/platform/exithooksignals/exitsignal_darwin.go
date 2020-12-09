package exithooksignals

import (
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func HookSignals(host abstractions.IServiceHost) {
	quitSig := make(chan os.Signal)
	signal.Notify(
		quitSig,
		syscall.SIGHUP,
		syscall.SIGINT,
		syscall.SIGTERM,
		syscall.SIGQUIT,
		syscall.SIGSTOP,
		syscall.SIGUSR1,
		syscall.SIGUSR2,
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
