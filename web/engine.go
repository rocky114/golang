package web

type HandlerFunc func(ctx *Context)

type Engine struct {
	router *router
}
