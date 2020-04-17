package gee

import (
	"fmt"
	"log"
	"net/http"
)

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

//引擎
type Engine struct {
	route map[string]HandlerFunc
}

func (engine *Engine) Get(pattern string, handlerFunc HandlerFunc) {
	engine.addRoute("GET", pattern, handlerFunc)

}

func (engine *Engine) Post(pattern string, handlerFunc HandlerFunc) {
	engine.addRoute("POST", pattern, handlerFunc)

}

//添加路由
func (engine *Engine) addRoute(method string, pattern string, handlerFunc HandlerFunc) {
	key := method + "-" + pattern
	log.Printf("Route %4s - %s", method, pattern)
	engine.route[key] = handlerFunc
}

//创建路由引擎
func New() *Engine {
	return &Engine{route: make(map[string]HandlerFunc)}
}

//启动
func (engine *Engine) Run(add string) (err error) {
	err = http.ListenAndServe(add, engine)
	return
}

//引擎必须实现ServeHttp(必须步骤)
func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	if handler, ok := engine.route[key];
		ok {
		handler(w, r);
	} else {
		fmt.Fprintf(w, "404 Not Found URL =%s ", r.URL.Path)
	}
}
