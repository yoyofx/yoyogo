package abstractions

import (
	"os"
)

const (
	// DefaultAddress is used if no other is specified.
	DefaultAddress = ":8080"
)

type IServer interface {
	GetAddr() string
	Run(context *HostBuilderContext) (e error)
	Shutdown()
}

func DetectAddress(addr ...string) string {
	if len(addr) > 0 && addr[0] != "" {
		return addr[0]
	}
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return DefaultAddress
}
