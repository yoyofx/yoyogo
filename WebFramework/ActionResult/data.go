package ActionResult

import (
	"io/ioutil"
	"net/http"
)

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

var octetstreamContentType = "application/octet-stream"

func FormFileStream(data []byte) Data {
	return Data{ContentType: octetstreamContentType, Data: data}
}

func FormFile(filename string) (Data, error) {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		return Data{}, err
	}
	return Data{ContentType: octetstreamContentType, Data: bytes}, nil
}
