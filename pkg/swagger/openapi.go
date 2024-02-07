package swagger

type OpenApi struct {
	Openapi    string                     `json:"openapi"`
	Info       Info                       `json:"info"`
	Host       string                     `json:"host"`
	BasePath   string                     `json:"basePath"`
	Schemes    []string                   `json:"schemes"`
	Servers    []Server                   `json:"servers"`
	Tags       []Tag                      `json:"tags"`
	Paths      map[string]map[string]Path `json:"paths"`
	Components map[string]interface{}     `json:"components"`
	Security   []map[string][]string      `json:"security"`
}

func NewOpenApi(baseUri string, info Info) *OpenApi {
	return &OpenApi{
		Openapi:    "3.0.0",
		Info:       info,
		Host:       baseUri,
		Servers:    []Server{{Url: baseUri}},
		Paths:      make(map[string]map[string]Path),
		Components: make(map[string]interface{}),
		Security:   make([]map[string][]string, 0),
	}
}

func (openapi *OpenApi) AddSecurityApiKey(name string) {
	openapi.Components[name] = map[string]string{
		"type": "apiKey",
		"name": name,
		"in":   "header",
	}
}

func (openapi *OpenApi) AddSecurityBasicAuth() {
	openapi.Components["securitySchemes"] = map[string]interface{}{
		"basicAuth": map[string]string{
			"type":   "http",
			"scheme": "basic",
		},
	}
}

func (openapi *OpenApi) AddSecurityBearerAuth() {
	openapi.Components["securitySchemes"] = map[string]interface{}{
		"bearerAuth": map[string]string{
			"type":         "http",
			"scheme":       "bearer",
			"bearerFormat": "JWT",
		},
	}

	openapi.Security = append(openapi.Security, map[string][]string{
		"bearerAuth": {}})
}
