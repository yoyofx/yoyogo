package Mvc

import "strings"

type RouteTemplate struct {
	Template        string
	ControllerName  string
	ActionName      string
	controllerIndex int
	actionIndex     int
	pathLen         int
}

func NewRouteTemplate(temp string) RouteTemplate {
	templateSplits := strings.Split(temp, "/")

	return RouteTemplate{
		Template:        temp,
		controllerIndex: findIndex("{controller}", templateSplits),
		actionIndex:     findIndex("{action}", templateSplits),
		pathLen:         len(templateSplits),
	}
}

func (template *RouteTemplate) Match(pathComponents []string) bool {
	if len(pathComponents) >= template.pathLen {
		template.ControllerName = pathComponents[template.GetControllerIndex()]
		template.ActionName = pathComponents[template.GetActionIndex()]
		return true
	}
	return false
}

func (template *RouteTemplate) GetControllerIndex() int {
	return template.controllerIndex
}

func (template *RouteTemplate) GetActionIndex() int {
	return template.actionIndex
}

func findIndex(it string, ins []string) int {
	idx := 0
Loop:
	for index, item := range ins {
		if it == item {
			idx = index
			break Loop
		}
	}
	return idx
}
