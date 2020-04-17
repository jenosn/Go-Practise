package gee

import "log"

type HandlerFunc func(ctx *Context)

type Engine struct {
	route *Route
}

func newEngine() *Engine {
	return &Engine{
		route: newRoute(),
	}
}

func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	log.Printf("Route %4s  %s",method,pattern)
	engine.route.addRoute(method, pattern, handlerFunc)
}
