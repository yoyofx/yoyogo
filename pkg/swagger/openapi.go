package swagger

type OpenApi struct {
	Openapi string                     `json:"openapi"`
	Info    Info                       `json:"info"`
	Paths   map[string]map[string]Path `json:"paths"`
}
