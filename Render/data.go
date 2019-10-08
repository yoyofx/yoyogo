package Render

import "net/http"

// Data contains ContentType and bytes data.
type Data struct {
	ContentType string
	Data        []byte
}

// WriteContentType (Data) writes custom ContentType.
func (r Data) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, []string{r.ContentType})
}

func (r Data) Render(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)
	_, err = w.Write(r.Data)
	return
}
