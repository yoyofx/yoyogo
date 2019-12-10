package Utils

import "reflect"

type MethodCaller struct {
	Object      interface{}
	MethodName  string
	methodInfo  reflect.Value
	paramsNum   int
	paramTypes  []reflect.Type
	paramValues []reflect.Value
	foundMethod bool
}

func NewMethodCaller(obj interface{}, funcName string) *MethodCaller {
	caller := &MethodCaller{
		Object:     obj,
		MethodName: funcName,
	}
	caller.foundMethod = caller.findMethod()
	if caller.foundMethod {
		return caller
	}
	return nil
}

func (method *MethodCaller) GetParamTypes() []reflect.Type {
	return method.paramTypes
}

func (method *MethodCaller) findMethod() bool {
	t := reflect.ValueOf(method.Object)
	method.methodInfo = t.MethodByName(method.MethodName)
	if !method.methodInfo.IsValid() {
		return false
	}
	methodType := method.methodInfo.Type()

	method.paramsNum = methodType.NumIn()
	method.paramTypes = make([]reflect.Type, method.paramsNum)

	for idx := 0; idx < method.paramsNum; idx++ {
		method.paramTypes[idx] = methodType.In(idx)
	}
	return true
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
