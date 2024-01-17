package swagger

type Path struct {
	Tags        []string                 `json:"tags"`
	Summary     string                   `json:"summary"`
	OperationId string                   `json:"operationId"`
	Parameters  []Parameters             `json:"parameters"`
	RequestBody RequestBody              `json:"requestBody"`
	Responses   map[string]ResponsesItem `json:"responses"`
	Security    []Security               `json:"security"`
}

type Parameters struct {
	Name        string `json:"name"`
	In          string `json:"in"`
	Description string `json:"description"`
	Required    bool   `json:"required"`
	Schema      struct {
		Type string `json:"type"`
	} `json:"schema"`
}
type RequestBody struct {
	Content map[string]ContentType `json:"content"`
}

type ContentType struct {
	Schema Schema `json:"schema"`
}

type Schema struct {
	Type       string              `json:"type"`
	Properties map[string]Property `json:"properties"`
	Required   []string            `json:"required"`
}
type Property struct {
	Description string `json:"description"`
	Type        string `json:"type"`
}

type ResponsesItem struct {
	Description string                 `json:"description"`
	Content     map[string]interface{} `json:"content"`
}

type Security struct {
	PetstoreAuth []string `json:"petstore_auth"`
}