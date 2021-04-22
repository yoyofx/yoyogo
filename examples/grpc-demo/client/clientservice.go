package main

import (
	"fmt"
	pb "grpc-demo/proto/helloworld"
)

type ClientService struct {
	helloworldApi *Api
}

func NewClientService(api *Api) *ClientService {
	return &ClientService{helloworldApi: api}
}

func (s *ClientService) Run() error {
	fmt.Println("host service Running")
	err := s.helloworldApi.SayRecord(&pb.HelloRequest{})
	if err != nil {
		return err
	}

	return nil
}

func (s *ClientService) Stop() error {
	fmt.Println("host service Stopping")
	return nil
}
