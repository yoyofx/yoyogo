package context

import (
	"errors"
	"github.com/yoyofx/yoyogo/utils/cast"
	"github.com/yoyofx/yoyogo/web/actionresult"
	"github.com/yoyofx/yoyogo/web/binding"
	"github.com/yoyofxteam/dependencyinjection"
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
	ctx.store = make(map[string]interface{})
	ctx.storeMutex.Unlock()
	binding.SetRequestMaxMemory(maxRequestSizeMemory)
}

//SetItem Set data in context.
func (ctx *HttpContext) SetItem(key string, val interface{}) {
	ctx.storeMutex.Lock()
	if ctx.store == nil {
		ctx.store = make(map[string]interface{})
	}
	ctx.store[key] = val
	ctx.storeMutex.Unlock()
}

// GetItem Get data in context.
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

// Bind BootStrap Binding
func (ctx *HttpContext) Bind(i interface{}) (err error) {
	req := ctx.Input.Request
	contentType := req.Header.Get(HeaderContentType)
	if req.Body == nil {
		err = errors.New("request body can't be empty")
		return err
	}
	bind := binding.Default(req.Method, contentType)
	err = bind.Bind(req, i)
	return err
}

//BindWithUri is a special bind
func (ctx *HttpContext) BindWithUri(i interface{}) (err error) {
	err = binding.Uri.BindUri(ctx.Input.QueryStrings(), i)
	return err
}

func (ctx *HttpContext) BindWithRouteData(i interface{}) (err error) {
	err = binding.Path.BindUri(ctx.Input.RouterData, i)
	return err
}

// BindWith Use Binding By Name
func (ctx *HttpContext) BindWith(i interface{}, bindEnum binding.Binding) (err error) {
	req := ctx.Input.Request
	switch bindEnum.Name() {
	case binding.JSON.Name():
		err = binding.JSON.Bind(req, i)
	case binding.XML.Name():
		err = binding.XML.Bind(req, i)
	case binding.Query.Name():
		err = binding.Query.Bind(req, i)
	case binding.YAML.Name():
		err = binding.YAML.Bind(req, i)
	case binding.FormMultipart.Name():
		err = binding.FormMultipart.Bind(req, i)
	case binding.ProtoBuf.Name():
		err = binding.ProtoBuf.Bind(req, i)
	case binding.MsgPack.Name():
		err = binding.MsgPack.Bind(req, i)
	case binding.Header.Name():
		err = binding.Header.Bind(req, i)
	default: // case MIMEPOSTForm:
		return binding.Form.Bind(req, i)
	}
	return err
}

// Query2Number Query String to number with default value
func Query2Number[N cast.Number](ctx *HttpContext, key string, defaultVal string) N {
	str := ctx.Input.QueryDefault(key, defaultVal)
	num, _ := cast.Str2Number[N](str)
	return num
}

// Redirect redirects the request
func (ctx *HttpContext) Redirect(code int, url string) {
	http.Redirect(ctx.Output.GetWriter(), ctx.Input.GetReader(), url, code)
}

// Render actionresult writes the response headers and calls render.actionresult to render data.
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
