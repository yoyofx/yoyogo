package web

// http server configuration
type ServerConfig struct {
	IsTLS                   bool
	Addr, CertFile, KeyFile string
}
