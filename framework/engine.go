package framework

import (
	"net/http"
)

type handlerFunc func(ctx *Context)

type engine struct {
	router *router
	*routerGroup
}

type routerGroup struct {
	prefix string
	engine *engine
}

func New() *engine {
	e := &engine{router: newRouter()}
	e.routerGroup = &routerGroup{engine: e}

	return e
}

func (group *routerGroup) Group(prefix string) *routerGroup {
	engine := group.engine
	newGroup := &routerGroup{
		prefix: group.prefix + prefix,
		engine: engine,
	}

	return newGroup
}

func (group *routerGroup) addRoute(method string, pattern string, handler handlerFunc) {
	group.engine.router.addRoute(method, group.prefix+pattern, handler)
}

func (group *routerGroup) GET(pattern string, handler handlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *routerGroup) POST(pattern string, handler handlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (engine *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	engine.router.handle(newContext(w, req))
}

func (engine *engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
