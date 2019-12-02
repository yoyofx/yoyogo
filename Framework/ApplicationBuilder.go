package YoyoGo

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/Middleware"
	"github.com/maxzhang1985/yoyogo/Router"
	"net/http"
)

// HTTP methods

const (
	// DefaultAddress is used if no other is specified.
	DefaultAddress = ":8080"
)

type ApplicationBuilder struct {
	hostContext   *HostBuildContext
	routerHandler Router.IRouterHandler
	middleware    middleware
	handlers      []Handler
}

func UseClassic() *ApplicationBuilder {
	return &ApplicationBuilder{}
}

//region Create the builder of Web host
func CreateDefaultWebHostBuilder(args []string, routerConfig func(router Router.IRouterBuilder)) *HostBuilder {
	return NewWebHostBuilder().
		UseServer(DefaultHttpServer(DefaultAddress)).
		Configure(func(app *ApplicationBuilder) {
			app.UseStatic("Static")
		}).
		UseRouter(routerConfig)
}

func NewApplicationBuilder(context *HostBuildContext) *ApplicationBuilder {
	routerHandler := Router.NewRouterHandler()
	recovery := Middleware.NewRecovery()
	logger := Middleware.NewLogger()
	router := Middleware.NewRouter(routerHandler)
	self := New(logger, recovery, router)
	self.routerHandler = routerHandler
	self.hostContext = context
	return self
}

func (self *ApplicationBuilder) UseMvc() *ApplicationBuilder {
	self.routerHandler = Router.NewRouterHandler()
	self.UseMiddleware(Middleware.NewLogger())
	self.UseMiddleware(Middleware.NewRecovery())
	self.UseMiddleware(Middleware.NewRouter(self.routerHandler))

	return self
}

func New(handlers ...Handler) *ApplicationBuilder {
	return &ApplicationBuilder{
		handlers:   handlers,
		middleware: build(handlers),
	}
}

func (n *ApplicationBuilder) UseMiddleware(handler Handler) {
	if handler == nil {
		panic("handler cannot be nil")
	}

	n.handlers = append(n.handlers, handler)
	//n.middleware = build(n.handlers)
}

func (n *ApplicationBuilder) Build() IRequestDelegate {
	n.middleware = build(n.handlers)
	return n
}

func (app *ApplicationBuilder) UseStatic(path string) {
	app.UseMiddleware(Middleware.NewStatic("Static"))
}

func (n *ApplicationBuilder) UseHandler(handler http.Handler) {
	n.UseMiddleware(wrap(handler))
}

func (n *ApplicationBuilder) UseHandlerFunc(handlerFunc func(rw http.ResponseWriter, r *http.Request)) {
	n.UseMiddleware(wrapFunc(handlerFunc))
}

func (n *ApplicationBuilder) UseFunc(handlerFunc HandlerFunc) {
	n.UseMiddleware(handlerFunc)
}

/*
Middleware of Server Handler , request port.
*/

func (yoyo *ApplicationBuilder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	yoyo.middleware.Invoke(Context.NewContext(w, r))
}
