package YoyoGo

import "net/http"

type HttpServer struct {
	IsTLS                   bool
	Addr, CertFile, KeyFile string
}

func DefaultHttpServer(addr string) HttpServer {
	return HttpServer{IsTLS: false, Addr: addr}
}

func (server HttpServer) Run(delegate IRequestDelegate) (e error) {
	if server.IsTLS {
		e = http.ListenAndServeTLS(server.Addr, server.CertFile, server.KeyFile, delegate)
	} else {
		e = http.ListenAndServe(server.Addr, delegate)
	}

	if e != nil {
		panic(e)
	}

	return nil
}
