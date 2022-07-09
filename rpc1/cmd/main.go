package main

import (
	"fmt"
	"go-example/rpc1"
	"log"
	"net"
	"sync"
	"time"
)

func startServer(addr chan string) {
	l, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Fatalln("network error:", err)
	}

	log.Println("start rpc server on", l.Addr())
	addr <- l.Addr().String()

	rpc1.Accept(l)
}

func main() {
	log.SetFlags(0)

	addr := make(chan string)
	go startServer(addr)

	client, err := rpc1.Dial("tcp", <-addr)
	if err != nil {
		log.Fatalln("client err", err)
	}
	defer func() { _ = client.Close() }()

	time.Sleep(time.Second)

	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			args := fmt.Sprintf("rpc req %d", i)
			var reply string
			if err := client.Call("Foo.Sum", args, &reply); err != nil {
				log.Fatalln("call Foo.Sum error: ", err)
			}
			log.Println("reply: ", reply)
		}()
	}

	wg.Wait()
}
