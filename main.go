package main

import (
	"go-example/web"
	"log"
	"net/http"
	"time"
)

func onlyForV2() web.HandlerFunc {
	return func(ctx *web.Context) {
		t := time.Now()

		log.Printf("[%d] %s in %v for group v2", ctx.StatusCode, ctx.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := web.New()
	r.Use(web.Logger()) // global midlleware
	r.GET("/", func(c *web.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *web.Context) {
			// expect /hello/geektutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
