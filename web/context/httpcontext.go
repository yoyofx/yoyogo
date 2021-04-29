package context

import (
	"errors"
	"github.com/yoyofx/yoyogo/dependencyinjection"
	"github.com/yoyofx/yoyogo/web/actionresult"
	"github.com/yoyofx/yoyogo/web/binding"
	"net/http"
	"sync"
)

const (
	defaultTagName = "param"
	jsonTagName    = "json"
)

var (
	defaultMaxMemory int64 = 32 << 20 // 32 MB
)

type H = map[string]interface{}

type HttpContext struct {
	Input            *Input
	Output           Output
	RequiredServices dependencyinjection.IServiceProvider
	store            map[string]interface{}
	storeMutex       *sync.RWMutex
	Result           interface{}
}

func NewContext(w http.ResponseWriter, r *http.Request, maxRequestSizeMemory int64, sp dependencyinjection.IServiceProvider) *HttpContext {
	if maxRequestSizeMemory <= defaultMaxMemory {
		maxRequestSizeMemory = defaultMaxMemory
	}
	ctx := &HttpContext{}
	ctx.init(w, r, maxRequestSizeMemory, sp)
	return ctx
}

func (ctx *HttpContext) init(w http.ResponseWriter, r *http.Request, maxRequestSizeMemory int64, sp dependencyinjection.IServiceProvider) {
	ctx.storeMutex = new(sync.RWMutex)
	ctx.Input = NewInput(r, maxRequestSizeMemory)
	ctx.Output = Output{Response: &CResponseWriter{w, 0, 0, nil}}
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

// Get JWT UserInfo
func (ctx *HttpContext) GetUser() map[string]interface{} {
	v := ctx.GetItem("userinfo")
	if v != nil {
		return v.(map[string]interface{})
	}
	return nil
}

//BootStrap Binding
func (ctx *HttpContext) Bind(i interface{}) (err error) {
	req := ctx.Input.Request
	contentType := req.Header.Get(HeaderContentType)
	if req.Body == nil {
		err = errors.New("request body can't be empty")
		return err
	}
	err = errors.New("request unsupported MediaType -> " + contentType)
	bind := binding.Default(req.Method, contentType)
	bind.Bind(req, i)
	return err
}

//Use Binding By Name
func (ctx *HttpContext) AppointBinding(i interface{}, bindEnum binding.Binding) (err error) {
	req := ctx.Input.Request
	contentType := req.Header.Get(HeaderContentType)
	err = errors.New("request unsupported MediaType -> " + contentType)
	switch bindEnum.Name() {
	case binding.JSON.Name():
		err = binding.JSON.Bind(req, i)
	case binding.XML.Name():
		err = binding.XML.Bind(req, i)
	case binding.Query.Name():
		err = binding.Query.Bind(req, i)
	case binding.Uri.Name():
		err = binding.Uri.BindUri(ctx.Input.GetAllParam(), i)
	case binding.YAML.Name():
		err = binding.YAML.Bind(req, i)
	case binding.FormMultipart.Name():
		err = binding.FormMultipart.Bind(req, i)
	}
	return err
}

// Redirect redirects the request
func (ctx *HttpContext) Redirect(code int, url string) {
	http.Redirect(ctx.Output.GetWriter(), ctx.Input.GetReader(), url, code)
}

// actionresult writes the response headers and calls render.actionresult to render data.
func (ctx *HttpContext) Render(code int, r actionresult.IActionResult) {

	if !bodyAllowedForStatus(code) {
		r.WriteContentType(ctx.Output.GetWriter())
		ctx.Output.SetStatusCodeNow()
		return
	}

	if err := r.Render(ctx.Output.GetWriter()); err != nil {
		panic(err)
	}

	ctx.Output.SetStatusCode(code)
}

func (ctx *HttpContext) GetSession() *Session {
	sessionId := ctx.GetItem("sessionId").(string)
	if sessionId == "" {
		return nil
	}
	mgr := ctx.GetItem("sessionMgr").(ISessionManager)
	return NewSession(sessionId, mgr)
}
