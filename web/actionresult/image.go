package actionresult

import "net/http"

type Image struct {
	Data []byte
}

// WriteContentType (Data) writes custom ContentType.
func (r Image) WriteContentType(w http.ResponseWriter) {
	writeContentType(w, []string{"image/png"})
}

func (r Image) Render(w http.ResponseWriter) (err error) {
	r.WriteContentType(w)
	_, err = w.Write(r.Data)
	return
}
