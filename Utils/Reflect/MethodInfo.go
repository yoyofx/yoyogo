package Reflect

import "reflect"

type MethodInfo struct {
	Name           string
	MethodInfoVal  reflect.Value
	MethodInfoType reflect.Type
	Parameters     []ParameterInfo
}

func (method MethodInfo) IsValid() bool {
	return method.MethodInfoVal.IsValid()
}

func (method MethodInfo) Invoke(params ...interface{}) []interface{} {
	paramsCount := len(method.Parameters)
	paramsValues := make([]reflect.Value, paramsCount)
	for idx := 0; idx < paramsCount; idx++ {
		method.Parameters[idx].ParameterValue = reflect.ValueOf(params[idx])
		paramsValues[idx] = method.Parameters[idx].ParameterValue
	}
	returns := method.MethodInfoVal.Call(paramsValues)
	outNum := method.MethodInfoType.NumOut()
	results := make([]interface{}, outNum)
	if len(returns) > 0 {
		for i, res := range returns {
			results[i] = res.Interface()
		}
	}
	return results
}
