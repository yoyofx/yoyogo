package console

import (
	"fmt"
	"github.com/yoyofx/yoyogo/abstractions"
)

type Server struct {
	exit chan int
}

func NewServer() *Server {
	return &Server{exit: make(chan int)}
}

func (server *Server) GetAddr() string {
	return ""
}

func (server *Server) Run(context *abstractions.HostBuilderContext) (e error) {
	fmt.Println("server running")
	s := <-server.exit
	fmt.Println("server exit ", s)
	return nil
}

func (server *Server) Shutdown() {
	server.exit <- 0
	//close(server.exit)
}
