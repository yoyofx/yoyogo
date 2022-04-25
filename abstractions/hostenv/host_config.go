package hostenv

type HostConfig struct {
	Name     string          `mapstructure:"name" config:"name"`
	Metadata string          `mapstructure:"metadata" config:"metadata"`
	Server   WebServerConfig `mapstructure:"server" config:"server"`
}
