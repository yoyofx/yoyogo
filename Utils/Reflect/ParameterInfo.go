package Reflect

import (
	"reflect"
)

type ParameterInfo struct {
	Name           string
	ParameterType  reflect.Type
	ParameterValue reflect.Value
}
