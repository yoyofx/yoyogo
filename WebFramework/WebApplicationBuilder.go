package YoyoGo

import (
	"github.com/yoyofx/yoyogo"
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/DependencyInjection"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Middleware"
	"github.com/yoyofx/yoyogo/WebFramework/Mvc"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
	"net/http"
)

// HTTP methods

//application builder struct
type WebApplicationBuilder struct {
	hostContext     *Abstractions.HostBuildContext // host build 's context
	routerBuilder   Router.IRouterBuilder          // route builder of interface
	middleware      middleware
	handlers        []Handler
	routeConfigures []func(Router.IRouterBuilder)          // endpoints router configure functions
	mvcConfigures   []func(builder *Mvc.ControllerBuilder) // mvc router configure functions
}

// create classic application builder
func UseClassic() *WebApplicationBuilder {
	return &WebApplicationBuilder{}
}

//region Create the builder of Web host
func CreateDefaultBuilder(routerConfig func(router Router.IRouterBuilder)) *Abstractions.HostBuilder {
	return NewWebHostBuilder().
		UseServer(DefaultHttpServer(YoyoGo.DefaultAddress)).
		Configure(func(app *WebApplicationBuilder) {
			app.UseStatic("/Static", "./Static")
			app.UseEndpoints(routerConfig)
		})
}

func CreateBlankWebBuilder() *WebHostBuilder {
	return NewWebHostBuilder()
}

// create application builder when combo all handlers to middleware
func New(handlers ...Handler) *WebApplicationBuilder {
	return &WebApplicationBuilder{
		handlers: handlers,
		//middleware: build(handlers),
	}
}

// create new web application builder
func NewWebApplicationBuilder() *WebApplicationBuilder {
	routerBuilder := Router.NewRouterBuilder()
	recovery := Middleware.NewRecovery()
	logger := Middleware.NewLogger()
	router := Middleware.NewRouter(routerBuilder)
	self := New(logger, recovery, router)
	self.routerBuilder = routerBuilder
	return self
}

// UseMvc after create builder , apply router and logger and recovery middleware
func (self *WebApplicationBuilder) UseMvc(configure func(builder *Mvc.ControllerBuilder)) *WebApplicationBuilder {
	if !self.routerBuilder.IsMvc() {
		self.routerBuilder.UseMvc(true)
	}
	self.mvcConfigures = append(self.mvcConfigures, configure)
	return self
}

func (self *WebApplicationBuilder) UseEndpoints(configure func(Router.IRouterBuilder)) *WebApplicationBuilder {
	self.routeConfigures = append(self.routeConfigures, configure)
	return self
}

func (this *WebApplicationBuilder) buildEndPoints() {
	for _, configure := range this.routeConfigures {
		configure(this.routerBuilder)
	}
}

func (this *WebApplicationBuilder) buildMvc(services *DependencyInjection.ServiceCollection) {
	if this.routerBuilder.IsMvc() {
		controllerBuilder := this.routerBuilder.GetMvcBuilder()
		for _, configure := range this.mvcConfigures {
			configure(controllerBuilder)
		}
		// add controllers to application services
		controllerDescriptorList := controllerBuilder.GetControllerDescriptorList()
		for _, descriptor := range controllerDescriptorList {
			services.AddSingletonByNameAndImplements(descriptor.ControllerName, descriptor.ControllerType, new(Mvc.IController))
		}
	}
}

// build and combo all middleware to request delegate (ServeHTTP(w http.ResponseWriter, r *http.Request))
// return Abstractions.IRequestDelegate type
func (this *WebApplicationBuilder) Build() interface{} {
	if this.hostContext == nil {
		panic("hostContext is nil! please set.")
	}
	//this.hostContext.HostingEnvironment
	this.middleware = build(this.handlers)
	this.buildEndPoints()
	this.buildMvc(this.hostContext.ApplicationServicesDef)
	return this
}

func (this *WebApplicationBuilder) SetHostBuildContext(context *Abstractions.HostBuildContext) {
	this.hostContext = context
}

// apply middleware in builder
func (app *WebApplicationBuilder) UseMiddleware(handler Handler) {
	if handler == nil {
		panic("handler cannot be nil")
	}
	app.handlers = append(app.handlers, handler)
}

// apply static middleware in builder
func (app *WebApplicationBuilder) UseStatic(patten string, path string) {
	app.UseMiddleware(Middleware.NewStatic(patten, path))
}

func (app *WebApplicationBuilder) UseStaticAssets() {
	app.UseMiddleware(Middleware.NewStaticWithConfig(app.hostContext.Configuration))
}

// apply handler middleware in builder
func (app *WebApplicationBuilder) UseHandler(handler http.Handler) {
	app.UseMiddleware(wrap(handler))
}

// apply handler func middleware in builder
func (app *WebApplicationBuilder) UseHandlerFunc(handlerFunc func(rw http.ResponseWriter, r *http.Request)) {
	app.UseMiddleware(wrapFunc(handlerFunc))
}

// apply handler func middleware in builder
func (app *WebApplicationBuilder) UseFunc(handlerFunc MiddlewareHandlerFunc) {
	app.UseMiddleware(handlerFunc)
}

/*
Middleware of Server Handler , request port.
*/
func (app *WebApplicationBuilder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.middleware.Invoke(Context.NewContext(w, r, app.hostContext.ApplicationServices))
}
