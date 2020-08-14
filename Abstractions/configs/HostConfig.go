package configs

type HostConfig struct {
	Name     string       `mapstructure:"name"`
	Metadata string       `mapstructure:"metadata"`
	Server   ServerConfig `mapstructure:"server"`
}
