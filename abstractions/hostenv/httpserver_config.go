package hostenv

// http server configuration
type HttpServerConfig struct {
	IsTLS    bool
	Addr     string
	CertFile string `mapstructure:"cert" config:"cert"`
	KeyFile  string `mapstructure:"key" config:"key"`
}
