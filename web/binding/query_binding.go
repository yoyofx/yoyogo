package binding

import (
	"fmt"
	"net/http"
)

type queryBinding struct{}

func (queryBinding) Name() string {
	return "query"
}

func (queryBinding) Bind(req *http.Request, obj interface{}) error {
	values := req.URL.Query()
	fmt.Println(values)
	if err := mapForm(obj, values); err != nil {
		return err
	}
	return validate(obj)
}
