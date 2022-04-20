package mvc

import (
	"github.com/yoyofx/yoyogo/web/context"
)

type RouteAttribute struct {
	Template   string `json:"routeTemplate"`
	Controller string `json:"controller"`
	Action     string `json:"action"`
	Method     string `json:"method"`
}

func NewRouteAttribute(template string) RouteAttribute {
	return RouteAttribute{Template: template}
}

//---------------------------------------------------------------------------

type RouteAttributeCollection struct {
	elementCount uint64
	enabled      bool
	processor    IRouterAttributeBuilder
	Collection   map[string]RouteAttribute
}

func NewRouteAttributeCollection() *RouteAttributeCollection {
	return &RouteAttributeCollection{
		enabled: false, elementCount: 0, Collection: make(map[string]RouteAttribute)}
}

func (collection *RouteAttributeCollection) Add(routeAttribute RouteAttribute) {
	collection.elementCount++
	collection.Collection[routeAttribute.Template] = routeAttribute
	collection.processor.Insert(routeAttribute.Method, routeAttribute.Template, EmptyHandler)
}

func (collection *RouteAttributeCollection) HasElement() bool {
	return collection.elementCount > 0
}

func (collection *RouteAttributeCollection) Enable() {
	collection.enabled = true
}

func (collection *RouteAttributeCollection) Disable() {
	collection.enabled = false
}

func (collection *RouteAttributeCollection) IsEnable() bool {
	return collection.enabled
}

func (collection *RouteAttributeCollection) SetProcessor(processor IRouterAttributeBuilder) {
	collection.processor = processor
}

func (collection *RouteAttributeCollection) Match(ctx *context.HttpContext, pathComponents []string, match *MatchMvcInfo) bool {
	if collection.IsEnable() && collection.HasElement() {
		originTemplate, found := collection.processor.Match(ctx, pathComponents)
		if found {
			attribute, hasElem := collection.Collection[originTemplate]
			if hasElem {
				match.ControllerName = attribute.Controller
				match.ActionName = attribute.Action
			}
			return hasElem
		}
	}
	return false
}

func EmptyHandler(ctx *context.HttpContext) {}
