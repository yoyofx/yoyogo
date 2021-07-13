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
	if s.helloworldApi != nil {
		err := s.helloworldApi.SayRecord(&pb.HelloRequest{})
		if err != nil {
			return err
		}
	} else {
		fmt.Println("grpc is not connected !!")
	}
	return nil
}

func (s *ClientService) Stop() error {
	fmt.Println("host service Stopping")
	return nil
}
