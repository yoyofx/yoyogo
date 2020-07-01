package Context

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/yoyofx/yoyogo/DependencyInjection"
	"github.com/yoyofx/yoyogo/Utils"
	"github.com/yoyofx/yoyogo/WebFramework/ActionResult"
	"io"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
	"sync"
)

const (
	defaultTagName = "param"
	jsonTagName    = "json"
)

type M = map[string]string

type HttpContext struct {
	Request          *http.Request
	Response         *responseWriter
	RouterData       url.Values
	RequiredServices DependencyInjection.IServiceProvider
	store            map[string]interface{}
	storeMutex       *sync.RWMutex
	Result           interface{}
}

func NewContext(w http.ResponseWriter, r *http.Request, sp DependencyInjection.IServiceProvider) *HttpContext {
	ctx := &HttpContext{}
	ctx.init(w, r, sp)
	return ctx
}

func (ctx *HttpContext) init(w http.ResponseWriter, r *http.Request, sp DependencyInjection.IServiceProvider) {
	ctx.storeMutex = new(sync.RWMutex)
	ctx.Response = &responseWriter{w, 0, 0, nil}
	ctx.Request = r
	ctx.RouterData = url.Values{}
	ctx.RequiredServices = sp
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
	ctx.Response.Header().Add("Set-Cookie", cookie.String())
}

//Get Cookie by Name
func (ctx *HttpContext) GetCookie(name string) string {
	cookie, err := ctx.Request.Cookie(name)
	if err != nil {
		return ""
	}
	return url.QueryEscape(cookie.Value)
}

func (c *HttpContext) Header(key, value string) {
	if value == "" {
		c.Response.Header().Del(key)
		return
	}
	c.Response.Header().Set(key, value)
}

//Get Post Params
func (ctx *HttpContext) PostForm() url.Values {
	_ = ctx.Request.ParseForm()
	return ctx.Request.PostForm
}

func (ctx *HttpContext) PostMultipartForm() url.Values {
	_ = ctx.Request.ParseMultipartForm(32 << 20)
	return ctx.Request.MultipartForm.Value
}

func (ctx *HttpContext) PostJsonForm() url.Values {
	ret := url.Values{}
	var jsonMap map[string]interface{}
	body := ctx.PostBody()
	_ = json.Unmarshal(body, &jsonMap)
	var strVal string
	for key, value := range jsonMap {
		switch value.(type) {
		case int32:
		case int64:
			strVal = strconv.Itoa(value.(int))
			break
		case float64:
			strVal = strconv.FormatFloat(value.(float64), 'f', -1, 64)
			break
		default:
			strVal = fmt.Sprint(value)
		}
		ret.Add(key, strVal)
	}
	return ret
}

func (ctx *HttpContext) GetAllParam() url.Values {
	form := url.Values{}
	content_type := ctx.Request.Header.Get("Content-Type")

	if strings.HasPrefix(content_type, MIMEApplicationForm) {
		form = ctx.PostForm()
	} else if strings.HasPrefix(content_type, MIMEMultipartForm) {
		form = ctx.PostMultipartForm()
	} else if strings.HasPrefix(content_type, MIMEApplicationJSON) {
		form = ctx.PostJsonForm()
	}

	Utils.MergeMap(form, ctx.QueryStrings())
	Utils.MergeMap(form, ctx.RouterData)
	return form
}

//Get Post Param
func (ctx *HttpContext) Param(name string) string {
	form := ctx.GetAllParam()
	if form[name] != nil {
		return form[name][0]
	}
	return ""
}

// PostBody returns data from the POST or PUT request body
func (ctx *HttpContext) PostBody() []byte {
	bts, err := ioutil.ReadAll(ctx.Request.Body)
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bts))
	if err != nil {
		return []byte{}
	}

	return bts
}

func (ctx *HttpContext) Bind(i interface{}) (err error) {
	req := ctx.Request
	ctype := req.Header.Get(HeaderContentType)
	if req.Body == nil {
		err = errors.New("request body can't be empty")
		return err
	}
	err = errors.New("request unsupported MediaType -> " + ctype)
	tagName := defaultTagName
	switch {
	case strings.HasPrefix(ctype, MIMEApplicationXML):
		err = xml.Unmarshal(ctx.PostBody(), i)
	case strings.HasPrefix(ctype, MIMEApplicationJSON):
		//tagName = jsonTagName
	default:
		// check is use json tag, fixed for issue #91
		//tagName = defaultTagName
		// no check content type for fixed issue #6

	}
	err = ConvertMapToStruct(tagName, i, ctx.GetAllParam())
	return err
}

