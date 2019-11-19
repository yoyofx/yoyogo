package Router

import "github.com/maxzhang1985/yoyogo/Context"

//func NewRouterHandler() *Trie {
//	return &Trie{
//		Component: "/",
//		Methods:   make(map[string]func(ctx *Context.HttpContext)),
//	}
//}

func NewRouterHandler() IRouterHandler {
	tree := &Trie{
		Component: "/",
		Methods:   make(map[string]func(ctx *Context.HttpContext)),
	}

	defaultRouterHandler := &DefaultRouterHandler{routerTree: tree}

	return defaultRouterHandler
}
