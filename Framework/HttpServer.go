package YoyoGo

import (
	"golang.org/x/net/context"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type HttpServer struct {
	IsTLS                   bool
	Addr, CertFile, KeyFile string
}

func DefaultHttpServer(addr string) HttpServer {
	return HttpServer{IsTLS: false, Addr: addr}
}

func DefaultHttps(addr string, cert string, key string) HttpServer {
	return HttpServer{IsTLS: true, Addr: addr, CertFile: cert, KeyFile: key}
}

func (server HttpServer) GetAddr() string {
	return server.Addr
}

func (server HttpServer) Run(delegate IRequestDelegate) (e error) {

	webserver := &http.Server{
		Addr:    server.Addr,
		Handler: delegate,
	}

	// 创建系统信号接收器
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	go func() {
		<-quit

		if err := webserver.Shutdown(context.Background()); err != nil {
			log.Fatal("Shutdown server:", err)
		}
	}()

	log.Println("Starting HTTP server...")

	if server.IsTLS {
		e = webserver.ListenAndServeTLS(server.CertFile, server.KeyFile)
	} else {
		e = webserver.ListenAndServe()
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
