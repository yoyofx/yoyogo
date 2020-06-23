package YoyoGo

func (self *WebHostBuilder) UseFastHttpByAddr(addr string) *WebHostBuilder {
	self.server = NewFastHttp(addr)
	return self
}

func (self *WebHostBuilder) UseFastHttp() *WebHostBuilder {
	self.server = NewFastHttp("")
	return self
}

func (self *WebHostBuilder) UseHttpByAddr(addr string) *WebHostBuilder {
	self.server = DefaultHttpServer(addr)
	return self
}

func (self *WebHostBuilder) UseHttp() *WebHostBuilder {
	self.server = DefaultHttpServer("")
	return self
}
