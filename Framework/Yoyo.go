package YoyoGo

import (
	"encoding/base64"
	"fmt"
	"github.com/maxzhang1985/yoyogo/Middleware"
	"log"
	"net/http"
	"os"
)

// HTTP methods
const (
	CONNECT = http.MethodConnect
	DELETE  = http.MethodDelete
	GET     = http.MethodGet
	HEAD    = http.MethodHead
	OPTIONS = http.MethodOptions
	PATCH   = http.MethodPatch
	POST    = http.MethodPost
	PUT     = http.MethodPut
	TRACE   = http.MethodTrace
)

var methods = [...]string{
	http.MethodConnect,
	http.MethodDelete,
	http.MethodGet,
	http.MethodHead,
	http.MethodOptions,
	http.MethodPatch,
	http.MethodPost,
	http.MethodPut,
	http.MethodTrace,
}

const (
	// DefaultAddress is used if no other is specified.
	DefaultAddress = ":8080"
)

type YoyoGo struct {
	router     *Middleware.RouterMiddleware
	middleware middleware
	handlers   []Handler
}

func UseClassic() *YoyoGo {
	return &YoyoGo{}
}

func UseMvc() *YoyoGo {
	logger := Middleware.NewLogger()
	router := Middleware.NewRouter()
	self := New(logger, router)
	self.router = router
	return self
}

func New(handlers ...Handler) *YoyoGo {
	return &YoyoGo{
		handlers:   handlers,
		middleware: build(handlers),
	}
}

func (n *YoyoGo) Use(handler Handler) {
	if handler == nil {
		panic("handler cannot be nil")
	}

	n.handlers = append(n.handlers, handler)
	n.middleware = build(n.handlers)
}
func (app *YoyoGo) UseStatic(path string) {
	app.Use(Middleware.NewStatic("Static"))

}

// UseFunc adds a Negroni-style handler function onto the middleware stack.

// UseHandler adds a http.Handler onto the middleware stack. Handlers are invoked in the order they are added to a Negroni.

//func (yoyo *YoyoGo) Map(relativePath string, handler func(ctx *Middleware.HttpContext)) {
//	yoyo.router.ReqFuncMap[relativePath] = handler
//}

func (yoyo *YoyoGo) printLogo() {
	logo, _ := base64.StdEncoding.DecodeString("CiBfICAgICBfICAgICAgICAgICAgICAgICAgICBfX18gICAgICAgICAgCiggKSAgICggKSAgICAgICAgICAgICAgICAgICggIF9gXCAgICAgICAgCmBcYFxfLycvJ18gICAgXyAgIF8gICAgXyAgIHwgKCAoXykgICBfICAgCiAgYFwgLycvJ19gXCAoICkgKCApIC8nX2BcIHwgfF9fXyAgLydfYFwgCiAgIHwgfCggKF8pICl8IChfKSB8KCAoXykgKXwgKF8sICkoIChfKSApCiAgIChfKWBcX19fLydgXF9fLCB8YFxfX18vJyhfX19fLydgXF9fXy8nCiAgICAgICAgICAgICAoIClffCB8ICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICBgXF9fXy8nICAgICAgICAgICAgICAgICAgICAgCg==")
	fmt.Println(string(logo))
}

func (n *YoyoGo) UseHandler(handler http.Handler) {
	n.Use(wrap(handler))
}

func (n *YoyoGo) UseHandlerFunc(handlerFunc func(rw http.ResponseWriter, r *http.Request)) {
	n.Use(wrapFunc(handlerFunc))
}

func (n *YoyoGo) UseFunc(handlerFunc HandlerFunc) {
	n.Use(handlerFunc)
}

func (yoyo *YoyoGo) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	yoyo.middleware.Invoke(Middleware.NewContext(w, r))
	//fmt.Println(r.URL.Path)
	//ctx := NewContext(w,r)
	//fun,ok := reqFuncMap[r.URL.Path]
	//if ok{
	//	fun(ctx)
	//	return
	//}
}

func (yoyo *YoyoGo) Run(addr ...string) {
	yoyo.printLogo()
	l := log.New(os.Stdout, "[yoyogo] ", 0)
	finalAddr := detectAddress(addr...)
	l.Printf("listening on %s", finalAddr)
	l.Fatal(http.ListenAndServe(finalAddr, yoyo))
}

func detectAddress(addr ...string) string {
	if len(addr) > 0 {
		return addr[0]
	}
	if port := os.Getenv("PORT"); port != "" {
		return ":" + port
	}
	return DefaultAddress
}

func (yoyo *YoyoGo) Map(method string, path string, handler func(ctx *Middleware.HttpContext)) {
	if len(path) < 1 || path[0] != '/' {
		panic("Path should be like '/yoyo/go'")
	}
	yoyo.router.Tree.Insert(method, path, handler)
}

// GET register GET request handler
func (yoyo *YoyoGo) GET(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(GET, path, handler)
}

// HEAD register HEAD request handler
func (yoyo *YoyoGo) HEAD(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(HEAD, path, handler)
}

// OPTIONS register OPTIONS request handler
func (yoyo *YoyoGo) OPTIONS(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(OPTIONS, path, handler)
}

// POST register POST request handler
func (yoyo *YoyoGo) POST(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(POST, path, handler)
}

// PUT register PUT request handler
func (yoyo *YoyoGo) PUT(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(PUT, path, handler)
}

// PATCH register PATCH request HandlerFunc
func (yoyo *YoyoGo) PATCH(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(PATCH, path, handler)
}

// DELETE register DELETE request handler
func (yoyo *YoyoGo) DELETE(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(DELETE, path, handler)
}

// CONNECT register CONNECT request handler
func (yoyo *YoyoGo) CONNECT(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(CONNECT, path, handler)
}

// TRACE register TRACE request handler
func (yoyo *YoyoGo) TRACE(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(TRACE, path, handler)
}

// Any register any method handler
func (yoyo *YoyoGo) Any(path string, handler func(ctx *Middleware.HttpContext)) {
	for _, m := range methods {
		yoyo.Map(m, path, handler)
	}
}
