package framework

import (
	"net/http"
)

type handlerFunc func(ctx *Context)

type engine struct {
	router     *router
	middleware *middleware
	*routerGroup
}

type routerGroup struct {
	prefix string
	engine *engine
}

func New() *engine {
	e := &engine{router: newRouter()}
	e.routerGroup = &routerGroup{engine: e}
	e.middleware = &middleware{}

	return e
}

func (group *routerGroup) Use(middlewares ...handlerFunc) {
	group.engine.middleware.insert(parsePattern(group.prefix), 0, middlewares)
}

func (group *routerGroup) Group(prefix string) *routerGroup {
	e := group.engine
	newGroup := &routerGroup{
		prefix: group.prefix + prefix,
		engine: e,
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
	middlewares := engine.middleware.find(parsePattern(req.URL.Path), 0)

	middlewares = append(middlewares, func(ctx *Context) {
		engine.router.handle(ctx)
	})

	newContext(w, req, middlewares).Next()
}

func (engine *engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
