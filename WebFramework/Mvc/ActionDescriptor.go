package Mvc

import "github.com/yoyofxteam/reflectx"

type ActionDescriptor struct {
	ActionName string
	MethodInfo reflectx.MethodInfo
}
