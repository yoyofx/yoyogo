package Mvc

import (
	"github.com/yoyofx/yoyogo/WebFramework/Context"
	"github.com/yoyofxteam/reflectx"
	"reflect"
	"strings"
)

// ControllerDescriptor
type ControllerDescriptor struct {
	ControllerName    string
	ControllerType    interface{} // ctor func of controller
	actionDescriptors map[string]ActionDescriptor
}

// NewControllerDescriptor create new controller descriptor
func NewControllerDescriptor(name string, controllerType reflect.Type, controllerCtor interface{}) ControllerDescriptor {

	instance := reflect.New(controllerType).Interface()
	actionList := reflectx.GetObjectMethodInfoList(instance)

	actionDescriptors := make(map[string]ActionDescriptor, len(actionList))

	for _, action := range actionList {
		actionName := strings.ToLower(action.Name)
		actionDescriptors[actionName] = ActionDescriptor{
			ActionName:   action.Name,
			ActionMethod: getHttpMethodByActionName(actionName),
			MethodInfo:   action,
		}
	}

	return ControllerDescriptor{name, controllerCtor, actionDescriptors}
}

// GetActionDescriptors get action descriptor list
func (controllerDescriptor ControllerDescriptor) GetActionDescriptors() []ActionDescriptor {

	values := make([]ActionDescriptor, 0, len(controllerDescriptor.actionDescriptors))
	for _, value := range controllerDescriptor.actionDescriptors {
		values = append(values, value)
	}
	return values
}

// GetActionDescriptorByName get action descriptor by name
func (controllerDescriptor ControllerDescriptor) GetActionDescriptorByName(actionName string) (ActionDescriptor, bool) {
	actionDescriptor, ok := controllerDescriptor.actionDescriptors[actionName]
	return actionDescriptor, ok
}

func getHttpMethodByActionName(actionNameLower string) string {
	methodName := "any"
	for _, m := range Context.Methods {
		method := strings.ToLower(m)
		if strings.HasPrefix(actionNameLower, method) {
			methodName = method
			break
		}
	}
	return methodName
}
