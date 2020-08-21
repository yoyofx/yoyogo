package YoyoGo

import (
	"github.com/yoyofx/yoyogo/Abstractions"
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofx/yoyogo/WebFramework/Middleware"
	"github.com/yoyofx/yoyogo/WebFramework/Mvc"
	"github.com/yoyofx/yoyogo/WebFramework/Router"
	"github.com/yoyofx/yoyogo/WebFramework/View"
	"net/http"
)

// HTTP methods

//application builder struct
type WebApplicationBuilder struct {
	hostContext     *Abstractions.HostBuildContext // host build 's context
	routerBuilder   Router.IRouterBuilder          // route builder of interface
	middleware      middleware
	handlers        []MiddlewareHandler
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
		UseServer(DefaultHttpServer(DefaultAddress)).
		Configure(func(app *WebApplicationBuilder) {
			app.UseStatic("/Static", "./Static")
			app.UseEndpoints(routerConfig)
		})
}

func CreateBlankWebBuilder() *WebHostBuilder {
	return NewWebHostBuilder()
}

// create application builder when combo all handlers to middleware
func New(handlers ...MiddlewareHandler) *WebApplicationBuilder {
	return &WebApplicationBuilder{
		handlers: handlers,
	}
}

// create new web application builder
func NewWebApplicationBuilder() *WebApplicationBuilder {
	routerBuilder := Router.NewRouterBuilder()
	recovery := Middleware.NewRecovery()
	logger := Middleware.NewLogger()
	router := Middleware.NewRouter(routerBuilder)
	jwt := Middleware.NewJwt()
	self := New(logger, recovery, jwt, router)
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
	this.routerBuilder.SetConfiguration(this.hostContext.Configuration)
	for _, configure := range this.routeConfigures {
		configure(this.routerBuilder)
	}
}

func (this *WebApplicationBuilder) buildMvc() {
	if this.routerBuilder.IsMvc() {
		controllerBuilder := this.routerBuilder.GetMvcBuilder()
		// add config for controller builder
		controllerBuilder.SetConfiguration(this.hostContext.Configuration)
		for _, configure := range this.mvcConfigures {
			configure(controllerBuilder)
		}
		// add view engine
		var viewEngine View.IViewEngine
		err := this.hostContext.HostServices.GetServiceByName(&viewEngine, "viewEngine")
		if err == nil {
			option := this.routerBuilder.GetMvcBuilder().GetRouterHandler().Options.ViewOption
			viewEngine.SetTemplatePath(option)
			controllerBuilder.SetViewEngine(viewEngine)
		}

		// add controllers to application services
		controllerDescriptorList := controllerBuilder.GetControllerDescriptorList()
		for _, descriptor := range controllerDescriptorList {
			this.hostContext.
				ApplicationServicesDef.
				AddSingletonByNameAndImplements(descriptor.ControllerName, descriptor.ControllerType, new(Mvc.IController))
		}
	}
}

func (this *WebApplicationBuilder) buildMiddleware() {
	for _, handler := range this.handlers {
		if configurationMdw, ok := handler.(Middleware.IConfigurationMiddleware); ok {
			configurationMdw.SetConfiguration(this.hostContext.Configuration)
		}
	}
	this.middleware = build(this.handlers)
}

//  this time is not build host.Context.HostServices , that add services define
func (this *WebApplicationBuilder) innerConfigures() {
	this.hostContext.
		ApplicationServicesDef.
		AddSingletonByNameAndImplements("viewEngine", View.CreateViewEngine, new(View.IViewEngine))
	//-------------------------  view engine ----------------------------------

}

// build and combo all middleware to request delegate (ServeHTTP(w http.ResponseWriter, r *http.Request))
// return Abstractions.IRequestDelegate type
func (this *WebApplicationBuilder) Build() interface{} {
	if this.hostContext == nil {
		panic("hostContext is nil! please set.")
	}

	this.buildMiddleware()
	this.buildEndPoints()
	this.buildMvc()
	return this
}

func (this *WebApplicationBuilder) SetHostBuildContext(context *Abstractions.HostBuildContext) {
	this.hostContext = context
	// has host.Context.HostServices
	if this.hostContext.ApplicationServicesDef != nil {
		this.innerConfigures()
	}
}

// apply middleware in builder
func (app *WebApplicationBuilder) UseMiddleware(handler MiddlewareHandler) {
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
Middleware of Server MiddlewareHandler , request port.
*/
func (app *WebApplicationBuilder) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	app.middleware.Invoke(Context.NewContext(w, r, app.hostContext.ApplicationServices))
}
