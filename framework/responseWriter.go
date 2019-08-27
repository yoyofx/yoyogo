package YoyoGo

import (
	"bufio"
	"net"
	"net/http"
)

// responseWriter implement ResponseWriter
type responseWriter struct {
	http.ResponseWriter
	status int
}

// WriteHeader sends an HTTP response header with status code,
// and stores the code
func (w *responseWriter) WriteHeader(code int) {
	w.status = code
	w.ResponseWriter.WriteHeader(code)
}

// Hijack implements the http.Hijacker interface to allow an HTTP handler to
// take over the connection
func (w *responseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	return w.ResponseWriter.(http.Hijacker).Hijack()
}
