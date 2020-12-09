package actionresult

import "net/http"

// actionresult interface is to be implemented by JSON, XML, HTML, YAML and so on.
type IActionResult interface {
	// actionresult writes data with custom ContentType.
	Render(http.ResponseWriter) error
	// WriteContentType writes custom ContentType.
	WriteContentType(w http.ResponseWriter)
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		w.Header().Set("content-type", value[0])
		//header["Content-Type"] = value
	}
}
