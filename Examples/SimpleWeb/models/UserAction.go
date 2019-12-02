package models

import (
	"math/rand"
	"strconv"
	"time"
)

type IUserAction interface {
	Login(name string) string
}

type UserAction struct {
	index int
}

func NewUserAction() *UserAction {
	rand.Seed(time.Now().Unix())
	rnd := rand.Intn(100)
	return &UserAction{index: rnd}
}

func (u UserAction) Login(name string) string {
	return "hello " + name + strconv.Itoa(u.index)
}
