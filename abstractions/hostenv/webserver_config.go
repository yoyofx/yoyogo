package hostenv

type WebServerConfig struct {
	ServerType     string           `mapstructure:"type"`
	Address        string           `mapstructure:"address"`
	MaxRequestSize int64            `mapstructure:"max_request_size"`
	Static         StaticConfig     `mapstructure:"static"`
	Tls            HttpServerConfig `mapstructure:"tls"`
}

type StaticConfig struct {
	Patten  string `mapstructure:"patten"`
	WebRoot string `mapstructure:"webroot"`
}
