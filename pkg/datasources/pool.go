package datasources

type DataSourcePool struct {
	InitCap     int `mapstructure:"init_cap" config:"init_cap"`
	MaxCap      int `mapstructure:"max_cap" config:"max_cap"`
	Idletimeout int `mapstructure:"idle_timeout" config:"idle_timeout"`
}
