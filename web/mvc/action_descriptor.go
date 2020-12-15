package mvc

import "github.com/yoyofxteam/reflectx"

type ActionDescriptor struct {
	ActionName   string
	ActionMethod string
	MethodInfo   reflectx.MethodInfo
}
