package Router

import (
	"github.com/maxzhang1985/yoyogo/WebFramework/Context"
	"net/url"
	"strings"
)

// EndPointRouterHandler node.
type EndPointRouterHandler struct {
	children  []*EndPointRouterHandler
	param     byte
	Component string
	Methods   map[string]func(ctx *Context.HttpContext)
}

func (endPoint *EndPointRouterHandler) Invoke(ctx *Context.HttpContext, pathComponents []string) func(ctx *Context.HttpContext) {
	var handler func(ctx *Context.HttpContext) = nil
	node := endPoint.search(pathComponents, ctx.RouterData)
	if node != nil && node.Methods[ctx.Request.Method] != nil {
		handler = node.Methods[ctx.Request.Method]
	}
	return handler
}

// Insert a node into the tree.
func (t *EndPointRouterHandler) Insert(method, path string, handler func(ctx *Context.HttpContext)) {
	components := strings.Split(path, "/")[1:]
Next:
	for _, component := range components {
		for _, child := range t.children {
			if child.Component == component {
				t = child
				continue Next
			}
		}
		newNode := &EndPointRouterHandler{Component: component,
			Methods: make(map[string]func(ctx *Context.HttpContext))}
		if len(component) > 0 {
			if component[0] == ':' || component[0] == '*' {
				newNode.param = component[0]
			}
		}
		t.children = append(t.children, newNode)
		t = newNode
	}
	t.Methods[method] = handler
}

// Search the tree.
func (t *EndPointRouterHandler) search(components []string, params url.Values) *EndPointRouterHandler {
Next:
	for _, component := range components {
		for _, child := range t.children {
			if child.Component == component || child.param == ':' || child.param == '*' {
				if child.param == '*' {
					return child
				}
				if child.param == ':' {
					params.Add(child.Component[1:], component)
				}
				t = child
				continue Next
			}
		}
		return nil // not found
	}
	return t
}
