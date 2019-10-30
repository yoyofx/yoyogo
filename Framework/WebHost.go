package YoyoGo

type WebHost struct {
	webServer       IServer
	requestDelegate IRequestDelegate
}

func NewWebHost(server IServer, request IRequestDelegate) WebHost {
	return WebHost{webServer: server, requestDelegate: request}
}

func (host WebHost) Run() {
	_ = host.webServer.Run(host.requestDelegate)
}
