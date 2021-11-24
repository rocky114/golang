package framework

import (
	"net/http"
	"strings"
)

type handlerFunc func(ctx *Context)

type engine struct {
	router *router
	*routerGroup
}

type routerGroup struct {
	prefix      string
	engine      *engine
	middlewares []map[string]handlerFunc
}

func New() *engine {
	e := &engine{router: newRouter()}
	e.routerGroup = &routerGroup{engine: e}

	return e
}

func (group *routerGroup) Use(middlewares ...handlerFunc) *routerGroup {
	for _, middleware := range middlewares {
		group.engine.routerGroup.middlewares = append(group.middlewares, map[string]handlerFunc{group.prefix: middleware})
	}

	return group
}

func (group *routerGroup) Group(prefix string) *routerGroup {
	engine := group.engine
	newGroup := &routerGroup{
		prefix:      group.prefix + prefix,
		middlewares: group.middlewares,
		engine:      engine,
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
	middlewares := make([]handlerFunc, 0)
	for _, item := range engine.routerGroup.middlewares {
		for prefix, middleware := range item {
			if strings.HasPrefix(req.URL.Path, prefix) {
				middlewares = append(middlewares, middleware)
			}
		}
	}

	middlewares = append(middlewares, func(ctx *Context) {
		engine.router.handle(ctx)
	})

	newContext(w, req, middlewares).Next()
}

func (engine *engine) Run(addr string) error {
	return http.ListenAndServe(addr, engine)
}
