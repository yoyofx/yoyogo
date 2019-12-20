package Utils

import "reflect"

// GetCtorFuncName
func GetCtorFuncName(ctorFunc interface{}) string {
	ctorVal := reflect.ValueOf(ctorFunc)
	if ctorVal.Kind() == reflect.Func {
		ctorType := ctorVal.Type()
		if ctorType.NumOut() < 1 {
			panic("not return controller type in ctor func !")
		}
		controllerType := ctorType.Out(0)
		if controllerType.Kind() != reflect.Ptr {
			panic("Controller type must be Ptr ! ")
		}
		controllerName := controllerType.Elem().Name()
		return controllerName
	}
	return ""
}
