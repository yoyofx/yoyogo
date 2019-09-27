package YoyoGo

import (
	"encoding/base64"
	"fmt"
	"github.com/maxzhang1985/yoyogo/Middleware"
	"github.com/maxzhang1985/yoyogo/Standard"
	"log"
	"net/http"
	"os"
)

// HTTP methods

const (
	// DefaultAddress is used if no other is specified.
	DefaultAddress = ":8080"
)

type YoyoGo struct {
	Mode       string
	router     *Middleware.RouterMiddleware
	Recovery   *Middleware.Recovery
	middleware middleware
	handlers   []Handler
}

func UseClassic() *YoyoGo {
	return &YoyoGo{Mode: Dev}
}

func UseMvc() *YoyoGo {
	recovery := Middleware.NewRecovery()
	logger := Middleware.NewLogger()
	router := Middleware.NewRouter()
	self := New(logger, recovery, router)
	self.router = router
	self.Recovery = recovery
	return self
}

func New(handlers ...Handler) *YoyoGo {
	return &YoyoGo{
		Mode:       Dev,
		handlers:   handlers,
		middleware: build(handlers),
	}
}
func (app *YoyoGo) SetMode(mode string) {
	app.Mode = mode
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

func (yoyo *YoyoGo) printLogo(l *log.Logger, port string) {
	logo, _ := base64.StdEncoding.DecodeString("CiBfICAgICBfICAgICAgICAgICAgICAgICAgICBfX18gICAgICAgICAgCiggKSAgICggKSAgICAgICAgICAgICAgICAgICggIF9gXCAgICAgICAgCmBcYFxfLycvJ18gICAgXyAgIF8gICAgXyAgIHwgKCAoXykgICBfICAgCiAgYFwgLycvJ19gXCAoICkgKCApIC8nX2BcIHwgfF9fXyAgLydfYFwgCiAgIHwgfCggKF8pICl8IChfKSB8KCAoXykgKXwgKF8sICkoIChfKSApCiAgIChfKWBcX19fLydgXF9fLCB8YFxfX18vJyhfX19fLydgXF9fXy8nCiAgICAgICAgICAgICAoIClffCB8ICAgICAgICAgICAgICAgICAgICAgCiAgICAgICAgICAgICBgXF9fXy8nICAgICAgICAgICAgICAgICAgICAgCg==")
	fmt.Println(string(logo))

	l.Printf("listening on %s", port)
	l.Printf("application is runing pid: %d", os.Getpid())
	l.Printf("runing in %s mode , switch on 'Prod' mode in production.", yoyo.Mode)
	l.Println(" - use Prod app.SetMode(Prod) ")
}

func (yoyo *YoyoGo) Run(addr ...string) {
	finalAddr := detectAddress(addr...)
	l := log.New(os.Stdout, "[yoyogo] ", 0)
	yoyo.printLogo(l, finalAddr)
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
	yoyo.Map(Std.GET, path, handler)
}

// HEAD register HEAD request handler
func (yoyo *YoyoGo) HEAD(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(Std.HEAD, path, handler)
}

// OPTIONS register OPTIONS request handler
func (yoyo *YoyoGo) OPTIONS(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(Std.OPTIONS, path, handler)
}

// POST register POST request handler
func (yoyo *YoyoGo) POST(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(Std.POST, path, handler)
}

// PUT register PUT request handler
func (yoyo *YoyoGo) PUT(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(Std.PUT, path, handler)
}

// PATCH register PATCH request HandlerFunc
func (yoyo *YoyoGo) PATCH(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(Std.PATCH, path, handler)
}

// DELETE register DELETE request handler
func (yoyo *YoyoGo) DELETE(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(Std.DELETE, path, handler)
}

// CONNECT register CONNECT request handler
func (yoyo *YoyoGo) CONNECT(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(Std.CONNECT, path, handler)
}

// TRACE register TRACE request handler
func (yoyo *YoyoGo) TRACE(path string, handler func(ctx *Middleware.HttpContext)) {
	yoyo.Map(Std.TRACE, path, handler)
}

// Any register any method handler
func (yoyo *YoyoGo) Any(path string, handler func(ctx *Middleware.HttpContext)) {
	for _, m := range Std.Methods {
		yoyo.Map(m, path, handler)
	}
}

func (yoyo *YoyoGo) Group(name string, routerBuilderFunc func(router *Middleware.RouterGroup)) {
	group := &Middleware.RouterGroup{
		Name:       name,
		RouterTree: yoyo.router.Tree,
	}
	if routerBuilderFunc == nil {
		panic("routerBuilderFunc is nil")
	}

	routerBuilderFunc(group)
}
