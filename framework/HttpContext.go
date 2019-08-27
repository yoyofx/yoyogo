package YoyoGo

import (
	"encoding/json"
	"net/http"
	"sync"
)

type HttpContext struct {
	Req          *http.Request
	Resp         *responseWriter
	store        map[string]interface{}
	storeMutex   *sync.RWMutex
}

func NewContext(w http.ResponseWriter, r *http.Request) *HttpContext {
	ctx := &HttpContext{}
	ctx.Init(w, r)
	return ctx
}

func (ctx *HttpContext) Init(w http.ResponseWriter, r *http.Request) {
	ctx.storeMutex = new(sync.RWMutex)
	ctx.Resp = &responseWriter{w, 0}
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

func (ctx *HttpContext) JSON(data interface{})  {
	ctx.Resp.Header().Set("Content-Type", "application/json")
	jsons, _ := json.Marshal(data)
	_, _ = ctx.Resp.Write(jsons)
}
