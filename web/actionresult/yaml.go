package actionresult

import (
	"gopkg.in/yaml.v2"
	"net/http"
)

type YAML struct {
	Data interface{}
}

var yamlContentType = []string{"application/x-yaml; charset=utf-8"}

func (r YAML) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, yamlContentType)
}

func (r YAML) Render(w http.ResponseWriter) error {
	r.WriteContentType(w)

	bytes, err := yaml.Marshal(r.Data)
	if err != nil {
		return err
	}

	_, err = w.Write(bytes)
	return err
}
