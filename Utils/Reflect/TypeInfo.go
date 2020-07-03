package Reflect

import (
	"errors"
	"reflect"
)

// TypeInfo
type TypeInfo struct {
	Name  string
	Type  reflect.Type
	IsPtr bool
}

func getTypeInfo(outType reflect.Type) (string, reflect.Type, bool) {
	var name string
	var cType reflect.Type
	isPtr := false
	if outType.Kind() != reflect.Ptr {
		name = outType.Name()
		cType = outType
	} else {
		isPtr = true
		name = outType.Elem().Name()
		cType = outType.Elem()
	}
	return name, cType, isPtr
}

func GetTypeInfo(ctorFunc interface{}) (TypeInfo, error) {
	var typeInfo TypeInfo
	var errorInfo error = nil
	ctorVal := reflect.ValueOf(ctorFunc)
	ctorType := ctorVal.Type()
	ctorKind := ctorVal.Kind()
	if ctorKind == reflect.Func {

		if ctorType.NumOut() < 1 {
			errorInfo = errors.New("Can not be return out type in ctor func !")
			return typeInfo, errorInfo
		}
		outType := ctorType.Out(0)
		typeInfo.Name, typeInfo.Type, typeInfo.IsPtr = getTypeInfo(outType)

	} else if ctorKind == reflect.Struct || ctorKind == reflect.Ptr {
		typeInfo.Name, typeInfo.Type, typeInfo.IsPtr = getTypeInfo(ctorType)
	}
	return typeInfo, errorInfo
}

func (typeInfo TypeInfo) CreateInstance() interface{} {
	return CreateInstance(typeInfo.Type)
}

//func CreateInstance(objectType reflect.Type) interface{} {
//	return reflect.New(objectType).Elem().Interface()
//}

func CreateInstance(objectType reflect.Type) interface{} {
	var ins reflect.Value

	ins = reflect.New(objectType)

	if objectType.Kind() == reflect.Struct {
		ins = ins.Elem()
	}

	return ins.Interface()
}
