package YoyoGo

import (
	"github.com/maxzhang1985/yoyogo/Context"
	"github.com/maxzhang1985/yoyogo/Controller"
	"github.com/maxzhang1985/yoyogo/DependencyInjection"
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
	routerBuilder Router.IRouterBuilder
	middleware    middleware
	handlers      []Handler
	Profile       string
	mvcConfigures []func(builder *Controller.ControllerBuilder)
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
		UseEndpoints(routerConfig)
}

// create new application builder
func NewApplicationBuilder() *ApplicationBuilder {
	routerBuilder := Router.NewRouterBuilder()
	recovery := Middleware.NewRecovery()
	logger := Middleware.NewLogger()
	router := Middleware.NewRouter(routerBuilder)
	self := New(logger, recovery, router)
	self.routerBuilder = routerBuilder
	return self
}

// after create builder , apply router and logger and recovery middleware
func (self *ApplicationBuilder) UseMvc() *ApplicationBuilder {
	self.routerBuilder.(*Router.DefaultRouterBuilder).SetMvc(true)
	return self
}

func (this *ApplicationBuilder) SetHostBuildContext(context *HostBuildContext) {
	this.hostContext = context
}

func (this *ApplicationBuilder) ConfigureMvcParts(configure func(builder *Controller.ControllerBuilder)) *ApplicationBuilder {
	this.mvcConfigures = append(this.mvcConfigures, configure)
	return this
}

func (this *ApplicationBuilder) buildMvc(services *DependencyInjection.ServiceCollection) {
	if this.routerBuilder.IsMvc() {
		controllerBuilder := Controller.NewControllerBuilder(services)
		for _, configure := range this.mvcConfigures {
			configure(controllerBuilder)
		}
	}
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
}

// build and combo all middleware to request delegate (ServeHTTP(w http.ResponseWriter, r *http.Request))
func (this *ApplicationBuilder) Build() IRequestDelegate {
	if this.hostContext == nil {
		panic("hostContext is nil! please set.")
	}

	this.hostContext.hostingEnvironment.Profile = this.Profile
	this.middleware = build(this.handlers)
	this.buildMvc(this.hostContext.applicationServicesDef)
	return this
}

func (app *ApplicationBuilder) SetEnvironment(mode string) {
	app.Profile = mode
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
