package Reflect

import (
	"errors"
	"reflect"
)

// TypeInfo
type TypeInfo struct {
	Name                string
	ValueType           reflect.Value
	Type                reflect.Type
	IsPtr               bool
	IsValidation        bool
	fieldInfoListCache  map[string]FieldInfo
	methodInfoListCache map[string]MethodInfo
}

// GetTypeInfo: get TypeInfo from instance
func GetTypeInfo(ctorFunc interface{}) (TypeInfo, error) {
	var typeInfo TypeInfo
	typeInfo.IsValidation = true
	var errorInfo error = nil
	ctorVal := reflect.ValueOf(ctorFunc)
	ctorType := ctorVal.Type()
	ctorKind := ctorVal.Kind()
	if ctorKind == reflect.Func {
		typeInfo.IsValidation = false
		if ctorType.NumOut() < 1 {
			errorInfo = errors.New("Can not be return out type in ctor func !")
			return typeInfo, errorInfo
		}
		outType := ctorType.Out(0)
		typeInfo.Name, typeInfo.Type, typeInfo.IsPtr = getStructOrPtrType(outType)

	} else if ctorKind == reflect.Struct || ctorKind == reflect.Ptr {
		typeInfo.IsValidation = true
		typeInfo.ValueType = ctorVal
		typeInfo.Name, typeInfo.Type, typeInfo.IsPtr = getStructOrPtrType(ctorType)
	}
	return typeInfo, errorInfo
}

// CreateInstance: create new instance by TypeInfo
func (typeInfo TypeInfo) CreateInstance() interface{} {
	return CreateInstance(typeInfo.Type)
}

func (typeInfo TypeInfo) LazyLoadFields() {

}

func (typeInfo TypeInfo) LazyLoadMethods() {

}

func (typeInfo TypeInfo) GetFields() []FieldInfo {

	return nil
}

// getStructOrPtrType: get Struct Or Ptr type (name , type , isPtr)
func getStructOrPtrType(outType reflect.Type) (string, reflect.Type, bool) {
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
