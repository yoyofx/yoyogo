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
	Kind                reflect.Kind
	IsPtr               bool
	CanSet              bool
	IsValidation        bool
	fieldInfoListCache  map[string]FieldInfo
	methodInfoListCache map[string]MethodInfo
}

// GetTypeInfo: get TypeInfo from instance
func GetTypeInfo(ctorFunc interface{}) (TypeInfo, error) {
	ctorVal := reflect.ValueOf(ctorFunc)
	return GetTypeInfoWithValueType(ctorVal, ctorVal.Type())
}

func GetTypeInfoWithValueType(ctorVal reflect.Value, ctorType reflect.Type) (TypeInfo, error) {
	var typeInfo TypeInfo
	typeInfo.IsValidation = true
	var errorInfo error = nil
	typeInfo.Kind = ctorVal.Kind()
	if typeInfo.Kind == reflect.Func {
		typeInfo.IsValidation = false
		if ctorType.NumOut() < 1 {
			errorInfo = errors.New("Can not be return out type in ctor func !")
			return typeInfo, errorInfo
		}
		outType := ctorType.Out(0)
		typeInfo.Name, typeInfo.Type, typeInfo.IsPtr = getStructOrPtrType(outType)

	} else if typeInfo.Kind == reflect.Struct || typeInfo.Kind == reflect.Ptr {
		typeInfo.IsValidation = true
		typeInfo.ValueType = ctorVal
		typeInfo.Name, typeInfo.Type, typeInfo.IsPtr = getStructOrPtrType(ctorType)
		if typeInfo.Kind == reflect.Ptr {
			typeInfo.ValueType = typeInfo.ValueType.Elem()
		}
	} else {
		errorInfo = errors.New("It's not ctor func or object instance !")
	}

	if ctorVal.CanSet() {
		typeInfo.CanSet = true
	}
	typeInfo.lazyLoadFields()
	typeInfo.LazyLoadMethods()
	return typeInfo, errorInfo
}

// HasFields the TypeInfo has fields , not empty
func (typeInfo TypeInfo) HasFields() bool {
	return len(typeInfo.fieldInfoListCache) > 0
}

// HasMethods the TypeInfo has methods , not empty
func (typeInfo TypeInfo) HasMethods() bool {
	return len(typeInfo.methodInfoListCache) > 0
}

func (typeInfo TypeInfo) GetFields() []FieldInfo {
	values := make([]FieldInfo, 0, len(typeInfo.fieldInfoListCache))
	for _, value := range typeInfo.fieldInfoListCache {
		values = append(values, value)
	}
	return nil
}

func (typeInfo TypeInfo) GetFieldByName(fieldName string) FieldInfo {
	if typeInfo.HasFields() {
		return typeInfo.fieldInfoListCache[fieldName]
	}
	panic("the TypeInfo is not fields")
}

func (typeInfo TypeInfo) GetMethods() []MethodInfo {
	values := make([]MethodInfo, 0, len(typeInfo.fieldInfoListCache))
	for _, value := range typeInfo.methodInfoListCache {
		values = append(values, value)
	}
	return nil
}

func (typeInfo TypeInfo) GetMethodByName(methodName string) MethodInfo {
	if typeInfo.HasMethods() {
		return typeInfo.methodInfoListCache[methodName]
	}
	panic("the TypeInfo is not methods")
}

// CreateInstance: create new instance by TypeInfo
func (typeInfo TypeInfo) CreateInstance() interface{} {
	return CreateInstance(typeInfo.Type)
}

func (typeInfo *TypeInfo) lazyLoadFields() {
	if len(typeInfo.fieldInfoListCache) == 0 {
		fieldNum := typeInfo.Type.NumField()
		typeInfo.fieldInfoListCache = make(map[string]FieldInfo, fieldNum)

		for i := 0; i < fieldNum; i++ {
			fieldT := typeInfo.Type.Field(i)
			fieldV := typeInfo.ValueType.Field(i)
			typeInfo.fieldInfoListCache[fieldT.Name] = FieldInfo{
				Name:  fieldT.Name,
				Type:  fieldT.Type,
				Kind:  fieldT.Type.Kind(),
				Tags:  fieldT.Tag,
				Value: fieldV,
			}
		}
	}
}

func (typeInfo *TypeInfo) LazyLoadMethods() {
	if len(typeInfo.methodInfoListCache) == 0 {
		methodList := GetObjectMethodInfoListWithValueType(typeInfo.Type, typeInfo.ValueType)
		methodNum := len(methodList)
		if methodNum > 0 {
			typeInfo.methodInfoListCache = make(map[string]MethodInfo, methodNum)
			for i := 0; i < methodNum; i++ {
				method := methodList[i]
				typeInfo.methodInfoListCache[method.Name] = method
			}
		}
	}
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
