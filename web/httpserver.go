package web

import (
	"github.com/yoyofx/yoyogo/abstractions"
	"golang.org/x/net/context"
	"log"
	"net/http"
)

type HttpServer struct {
	IsTLS                   bool
	Addr, CertFile, KeyFile string
	webserver               *http.Server
}

func DefaultHttpServer(addr string) *HttpServer {
	return &HttpServer{IsTLS: false, Addr: addr}
}

func DefaultHttps(addr string, cert string, key string) *HttpServer {
	return &HttpServer{IsTLS: true, Addr: addr, CertFile: cert, KeyFile: key}
}

func (server *HttpServer) GetAddr() string {
	return server.Addr
}

func (server *HttpServer) Run(context *abstractions.HostBuilderContext) (e error) {
	addr := server.Addr
	if server.Addr == "" {
		addr = context.HostingEnvironment.Addr
	}

	server.webserver = &http.Server{
		Addr:    addr,
		Handler: context.RequestDelegate.(IRequestDelegate),
	}

	context.ApplicationCycle.StartApplication()

	if server.IsTLS {
		e = server.webserver.ListenAndServeTLS(server.CertFile, server.KeyFile)
	} else {
		e = server.webserver.ListenAndServe()
	}
	if e != nil {
		if e == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}

	return nil
}

func (server *HttpServer) Shutdown() {
	if err := server.webserver.Shutdown(context.Background()); err != nil {
		log.Fatal("Shutdown server:", err)
	}
	log.Fatal("Shutdown HTTP server...")
}
