package gee

import "fmt"

type Route struct {
	handlers map[string]HandlerFunc
}

func newRoute() *Route {
	return &Route{
		handlers: make(map[string]HandlerFunc),
	}
}

//添加路由
func (route *Route) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	route.handlers[key] = handlerFunc
}

func (r *Route) handler(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		fmt.Fprintf(c.Writer, "404 Not Fount URL = %s", c.Path)
	}

}
