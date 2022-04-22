package router

import (
	"github.com/yoyofx/yoyogo/web/context"
	"net/url"
	"strings"
)

// EndPointRouterHandler node.
type EndPointRouterHandler struct {
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
func (t *EndPointRouterHandler) Insert(method, path string, handler func(ctx *context.HttpContext)) {
	t.fullURL = &path
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
			Methods: make(map[string]func(ctx *context.HttpContext))}
		newNode.fullURL = &path
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

func (t *EndPointRouterHandler) Match(ctx *context.HttpContext, pathComponents []string) (string, bool) {
	node := t.search(pathComponents, ctx.Input.RouterData)
	if node != nil {
		return *node.fullURL, true
	}
	return "", false
}

// Search the tree.
func (t *EndPointRouterHandler) search(components []string, params url.Values) *EndPointRouterHandler {
Next:
	for cidx, component := range components {
		if t.Component == component && cidx == 0 {
			continue
		} else if t.Component != "/" && t.Component != component && cidx == 0 {
			return nil
		}
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
