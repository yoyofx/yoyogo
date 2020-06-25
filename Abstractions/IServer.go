package Abstractions

import (
	"github.com/yoyofx/yoyogo"
	"os"
)

type IServer interface {
	GetAddr() string
	Run(context *HostBuildContext) (e error)
	Shutdown()
}

func DetectAddress(addr ...string) string {
	if len(addr) > 0 && addr[0] != "" {
		return addr[0]
	}
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return YoyoGo.DefaultAddress
}
