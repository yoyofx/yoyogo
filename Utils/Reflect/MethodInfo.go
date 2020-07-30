package Reflect

//
//import (
//	"reflect"
//)
//
//// Method Info
//type MethodInfo struct {
//	Name           string          //Method Name
//	MethodInfoVal  reflect.Value   //method value
//	MethodInfoType reflect.Type    //method type
//	Parameters     []ParameterInfo //method's Parameters
//	OutType        reflect.Type    //function's return type.
//}
//
//// IsValid : method is valid
//func (method MethodInfo) IsValid() bool {
//	return method.MethodInfoVal.IsValid()
//}
//
//// Invoke : invoke the method with interface params.
//func (method MethodInfo) Invoke(params ...interface{}) []interface{} {
//	paramsCount := len(method.Parameters)
//	paramsValues := make([]reflect.Value, paramsCount)
//	for idx := 0; idx < paramsCount; idx++ {
//		method.Parameters[idx].ParameterValue = reflect.ValueOf(params[idx])
//		paramsValues[idx] = method.Parameters[idx].ParameterValue
//	}
//
//	return method.InvokeWithValue(paramsValues...)
//}
//
//// InvokeWithValue: invoke the method with value params.
//func (method MethodInfo) InvokeWithValue(paramsValues ...reflect.Value) []interface{} {
//	returns := method.MethodInfoVal.Call(paramsValues)
//	outNum := method.MethodInfoType.NumOut()
//	results := make([]interface{}, outNum)
//	if len(returns) > 0 {
//		for i, res := range returns {
//			results[i] = res.Interface()
//		}
//	}
//	return results
//}
//
//// AsTypeInfo : convert method to TypeInfo
//func (method MethodInfo) AsTypeInfo() (TypeInfo, error) {
//	return GetTypeInfoWithValueType(method.MethodInfoVal, method.MethodInfoType)
//}
