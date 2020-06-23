package ActionResult

import "net/http"

// ActionResult interface is to be implemented by JSON, XML, HTML, YAML and so on.
type IActionResult interface {
	// ActionResult writes data with custom ContentType.
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
