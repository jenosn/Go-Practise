package gee

import (
    "log"
    "net/http"
)

type HandlerFunc func(ctx *Context)

type Engine struct {
    route *Route
}

func NewEngine() *Engine {
    return &Engine{
        route: newRoute(),
    }
}

func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
    log.Printf("Route %4s  %s", method, pattern)
    engine.route.addRoute(method, pattern, handlerFunc)
}

func (engine *Engine) Get(pattern string, handlerFunc HandlerFunc) {
    log.Printf("Route %4s", pattern)
    engine.route.addRoute("GET", pattern, handlerFunc)
}

func (engine *Engine) Post(pattern string, handlerFunc HandlerFunc) {
    log.Printf("Route %4s ", pattern)
    engine.route.addRoute("POST", pattern, handlerFunc)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    c := newContext(w, r)
    engine.route.handler(c)
}

func (engine *Engine) Run(add string) {
    http.ListenAndServe(add, engine)
}
