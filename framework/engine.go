package framework

import (
	"net/http"
)

type handlerFunc func(ctx *Context)

type routerGroup struct {
	prefix     string
	router     *router
	middleware *middleware
}

func New() *routerGroup {
	return &routerGroup{router: newRouter(), middleware: newMiddleware()}
}

func (group *routerGroup) Use(middlewares ...handlerFunc) {
	group.middleware.insert(parsePattern(group.prefix), 0, middlewares)
}

func (group *routerGroup) Group(prefix string) *routerGroup {
	newGroup := &routerGroup{
		prefix:     group.prefix + prefix,
		router:     group.router,
		middleware: group.middleware,
	}

	return newGroup
}

func (group *routerGroup) addRoute(method string, pattern string, handler handlerFunc) {
	group.router.addRoute(method, group.prefix+pattern, handler)
}

func (group *routerGroup) GET(pattern string, handler handlerFunc) {
	group.addRoute("GET", pattern, handler)
}

func (group *routerGroup) POST(pattern string, handler handlerFunc) {
	group.addRoute("POST", pattern, handler)
}

func (group *routerGroup) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	middlewares := group.middleware.find(parsePattern(req.URL.Path), 0)

	middlewares = append(middlewares, func(ctx *Context) {
		group.router.handle(ctx)
	})

	newContext(w, req, middlewares).Next()
}

func (group *routerGroup) Run(addr string) error {
	return http.ListenAndServe(addr, group)
}
