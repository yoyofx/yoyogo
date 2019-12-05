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

//application builder struct
type ApplicationBuilder struct {
	hostContext   *HostBuildContext
	routerHandler Router.IRouterHandler
	middleware    middleware
	handlers      []Handler
}

// create classic application builder
func UseClassic() *ApplicationBuilder {
	return &ApplicationBuilder{}
}

//region Create the builder of Web host
func CreateDefaultBuilder(routerConfig func(router Router.IRouterBuilder)) *HostBuilder {
	return NewWebHostBuilder().
		UseServer(DefaultHttpServer(DefaultAddress)).
		Configure(func(app *ApplicationBuilder) {
			app.UseStatic("Static")
		}).
		UseRouter(routerConfig)
}

// create new application builder
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

// after create builder , apply router and logger and recovery middleware
func (self *ApplicationBuilder) UseMvc() *ApplicationBuilder {
	self.routerHandler = Router.NewRouterHandler()
	self.UseMiddleware(Middleware.NewLogger())
	self.UseMiddleware(Middleware.NewRecovery())
	self.UseMiddleware(Middleware.NewRouter(self.routerHandler))

	return self
}

// create application builder when combo all handlers to middleware
func New(handlers ...Handler) *ApplicationBuilder {
	return &ApplicationBuilder{
		handlers:   handlers,
		middleware: build(handlers),
	}
}

// apply middleware in builder
func (app *ApplicationBuilder) UseMiddleware(handler Handler) {
	if handler == nil {
		panic("handler cannot be nil")
	}

	app.handlers = append(app.handlers, handler)
	//n.middleware = build(n.handlers)
}

// build and combo all middleware to request delegate (ServeHTTP(w http.ResponseWriter, r *http.Request))
func (app *ApplicationBuilder) Build() IRequestDelegate {
	app.middleware = build(app.handlers)
	return app
}

func (app *ApplicationBuilder) SetEnvironment(mode string) {
	app.hostContext.hostingEnvironment.Profile = mode
}

// apply static middleware in builder
func (app *ApplicationBuilder) UseStatic(path string) {
	app.UseMiddleware(Middleware.NewStatic("Static"))
}

// apply handler middleware in builder
func (app *ApplicationBuilder) UseHandler(handler http.Handler) {
	app.UseMiddleware(wrap(handler))
}

// apply handler func middleware in builder
func (app *ApplicationBuilder) UseHandlerFunc(handlerFunc func(rw http.ResponseWriter, r *http.Request)) {
	app.UseMiddleware(wrapFunc(handlerFunc))
}

// apply handler func middleware in builder
func (app *ApplicationBuilder) UseFunc(handlerFunc HandlerFunc) {
	app.UseMiddleware(handlerFunc)
}

/*
Middleware of Server Handler , request port.
*/
func (app *ApplicationBuilder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.middleware.Invoke(Context.NewContext(w, r, app.hostContext.applicationServices))
}
