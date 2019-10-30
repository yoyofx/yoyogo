package YoyoGo

import "os"

type IServer interface {
	Run(addr string) error
	RunOverTLS(addr, certFile, keyFile string) error
}

func detectAddress(addr ...string) string {
	if len(addr) > 0 {
		return addr[0]
	}
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return DefaultAddress
}
