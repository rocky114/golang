package framework

import (
	"errors"
	"strings"
)

type router struct {
	roots    map[string]*node
	handlers map[string]handlerFunc
}

func newRouter() *router {
	return &router{
		roots:    make(map[string]*node),
		handlers: make(map[string]handlerFunc),
	}
}

func parsePattern(pattern string) []string {
	parts := make([]string, 0)

	if len(pattern) == 0 {
		return parts
	}

	for _, part := range strings.Split(pattern, "/") {
		if len(part) == 0 {
			continue
		}

		parts = append(parts, part)
	}

	return parts
}

func (r *router) addRoute(method string, pattern string, handler handlerFunc) {
	if _, ok := r.roots[method]; !ok {
		r.roots[method] = &node{}
	}

	key := method + "_" + pattern
	r.roots[method].insert(pattern, parsePattern(pattern), 0)
	r.handlers[key] = handler
}

func (r *router) getRoute(method, path string) (string, error) {
	if _, ok := r.roots[method]; !ok {
		return "", errors.New("route not found")
	}

	pattern, err := r.roots[method].find(parsePattern(path), 0)
	if err != nil {
		return "", err
	}

	return pattern, nil
}

func (r *router) handle(c *Context) {
	pattern, err := r.getRoute(c.Method, c.Path)
	if err != nil {
		_, _ = c.Writer.Write([]byte("not found"))
		return
	}

	key := c.Method + "_" + pattern
	r.handlers[key](c)
}
