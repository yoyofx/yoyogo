package Mvc

import (
	"github.com/yoyofx/yoyogo/Utils/Reflect"
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
func NewControllerDescriptor(name string, cType reflect.Type, controllerCtor interface{}) ControllerDescriptor {

	instance := Reflect.CreateInstance(cType)
	actionList := Reflect.GetObjectMehtodInfoList(instance)

	actionDescriptors := make(map[string]ActionDescriptor, len(actionList))

	for _, action := range actionList {
		actionDescriptors[action.Name] = ActionDescriptor{
			ActionName: strings.ToLower(action.Name),
			MethodInfo: action,
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
func (controllerDescriptor ControllerDescriptor) GetActionDescriptorByName(actionName string) ActionDescriptor {
	return controllerDescriptor.actionDescriptors[actionName]
}
