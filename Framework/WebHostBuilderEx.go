package YoyoGo

func (self *HostBuilder) UseFastHttpByAddr(addr string) *HostBuilder {
	self.server = NewFastHttp(addr)
	return self
}

func (self *HostBuilder) UseFastHttp() *HostBuilder {
	self.server = NewFastHttp("")
	return self
}

func (self *HostBuilder) UseHttpByAddr(addr string) *HostBuilder {
	self.server = DefaultHttpServer(addr)
	return self
}

func (self *HostBuilder) UseHttp() *HostBuilder {
	self.server = DefaultHttpServer("")
	return self
}
