package Abstractions

type IConfiguration interface {
	Get(name string) interface{}
	GetString(name string) string
	GetInt(name string) int
	GetSection(name string) IConfiguration
	Unmarshal(interface{})
	GetProfile() string
	GetConfDir() string
}
