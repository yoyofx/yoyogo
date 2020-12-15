package abstractions

type IConfiguration interface {
	Get(name string) interface{}
	GetString(name string) string
	GetBool(name string) bool
	GetInt(name string) int
	GetSection(name string) IConfiguration
	Unmarshal(interface{})
	GetProfile() string
	GetConfDir() string
}
