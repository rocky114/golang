package main

import (
	"fmt"
	"go-example/web"
	"net/http"
)

func main() {
	fmt.Println("main")

	r := web.New()
	r.GET("/", func(ctx *web.Context) {
		ctx.HTML(http.StatusOK, "<h1>Hello</h1>")
	})

	r.GET("/hello", func(ctx *web.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Query("name"), ctx.Path)
	})

	r.GET("/hello/:name", func(ctx *web.Context) {
		ctx.String(http.StatusOK, "hello %s, you're at %s\n", ctx.Params["name"], ctx.Path)
	})

	r.GET("/assets/*filepath", func(ctx *web.Context) {
		ctx.JSON(http.StatusOK, web.H{"filepath": ctx.Params["filepath"]})
	})

	r.Run(":8888")
}
