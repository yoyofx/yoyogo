package Abstractions

type IConfiguration interface {
	Get(name string) interface{}
	GetSection(name string) IConfiguration
	Unmarshal(interface{})
}
