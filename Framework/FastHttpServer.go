package YoyoGo

import (
	"github.com/valyala/fasthttp"
	"github.com/valyala/fasthttp/fasthttpadaptor"
)

type FastHttpServer struct {
	Addr, CertFile, KeyFile string
}

func NewFastHttpServer(addr string, certfile string, keyfile string) FastHttpServer {
	return FastHttpServer{Addr: addr, CertFile: certfile, KeyFile: keyfile}
}

func (server FastHttpServer) GetAddr() string {
	return server.Addr
}

func (server FastHttpServer) Run(delegate IRequestDelegate) (e error) {

	fastHttpHandler := fasthttpadaptor.NewFastHTTPHandler(delegate)
	e = fasthttp.ListenAndServe(server.Addr, fastHttpHandler)

	if e != nil {
		panic(e)
	}

	return e
}
