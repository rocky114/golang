package main

import (
	"fmt"
	"go-example/web"
	"net/http"
)

func main() {
	engine := web.New()
	engine.Use(func(ctx *web.Context) {
		fmt.Println("global middleware")
		ctx.Next()
	})

	engine.Use(func(ctx *web.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.String(http.StatusInternalServerError, fmt.Sprintf("internal server error: %s", err))
			}
		}()

		ctx.Next()
	})

	engine.GET("/", func(ctx *web.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	engine.GET("/panic", func(ctx *web.Context) {
		panic("server panic")
	})

	engine.GET("/hello", func(ctx *web.Context) {
		ctx.String(http.StatusOK, "hello !!!")
	})

	engine.GET("/hello/:name", func(ctx *web.Context) {
		ctx.String(http.StatusOK, "hello :name")
	})

	v1 := engine.Group("/api")
	v1.Use(func(ctx *web.Context) {
		fmt.Println("v1 middleware")
		ctx.Next()
	})
	v1.GET("/version", func(ctx *web.Context) {
		ctx.String(http.StatusOK, "v1")
	})

	v2 := v1.Group("/home")
	v2.GET("/admin", func(ctx *web.Context) {
		ctx.String(http.StatusOK, "/api/home/admin")
	})

	err := engine.Run(":9999")
	fmt.Println(err)
}
