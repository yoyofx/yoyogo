package Middleware

import (
	"encoding/json"
	"io"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path"
	"sync"
)

type HttpContext struct {
	Req        *http.Request
	Resp       *responseWriter
	store      map[string]interface{}
	storeMutex *sync.RWMutex
}

func NewContext(w http.ResponseWriter, r *http.Request) *HttpContext {
	ctx := &HttpContext{}
	ctx.init(w, r)
	return ctx
}

func (ctx *HttpContext) init(w http.ResponseWriter, r *http.Request) {
	ctx.storeMutex = new(sync.RWMutex)
	ctx.Resp = &responseWriter{w, 0, 0, nil}
	ctx.Req = r
	ctx.storeMutex.Lock()
	ctx.store = nil
	ctx.storeMutex.Unlock()
}

//Set data in context.
func (ctx *HttpContext) SetItem(key string, val interface{}) {
	ctx.storeMutex.Lock()
	if ctx.store == nil {
		ctx.store = make(map[string]interface{})
	}
	ctx.store[key] = val
	ctx.storeMutex.Unlock()
}

// Get data in context.
func (ctx *HttpContext) GetItem(key string) interface{} {
	ctx.storeMutex.RLock()
	v := ctx.store[key]
	ctx.storeMutex.RUnlock()
	return v
}

//Set Cookie value
func (ctx *HttpContext) SetCookie(name, value string) {
	cookie := &http.Cookie{
		Name:     name,
		Value:    value,
		Path:     "/",
		MaxAge:   0,
		HttpOnly: true,
	}
	ctx.Resp.Header().Add("Set-Cookie", cookie.String())
}

//Get Cookie by Name
func (ctx *HttpContext) GetCookie(name string) string {
	cookie, err := ctx.Req.Cookie(name)
	if err != nil {
		return ""
	}
	return url.QueryEscape(cookie.Value)
}

//Get Post Params
func (ctx *HttpContext) Params() url.Values {
	return ctx.Req.Form
}

//Get Post Param
func (ctx *HttpContext) Param(name string) string {
	if ctx.Params()[name] != nil {
		return ctx.Params()[name][0]
	}
	return ""
}

// Get Query string
func (ctx *HttpContext) QueryString() url.Values {
	return ctx.Req.URL.Query()
}

// Redirect redirects the request
func (ctx *HttpContext) Redirect(code int, url string) {
	http.Redirect(ctx.Resp, ctx.Req, url, code)
}

// Path returns URL Path string.
func (ctx *HttpContext) Path() string {
	return ctx.Req.URL.Path
}

// Referer returns request referer.
func (ctx *HttpContext) Referer() string {
	return ctx.Req.Header.Get("Referer")
}

// UserAgent returns http request UserAgent
func (ctx *HttpContext) UserAgent() string {
	return ctx.Req.Header.Get("User-Agent")
}

//Get Http Method.
func (ctx *HttpContext) Method() string {
	return ctx.Req.Method
}

//Get Http Status Code.
func (ctx *HttpContext) Status() int {
	return ctx.Resp.status
}

// FormFile gets file from request.
func (ctx *HttpContext) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return ctx.Req.FormFile(key)
}

// SaveFile saves the form file and
// returns the filename.
func (ctx *HttpContext) SaveFile(name, saveDir string) (string, error) {
	fr, handle, err := ctx.FormFile(name)
	if err != nil {
		return "", err
	}
	defer fr.Close()
	fw, err := os.OpenFile(path.Join(saveDir, handle.Filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return "", err
	}
	defer fw.Close()

	_, err = io.Copy(fw, fr)
	return handle.Filename, err
}

// Write Error Response.
func (ctx *HttpContext) Error(code int, error string) {
	http.Error(ctx.Resp, error, code)
}

// Write Byte[] Response.
func (ctx *HttpContext) Write(data []byte) (n int, err error) {
	return ctx.Resp.Write(data)
}

// Text response text data.
func (ctx *HttpContext) Text(code int, body string) error {
	ctx.Resp.Header().Set("Content-Type", "text/plain; charset=utf-8")
	ctx.Resp.WriteHeader(code)
	_, err := ctx.Resp.Write([]byte(body))
	return err
}

// Write Json Response.
func (ctx *HttpContext) JSON(data interface{}) {
	ctx.Resp.Header().Set("Content-Type", "application/json")
	jsons, _ := json.Marshal(data)
	_, _ = ctx.Resp.Write(jsons)
}

// JSONP return JSONP data.
func (ctx *HttpContext) JSONP(code int, callback string, data interface{}) error {
	j, err := json.Marshal(data)
	if err != nil {
		return err
	}
	ctx.Resp.Header().Set("Content-Type", "application/javascript; charset=utf-8")
	ctx.Resp.WriteHeader(code)
	_, _ = ctx.Resp.Write([]byte(callback + "("))
	_, _ = ctx.Resp.Write(j)
	_, _ = ctx.Resp.Write([]byte(");"))
	return nil
}
