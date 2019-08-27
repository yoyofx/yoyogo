package YoyoGo

import (
	"log"
	"net/http"
	"os"
)

type M = map[string]string

const (
	// DefaultAddress is used if no other is specified.
	DefaultAddress = ":8080"
)

type YoyoGo struct {
	middleware middleware
	handlers   []Handler
}

func Classic() *YoyoGo {
	return &YoyoGo{}
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

// UseFunc adds a Negroni-style handler function onto the middleware stack.

// UseHandler adds a http.Handler onto the middleware stack. Handlers are invoked in the order they are added to a Negroni.

var reqFuncMap = make(map[string]func(ctx *HttpContext))

func (yoyo *YoyoGo) Map(relativePath string, handler func(ctx *HttpContext)) {
	reqFuncMap[relativePath] = handler
}

func (yoyo *YoyoGo) Run(addr ...string) {
	l := log.New(os.Stdout, "[yoyofx] ", 0)
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

	yoyo.middleware.Invoke(NewContext(w, r))
	//fmt.Println(r.URL.Path)
	//ctx := NewContext(w,r)
	//fun,ok := reqFuncMap[r.URL.Path]
	//if ok{
	//	fun(ctx)
	//	return
	//}
}
