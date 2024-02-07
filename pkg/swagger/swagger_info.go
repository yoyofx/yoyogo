package swagger

// Info swagger info 根节点/**
type Info struct {
	Title          string  `json:"title"`
	Description    string  `json:"description"`
	TermsOfService string  `json:"termsOfService"`
	Contact        Contact `json:"contact"`
	License        License `json:"license"`
	Version        string  `json:"version"`
}

type Contact struct {
	Name  string `json:"name"`
	Url   string `json:"url"`
	Email string `json:"email"`
}

type License struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type Server struct {
	Url string `json:"url"`
}

type Tag struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	//ExternalDocs struct {
	//	Description string `json:"description"`
	//	Url         string `json:"url"`
	//}
}
