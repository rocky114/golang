package framework

import (
	"net/http"
)

type handlerFunc func(ctx *Context)

type engine struct {
	router *router
}

func New() *engine {
	return &engine{
		router: newRouter(),
	}
}

func (engine *engine) addRoute(method string, pattern string, handler handlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

func (engine *engine) GET(pattern string, handler handlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

func (engine *engine) POST(pattern string, handler handlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

func (engine *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	engine.router.handle(newContext(w, req))
}

func (engine *engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
