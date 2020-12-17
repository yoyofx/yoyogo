package hostenv

type HostConfig struct {
	Name     string          `mapstructure:"name"`
	Metadata string          `mapstructure:"metadata"`
	Server   WebServerConfig `mapstructure:"server"`
}
