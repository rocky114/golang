package main

import (
	"fmt"
	"go-example/framework"
	"net/http"
)

func main() {
	engine := framework.New()
	engine.GET("/", func(ctx *framework.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	engine.GET("/hello", func(ctx *framework.Context) {
		ctx.String(http.StatusOK, "hello !!!")
	})

	engine.GET("/hello/:name", func(ctx *framework.Context) {
		ctx.String(http.StatusOK, "hello :name")
	})

	err := engine.Run(":9999")
	fmt.Println(err)
}
