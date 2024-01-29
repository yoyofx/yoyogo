package swagger

type OpenApi struct {
	Openapi string                     `json:"openapi"`
	Info    Info                       `json:"info"`
	Servers []Server                   `json:"servers"`
	Tags    []Tag                      `json:"tags"`
	Paths   map[string]map[string]Path `json:"paths"`
}
