package Context

import "net/http"

type Output struct {
	Response *responseWriter
}

//Set Cookie value
func (output Output) SetCookie(name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	}
	output.Response.Header().Add("Set-Cookie", cookie.String())
}

func (output Output) Status() int {
	return output.Response.Status()
}

func (output Output) GetWriter() *responseWriter {
	return output.Response
}

func (output Output) SetStatus(status int) {
	output.Response.WriteHeader(status)
}

func (output Output) SetStatusCode(status int) {
	output.Response.SetStatus(status)
}

func (output Output) SetStatusCodeNow() {
	output.Response.WriteHeaderNow()
}

// Write Byte[] Response.
func (output Output) Write(data []byte) (n int, err error) {
	return output.Response.Write(data)
}

func (output Output) Header(key, value string) {
	if value == "" {
		output.Response.Header().Del(key)
		return
	}
	output.Response.Header().Set(key, value)
}

// Write Error Response.
func (output Output) Error(code int, error string) {
	http.Error(output.Response, error, code)
}
