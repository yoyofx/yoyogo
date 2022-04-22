package models

type MyConfig struct {
	Name     string
	Url      string
	UserName string
	Password string //`mapstructure:"password"
	Debug    bool
	Env      string
}

func (config MyConfig) GetSection() string {
	return "yoyogo.datasource.db"
}
