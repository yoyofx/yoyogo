package YoyoGo

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type FastHttpServer struct {
	IsTLS                   bool
	Addr, CertFile, KeyFile string
}

func NewFastHttp(addr string) FastHttpServer {
	return FastHttpServer{IsTLS: false, Addr: addr}
}

func NewFastHttps(addr string, cert string, key string) FastHttpServer {
	return FastHttpServer{IsTLS: true, Addr: addr, CertFile: cert, KeyFile: key}
}

func (server FastHttpServer) GetAddr() string {
	return server.Addr
}

func (server FastHttpServer) Run(delegate IRequestDelegate) (e error) {

	fastHttpHandler := fasthttpadaptor.NewFastHTTPHandler(delegate)
	e = fasthttp.ListenAndServe(server.Addr, fastHttpHandler)
	webserver := &fasthttp.Server{
		Handler: fastHttpHandler,
	}

	// 创建系统信号接收器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit

		if err := webserver.Shutdown(); err != nil {
			log.Fatal("Shutdown server:", err)
		}
	}()

	log.Println("Starting HTTP server...")

	if server.IsTLS {
		e = webserver.ListenAndServeTLS(server.Addr, server.CertFile, server.KeyFile)
	} else {
		e = webserver.ListenAndServe(server.Addr)
	}
	if e != nil {
		if e == http.ErrServerClosed {
			log.Print("Server closed under request")
		} else {
			log.Fatal("Server closed unexpected")
		}
	}

	return e
}
