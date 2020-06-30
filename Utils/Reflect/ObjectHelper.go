package Reflect

import (
	"reflect"
)

// GetCtorFuncOutTypeName get ctor function return type's name.
func GetCtorFuncOutTypeName(ctorFunc interface{}) (string, reflect.Type) {
	ctorVal := reflect.ValueOf(ctorFunc)
	if ctorVal.Kind() == reflect.Func {
		ctorType := ctorVal.Type()
		if ctorType.NumOut() < 1 {
			panic("not return controller type in ctor func !")
		}
		controllerType := ctorType.Out(0)
		if controllerType.Kind() != reflect.Ptr {
			panic("Mvc type must be Ptr ! ")
		}
		controllerName := controllerType.Elem().Name()
		return controllerName, controllerType.Elem()
	}
	return "", nil
}

func CreateInstance(objectType reflect.Type) interface{} {
	return reflect.New(objectType).Interface()
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
