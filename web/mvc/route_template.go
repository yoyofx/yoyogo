package mvc

import "strings"

type RouteTemplate struct {
	Template        string
	TemplateSplits  []string
	controllerIndex int
	actionIndex     int
	pathLen         int
}

type MatchMvcInfo struct {
	ControllerName string
	ActionName     string
}

func NewRouteTemplate(temp string) *RouteTemplate {
	templateSplits := strings.Split(temp, "/")

	return &RouteTemplate{
		Template:        temp,
		controllerIndex: findIndex("{controller}", templateSplits),
		actionIndex:     findIndex("{action}", templateSplits),
		pathLen:         len(templateSplits),
		TemplateSplits:  templateSplits,
	}
}

func (template *RouteTemplate) Match(pathComponents []string, matchinfo *MatchMvcInfo) bool {
	if len(pathComponents) >= template.pathLen {
		matchinfo.ControllerName = pathComponents[template.GetControllerIndex()]
		matchinfo.ControllerName = strings.ToLower(matchinfo.ControllerName)
		if !strings.Contains(matchinfo.ControllerName, "controller") {
			matchinfo.ControllerName += "controller"
		}

		matchinfo.ActionName = strings.ToLower(pathComponents[template.GetActionIndex()])

		for index, item := range pathComponents {
			if index != template.GetControllerIndex() && index != template.GetActionIndex() {
				if item != template.TemplateSplits[index] {
					return false
				}
			}
		}

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
