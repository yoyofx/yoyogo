package router

import (
	"github.com/yoyofx/yoyogo/web/context"
	"net/url"
	"sort"
	"strings"
)

// EndPointRouterHandler node.
type EndPointRouterHandler struct {
	// 需要数组排序,Component是否包含:和*,如果包含则放在最后
	children  []*EndPointRouterHandler
	param     byte
	Component string
	fullURL   *string
	Methods   map[string]func(ctx *context.HttpContext)
}

func (endPoint *EndPointRouterHandler) Invoke(ctx *context.HttpContext, pathComponents []string) func(ctx *context.HttpContext) {
	var handler func(ctx *context.HttpContext) = nil
	node := endPoint.search(pathComponents, ctx.Input.RouterData)
	if node != nil && node.Methods[ctx.Input.Method()] != nil {
		handler = node.Methods[ctx.Input.Method()]
	}
	return handler
}

// Insert a node into the tree.
func (endPoint *EndPointRouterHandler) Insert(method, path string, handler func(ctx *context.HttpContext)) {
	endPoint.fullURL = &path
	components := strings.Split(path, "/")[1:]
Next:
	for _, component := range components {
		for _, child := range endPoint.children {
			if child.Component == component {
				endPoint = child
				continue Next
			}
		}
		newNode := &EndPointRouterHandler{Component: component,
			Methods: make(map[string]func(ctx *context.HttpContext))}
		newNode.fullURL = &path
		if len(component) > 0 {
			if component[0] == ':' || component[0] == '*' {
				newNode.param = component[0]
			}
		}
		endPoint.children = append(endPoint.children, newNode)
		endPoint = newNode
	}
	endPoint.Methods[method] = handler
}

func (endPoint *EndPointRouterHandler) Match(ctx *context.HttpContext, pathComponents []string) (string, bool) {
	node := endPoint.search(pathComponents, ctx.Input.RouterData)
	if node != nil {
		return *node.fullURL, true
	}
	return "", false
}

// Search the tree.
func (endPoint *EndPointRouterHandler) search(components []string, params url.Values) *EndPointRouterHandler {
Next:
	for cidx, component := range components {
		if endPoint.Component == component && cidx == 0 {
			continue
		} else if endPoint.Component != "/" && endPoint.Component != component && cidx == 0 {
			return nil
		}

		sort.Slice(endPoint.children, endPoint.Less)
		for _, child := range endPoint.children {

			if child.Component == component || child.param == ':' || child.param == '*' {
				if child.param == '*' {
					return child
				}
				if child.param == ':' {
					params.Add(child.Component[1:], component)
				}
				endPoint = child
				continue Next
			}
		}
		return nil // not found
	}
	return endPoint
}

// Less sort by EndPointRouterHandler
// Less function for sort,
func (endPoint *EndPointRouterHandler) Less(i, j int) bool {
	if endPoint.children[i].param == ':' || endPoint.children[i].param == '*' {
		return false
	}
	return true
}
