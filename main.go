package main

import (
	"fmt"
	"go-example/framework"
	"net/http"
)

func main() {
	engine := framework.New()
	engine.GET("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello world"))
	})

	engine.GET("/hello", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello !!!"))
	})

	engine.GET("/hello/:name", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = writer.Write([]byte("hello :name !!!"))
	})

	err := engine.Run(":9999")
	fmt.Println(err)
}
