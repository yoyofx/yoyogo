package Middleware

import (
	"net/url"
	"strings"
)

// Trie node.
type Trie struct {
	children  []*Trie
	param     byte
	Component string
	Methods   map[string]func(ctx *HttpContext)
}

// Insert a node into the tree.
func (t *Trie) Insert(method, path string, handler func(ctx *HttpContext)) {
	components := strings.Split(path, "/")[1:]
Next:
	for _, component := range components {
		for _, child := range t.children {
			if child.Component == component {
				t = child
				continue Next
			}
		}
		newNode := &Trie{Component: component,
			Methods: make(map[string]func(ctx *HttpContext))}
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
func (t *Trie) Search(components []string, params url.Values) *Trie {
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
