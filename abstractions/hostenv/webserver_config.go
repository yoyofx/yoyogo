package hostenv

type WebServerConfig struct {
	ServerType     string           `mapstructure:"type" config:"type"`
	Address        string           `mapstructure:"address" config:"address"`
	MaxRequestSize int64            `mapstructure:"max_request_size" config:"max_request_size"`
	Static         StaticConfig     `mapstructure:"static" config:"static"`
	Tls            HttpServerConfig `mapstructure:"tls" config:"tls"`
}

type StaticConfig struct {
	Patten  string `mapstructure:"patten" config:"patten"`
	WebRoot string `mapstructure:"webroot" config:"webroot"`
}
