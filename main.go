package main

import (
	"fmt"
	"go-example/framework"
	"log"
	"net/http"
	"time"
)

func main() {
	engine := framework.New()
	engine.Use(func(ctx *framework.Context) {
		t := time.Now()

		ctx.Next()

		log.Printf("%s in %v", ctx.Req.RequestURI, time.Since(t))
	})

	engine.GET("/", func(ctx *framework.Context) {
		ctx.String(http.StatusOK, "hello world")
	})

	engine.GET("/hello", func(ctx *framework.Context) {
		ctx.String(http.StatusOK, "hello !!!")
	})

	engine.GET("/hello/:name", func(ctx *framework.Context) {
		ctx.String(http.StatusOK, "hello :name")
	})

	v1 := engine.Group("/api")
	v1.GET("/version", func(ctx *framework.Context) {
		ctx.String(http.StatusOK, "v1")
	})

	v2 := v1.Group("/home")
	v2.GET("/admin", func(ctx *framework.Context) {
		ctx.String(http.StatusOK, "/api/home/admin")
	})

	err := engine.Run(":9999")
	fmt.Println(err)
}
