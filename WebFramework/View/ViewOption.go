package View

type Option struct {
	Path     string   `mapstructure:"path"`
	Includes []string `mapstructure:"includes"`
}
