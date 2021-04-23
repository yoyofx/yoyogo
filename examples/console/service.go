package main

import "fmt"

type Service1 struct {
}

func NewService() *Service1 {
	return &Service1{}
}

func (s *Service1) Run() error {
	fmt.Println("host service Running")
	return nil
}

func (s *Service1) Stop() error {
	fmt.Println("host service Stopping")
	return nil
}
