package YoyoGo

import (
	"github.com/lucas-clemente/quic-go/http3"
)

type HttpQUICServer struct {
	Addr, CertFile, KeyFile string
}

func (server HttpQUICServer) GetAddr() string {
	return server.Addr
}

func (server HttpQUICServer) Run(delegate IRequestDelegate) (e error) {

	e = http3.ListenAndServeQUIC(server.Addr, server.CertFile, server.KeyFile, delegate)

	if e != nil {
		panic(e)
	}

	return e
}
