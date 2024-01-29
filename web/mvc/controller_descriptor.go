package mvc

import (
	"errors"
	"github.com/yoyofx/yoyogo/utils"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofxteam/reflectx"
	"reflect"
	"strings"
)

// ControllerDescriptor
type ControllerDescriptor struct {
	ControllerName    string
	Descriptor        string
	ControllerType    interface{} // ctor func of controller
	actionDescriptors map[string]ActionDescriptor
}

// NewControllerDescriptor create new controller descriptor
func NewControllerDescriptor(name string, controllerType reflect.Type, controllerCtor interface{}) (ControllerDescriptor, error) {

	fieldApiController := controllerType.Field(0)
	if fieldApiController.Name != "ApiController" {
		return ControllerDescriptor{}, errors.New("controller must be embed field0 ApiController")
	}
	controllerDoc := fieldApiController.Tag.Get("doc")

	instance := reflect.New(controllerType).Interface()
	actionList := reflectx.GetObjectMethodInfoList(instance)

	actionDescriptors := make(map[string]ActionDescriptor, len(actionList))

	for _, action := range actionList {
		actionName := strings.ToLower(action.Name)
		if !utils.ContainsStr([]string{"apiresult", "fail", "ok", "setviewengine", "view", "getname"}, actionName) {
			actionDescriptors[actionName] = ActionDescriptor{
				ActionName:   action.Name,
				ActionMethod: getHttpMethodByActionName(actionName),
				MethodInfo:   action,
			}
		}
	}

	return ControllerDescriptor{name, controllerDoc, controllerCtor, actionDescriptors}, nil
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
	for _, m := range context.Methods {
		method := strings.ToLower(m)
		if strings.HasPrefix(actionNameLower, method) {
			methodName = method
			break
		}
	}
	return methodName
}
