package context

import (
	"bufio"
	"errors"
	"net"
	"net/http"
)

// IResponseWriter is a wrapper around http.ResponseWriter that provides extra information about
// the response. It is recommended that middleware handlers use this construct to wrap a responsewriter
// if the functionality calls for it.
type IResponseWriter interface {
	http.ResponseWriter
	http.Flusher
	// Status returns the status code of the response or 0 if the response has
	// not been written
	Status() int
	// Written returns whether or not the IResponseWriter has been written.
	Written() bool
	// Size returns the size of the response body.
	Size() int
	// Before allows for a function to be called before the IResponseWriter has been written to. This is
	// useful for setting headers or any other operations that must happen before a response has been written.
	Before(func(IResponseWriter))
}

type beforeFunc func(IResponseWriter)

// NewResponseWriter creates a IResponseWriter that wraps an http.ResponseWriter
func NewResponseWriter(rw http.ResponseWriter) IResponseWriter {
	nrw := &CResponseWriter{
		ResponseWriter: rw,
	}

	if _, ok := rw.(http.CloseNotifier); ok {
		return &responseWriterCloseNotifer{nrw}
	}

	return nrw
}

type CResponseWriter struct {
	http.ResponseWriter
	status      int
	size        int
	beforeFuncs []beforeFunc
}

func (rw *CResponseWriter) SetStatus(code int) {
	rw.status = code
}

func (rw *CResponseWriter) WriteHeader(s int) {
	rw.status = s
	rw.callBefore()
	rw.ResponseWriter.WriteHeader(s)
}

func (w *CResponseWriter) WriteHeaderNow() {
	if !w.Written() {
		w.size = 0
		w.ResponseWriter.WriteHeader(w.status)
	}
}

func (rw *CResponseWriter) Write(b []byte) (int, error) {
	//if !rw.Written() {
	//	// The status will be StatusOK if WriteHeader has not been called yet
	//	rw.WriteHeader(http.StatusOK)
	//}
	size, err := rw.ResponseWriter.Write(b)
	rw.size += size
	return size, err
}

func (rw *CResponseWriter) Status() int {
	return rw.status
}

func (rw *CResponseWriter) Size() int {
	return rw.size
}

func (rw *CResponseWriter) Written() bool {
	return rw.status != 0
}

func (rw *CResponseWriter) Before(before func(IResponseWriter)) {
	rw.beforeFuncs = append(rw.beforeFuncs, before)
}

func (rw *CResponseWriter) Hijack() (net.Conn, *bufio.ReadWriter, error) {
	hijacker, ok := rw.ResponseWriter.(http.Hijacker)
	if !ok {
		return nil, nil, errors.New("the IResponseWriter doesn't support the Hijacker interface")
	}
	return hijacker.Hijack()
}

func (rw *CResponseWriter) callBefore() {
	for i := len(rw.beforeFuncs) - 1; i >= 0; i-- {
		rw.beforeFuncs[i](rw)
	}
}

func (rw *CResponseWriter) Flush() {
	flusher, ok := rw.ResponseWriter.(http.Flusher)
	if ok {
		if !rw.Written() {
			// The status will be StatusOK if WriteHeader has not been called yet
			rw.WriteHeader(http.StatusOK)
		}
		flusher.Flush()
	}
}

type responseWriterCloseNotifer struct {
	*CResponseWriter
}

func (rw *responseWriterCloseNotifer) CloseNotify() <-chan bool {
	return rw.ResponseWriter.(http.CloseNotifier).CloseNotify()
}
