package web

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer http.ResponseWriter
	Req    *http.Request
	Path   string
	Method string

	index       int
	middlewares []handlerFunc
}

func newContext(w http.ResponseWriter, req *http.Request, middlewares []handlerFunc) *Context {
	return &Context{
		Writer:      w,
		Req:         req,
		Path:        req.URL.Path,
		Method:      req.Method,
		middlewares: middlewares,
		index:       -1,
	}
}

func (c *Context) Next() {
	c.index++
	c.middlewares[c.index](c)
}

func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

func (c *Context) Status(code int) {
	c.Writer.WriteHeader(code)
}

func (c *Context) SetHeader(key, value string) {
	c.Writer.Header().Set(key, value)
}

func (c *Context) String(code int, data string) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/plain")
	_, _ = c.Writer.Write([]byte(data))
}

func (c *Context) JSON(code int, data interface{}) {
	c.Status(code)
	c.SetHeader("Content-Type", "application/json")

	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(data); err != nil {
		http.Error(c.Writer, err.Error(), http.StatusInternalServerError)
	}
}

func (c *Context) HTML(code int, data string) {
	c.Status(code)
	c.SetHeader("Content-Type", "text/html")
	_, _ = c.Writer.Write([]byte(data))
}
