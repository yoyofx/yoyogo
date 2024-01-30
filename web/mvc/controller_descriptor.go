package mvc

import (
	"errors"
	"github.com/nacos-group/nacos-sdk-go/common/logger"
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
	var actionDescriptor ActionDescriptor
	for _, action := range actionList {
		actionName := strings.ToLower(action.Name)
		if !utils.ContainsStr([]string{"apiresult", "fail", "ok", "setviewengine", "view", "getname"}, actionName) {
			actionDescriptor = ActionDescriptor{
				ActionName:   action.Name,
				ActionMethod: getHttpMethodByActionName(actionName),
				MethodInfo:   action,
			}
			// Action Descriptors
			attributeRoute, err := addAttributeRouteActionDescriptor(name, actionDescriptor)
			if err != nil {
				logger.Error(err.Error())
			} else {
				if attributeRoute != nil {
					actionDescriptor.IsAttributeRoute = true
					actionDescriptor.Route = attributeRoute
				}
			}
			actionDescriptors[actionName] = actionDescriptor
		}
	}

	return ControllerDescriptor{name, controllerDoc, controllerCtor, actionDescriptors}, nil
}

func addAttributeRouteActionDescriptor(controllerName string, desc ActionDescriptor) (*RouteAttribute, error) {
	for _, parameter := range desc.MethodInfo.Parameters {
		paramType := parameter.ParameterType
		if paramType.Kind() == reflect.Ptr {
			paramType = paramType.Elem()
		}
		if paramType.Kind() == reflect.Struct {
			if paramType.NumField() > 0 {
				fieldName := paramType.Field(0).Name
				if fieldName == "RequestBody" || fieldName == "RequestGET" || fieldName == "RequestPOST" {
					routeTemplate := paramType.Field(0).Tag.Get("route")
					if routeTemplate != "" {
						routeAttr := NewRouteAttribute(routeTemplate)
						routeAttr.Controller = controllerName
						routeAttr.Action = strings.ToLower(desc.ActionName)
						if fieldName == "RequestBody" || fieldName == "RequestPOST" {
							routeAttr.Method = "POST"
						} else {
							routeAttr.Method = "GET"
						}

						logger.Debug("add mvc controller action for attributes:{[%s/%s],method=[%s]}", strings.Replace(controllerName, "controller", "", -1), strings.ToLower(desc.ActionName), strings.ToUpper(desc.ActionMethod))
						return &routeAttr, nil
					}
				}
			}
		}
	}
	return nil, errors.New("not found route attribute")
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
