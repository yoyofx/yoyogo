package Reflect

import (
	"reflect"
)

// GetCtorFuncOutTypeName get ctor function return type's name.
func GetCtorFuncOutTypeName(ctorFunc interface{}) (string, reflect.Type) {
	typeInfo, err := GetTypeInfo(ctorFunc)
	if err != nil {
		panic(err.Error())
	}
	return typeInfo.Name, typeInfo.Type
}

// getMehtodInfo get method info
func getMehtodInfo(method reflect.Method, methodValue reflect.Value) MethodInfo {
	methodInfo := MethodInfo{}
	methodInfo.MethodInfoVal = methodValue
	methodInfo.MethodInfoType = methodValue.Type()
	methodInfo.Name = method.Name
	paramsCount := methodInfo.MethodInfoType.NumIn()
	methodInfo.Parameters = make([]ParameterInfo, paramsCount)

	for idx := 0; idx < paramsCount; idx++ {
		methodInfo.Parameters[idx].ParameterType = methodInfo.MethodInfoType.In(idx)
		methodInfo.Parameters[idx].Name = methodInfo.Parameters[idx].ParameterType.Name()
	}

	return methodInfo
}

func GetObjectMehtodInfoList(object interface{}) []MethodInfo {
	objectType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)

	methodCount := objValue.NumMethod()
	methodInfos := make([]MethodInfo, methodCount)
	for idx := 0; idx < methodCount; idx++ {
		methodInfo := getMehtodInfo(objectType.Method(idx), objValue.Method(idx))
		methodInfos[idx] = methodInfo
	}

	return methodInfos
}

func GetObjectMehtodInfoByName(object interface{}, methodName string) MethodInfo {
	objType := reflect.TypeOf(object)
	objValue := reflect.ValueOf(object)
	methodType, _ := objType.MethodByName(methodName)
	methodInfo := getMehtodInfo(methodType, objValue.MethodByName(methodName))
	return methodInfo
}
