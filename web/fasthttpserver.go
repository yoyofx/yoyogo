package web

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"github.com/yoyofx/yoyogo/abstractions"
	"log"
	"net/http"
)

type FastHttpServer struct {
	IsTLS                   bool
	Addr, CertFile, KeyFile string
	webserver               *fasthttp.Server
}

func NewFastHttp(addr string) *FastHttpServer {
	return &FastHttpServer{IsTLS: false, Addr: addr}
}

func NewFastHttps(addr string, cert string, key string) *FastHttpServer {
	return &FastHttpServer{IsTLS: true, Addr: addr, CertFile: cert, KeyFile: key}
}

func NewFastHttpByConfig(config ServerConfig) *FastHttpServer {
	if config.IsTLS {
		return NewFastHttps(config.Addr, config.CertFile, config.KeyFile)
	} else {
		return NewFastHttp(config.Addr)
	}
}

func (server *FastHttpServer) GetAddr() string {
	return server.Addr
}

func (server *FastHttpServer) Run(context *abstractions.HostBuilderContext) (e error) {

	fastHttpHandler := fasthttpadaptor.NewFastHTTPHandler(context.RequestDelegate.(IRequestDelegate))

	server.webserver = &fasthttp.Server{
		Handler: fastHttpHandler,
	}

	addr := server.Addr
	if server.Addr == "" {
		addr = context.HostingEnvironment.Addr
	}
	context.ApplicationCycle.StartApplication()
	if server.IsTLS {
		e = server.webserver.ListenAndServeTLS(addr, server.CertFile, server.KeyFile)
	} else {
		e = server.webserver.ListenAndServe(addr)
	}
	if e != nil {
		if e == http.ErrServerClosed {
			log.Println("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected", e)
		}
	}

	return e
}

func (server *FastHttpServer) Shutdown() {
	log.Println("Shutdown HTTP server...")
	if err := server.webserver.Shutdown(); err != nil {
		log.Fatal("Shutdown server:", err)
	}
}
