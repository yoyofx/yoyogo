package models

type MyConfig struct {
	Name     string
	Url      string
	UserName string
	Password string //`mapstructure:"password"
	Debug    bool
}

func (config MyConfig) GetSection() string {
	return "yoyogo.datasource.db"
}