// RemoteIP RemoteAddr to an "IP" address
func (ctx *HttpContext) RemoteIP() string {
	host, _, _ := net.SplitHostPort(ctx.Request.RemoteAddr)
	return host
}

// RealIP returns the first ip from 'X-Forwarded-For' or 'X-Real-IP' header key
// if not exists data, returns request.RemoteAddr
// fixed for #164
func (ctx *HttpContext) RealIP() string {
	if ip := ctx.Request.Header.Get(HeaderXForwardedFor); ip != "" {
		return strings.Split(ip, ", ")[0]
	}
	if ip := ctx.Request.Header.Get(HeaderXRealIP); ip != "" {
		return ip
	}
	host, _, _ := net.SplitHostPort(ctx.Request.RemoteAddr)
	return host
}

// FullRemoteIP RemoteAddr to an "IP:port" address
func (ctx *HttpContext) FullRemoteIP() string {
	fullIp := ctx.Request.RemoteAddr
	return fullIp
}

// IsAJAX returns if it is a ajax request
func (ctx *HttpContext) IsAJAX() bool {
	return strings.Contains(ctx.Request.Header.Get(HeaderXRequestedWith), "XMLHttpRequest")
}

func (ctx *HttpContext) IsWebsocket() bool {
	if strings.Contains(strings.ToLower(ctx.Request.Header.Get("Connection")), "upgrade") &&
		strings.EqualFold(ctx.Request.Header.Get("Upgrade"), "websocket") {
		return true
	}
	return false
}

// Url get request url
func (ctx *HttpContext) Url() string {
	return ctx.Request.URL.String()
}

// Get Query string
func (ctx *HttpContext) QueryStrings() url.Values {

	queryForm, err := url.ParseQuery(ctx.Request.URL.RawQuery)
	if err == nil {
		return queryForm
	}
	return nil
}

// Get Query String By Key
// GET /path?id=1234&name=Manu&value=
// c.Query("id") == "1234"
// c.Query("name") == "Manu"
// c.Query("value") == ""
// c.Query("wtf") == ""
func (ctx *HttpContext) Query(key string) string {
	return ctx.QueryStrings().Get(key)
}

func (ctx *HttpContext) QueryStringOrDefault(key string, defaultval string) string {
	val := ctx.QueryStrings().Get(key)
	if val == "" {
		val = defaultval
	}
	return val
}

// Redirect redirects the request
func (ctx *HttpContext) Redirect(code int, url string) {
	http.Redirect(ctx.Response, ctx.Request, url, code)
}

// Path returns URL Path string.
func (ctx *HttpContext) Path() string {
	return ctx.Request.URL.Path
}

// Referer returns request referer.
func (ctx *HttpContext) Referer() string {
	return ctx.Request.Header.Get("Referer")
}

// UserAgent returns http request UserAgent
func (ctx *HttpContext) UserAgent() string {
	return ctx.Request.Header.Get("User-Agent")
}

//Get Http Method.
func (ctx *HttpContext) Method() string {
	return ctx.Request.Method
}

//Get Http Status Code.
func (ctx *HttpContext) GetStatus() int {
	return ctx.Response.status
}

// Status sets the HTTP response code.
func (ctx *HttpContext) Status(code int) {
	ctx.Response.SetStatus(code)
}

// FormFile gets file from request.
func (ctx *HttpContext) FormFile(key string) (multipart.File, *multipart.FileHeader, error) {
	return ctx.Request.FormFile(key)
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
	http.Error(ctx.Response, error, code)
}

// Write Byte[] Response.
func (ctx *HttpContext) Write(data []byte) (n int, err error) {
	return ctx.Response.Write(data)
}

func bodyAllowedForStatus(status int) bool {
	switch {
	case status >= 100 && status <= 199:
		return false
	case status == http.StatusNoContent:
		return false
	case status == http.StatusNotModified:
		return false
	}
	return true
}

// ActionResult writes the response headers and calls render.ActionResult to render data.
func (c *HttpContext) Render(code int, r ActionResult.IActionResult) {

	if !bodyAllowedForStatus(code) {
		r.WriteContentType(c.Response)
		c.Response.WriteHeaderNow()
		return
	}

	if err := r.Render(c.Response.ResponseWriter); err != nil {
		panic(err)
	}

	c.Status(code)
}
