package configs

type ServerConfig struct {
	ServerType     string       `mapstructure:"type"`
	Address        string       `mapstructure:"address"`
	MaxRequestSize string       `mapstructure:"max_request_size"`
	Static         StaticConfig `mapstructure:"static"`
}

type StaticConfig struct {
	Patten  string `mapstructure:"patten"`
	WebRoot string `mapstructure:"webroot"`
}
