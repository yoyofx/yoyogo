package configs

type HostConfig struct {
	Name     string       `mapstructure:"name"`
	Metadata string       `mapstructure:"metadata"`
	Profile  string       `mapstructure:"profile"`
	Server   ServerConfig `mapstructure:"server"`
}
