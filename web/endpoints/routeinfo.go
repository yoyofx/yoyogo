package endpoints

import (
	"github.com/yoyofx/yoyogo/abstractions/xlog"
	"github.com/yoyofx/yoyogo/web/context"
	"github.com/yoyofx/yoyogo/web/mvc"
	"github.com/yoyofx/yoyogo/web/router"
	"strings"
)

type controllerInfo struct {
	ControllerName string
	ActionList     []actionInfo
}
type actionInfo struct {
	Mehtod string
	Func   string
}

func UseRouteInfo(router router.IRouterBuilder) {
	xlog.GetXLogger("Endpoint").Debug("loaded router information endpoint.")

	router.GET("/actuator/route/info", func(ctx *context.HttpContext) {
		var builder *mvc.ControllerBuilder
		_ = ctx.RequiredServices.GetService(&builder)

		descriptorList := builder.GetControllerDescriptorList()

		controllerList := make([]controllerInfo, 0)
		for _, desc := range descriptorList {
			actionList := make([]actionInfo, 0)
			for _, action := range desc.GetActionDescriptors() {
				actionInfo := actionInfo{}
				actionInfo.Mehtod = strings.ToUpper(action.ActionMethod)
				actionInfo.Func = action.ActionName
				actionList = append(actionList, actionInfo)
			}

			controllerList = append(controllerList, controllerInfo{ControllerName: desc.ControllerName, ActionList: actionList})
		}
		retMap := make(map[string]interface{})
		retMap["MVC"] = controllerList

		retMap["Route"] = router.GetRouteInfo()

		ctx.JSON(200, retMap)
	})
}
