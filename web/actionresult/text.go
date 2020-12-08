package actionresult

import (
	"fmt"
	"io"
	"net/http"
)

type Text struct {
	Format string
	Data   []interface{}
}

var plainContentType = []string{"text/plain; charset=utf-8"}

func (r Text) Render(w http.ResponseWriter) (err error) {
	writeContentType(w, plainContentType)
	if len(r.Data) > 0 {
		_, err = fmt.Fprintf(w, r.Format, r.Data...)
		return
	}
	_, err = io.WriteString(w, r.Format)
	return
}

func (r Text) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, plainContentType)
}
