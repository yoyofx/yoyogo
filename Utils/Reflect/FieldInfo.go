package Reflect

import "reflect"

type FieldInfo struct {
	Name  string
	Type  reflect.Type
	Value interface{}
}
