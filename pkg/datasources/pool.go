package datasources

type DataSourcePool struct {
	InitCap     int `mapstructure:"init_cap"`
	MaxCap      int `mapstructure:"max_cap"`
	Idletimeout int `mapstructure:"idle_timeout"`
}
