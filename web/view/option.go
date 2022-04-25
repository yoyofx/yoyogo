package view

type Option struct {
	Path     string   `mapstructure:"path" config:"path"`
	Includes []string `mapstructure:"includes" config:"includes"`
}
