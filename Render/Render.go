package Render

import "net/http"

// Render interface is to be implemented by JSON, XML, HTML, YAML and so on.
type Render interface {
	// Render writes data with custom ContentType.
	Render(http.ResponseWriter) error
	// WriteContentType writes custom ContentType.
	WriteContentType(w http.ResponseWriter)
}

func writeContentType(w http.ResponseWriter, value []string) {
	header := w.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = value
	}
}

//var (
//	_ Render     = JSON{}
//	_ Render     = IndentedJSON{}
//	_ Render     = SecureJSON{}
//	_ Render     = JsonpJSON{}
//	_ Render     = XML{}
//	_ Render     = String{}
//	_ Render     = Redirect{}
//	_ Render     = Data{}
//	_ Render     = HTML{}
//	_ HTMLRender = HTMLDebug{}
//	_ HTMLRender = HTMLProduction{}
//	_ Render     = YAML{}
//	_ Render     = MsgPack{}
//	_ Render     = Reader{}
//	_ Render     = AsciiJSON{}
//	_ Render     = ProtoBuf{}
//)
