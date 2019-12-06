package Utils

import "reflect"

type MethodCaller struct {
	Object      interface{}
	MethodName  string
	methodInfo  reflect.Value
	paramsNum   int
	paramTypes  []reflect.Type
	paramValues []reflect.Value
}

func NewMethodCaller(obj interface{}, funcName string) *MethodCaller {
	caller := &MethodCaller{
		Object:     obj,
		MethodName: funcName,
	}
	caller.Init()

	return caller
}

func (method *MethodCaller) Init() {
	t := reflect.ValueOf(method.Object)
	method.methodInfo = t.MethodByName(method.MethodName)
	methodType := method.methodInfo.Type()

	method.paramsNum = methodType.NumIn()
	paramTypes := make([]reflect.Type, method.paramsNum)

	for idx := 0; idx < method.paramsNum; idx++ {
		paramTypes[idx] = methodType.In(idx)
	}
}

func (method *MethodCaller) Invoke(params ...interface{}) []interface{} {
	method.paramValues = make([]reflect.Value, method.paramsNum)
	for idx := 0; idx < method.paramsNum; idx++ {
		method.paramValues[idx] = reflect.ValueOf(params[idx])
	}
	returns := method.methodInfo.Call(method.paramValues)
	outNum := method.methodInfo.Type().NumOut()
	results := make([]interface{}, outNum)
	if len(returns) > 0 {
		for i, res := range returns {
			results[i] = res.Interface()
		}
	}
	return results
}
