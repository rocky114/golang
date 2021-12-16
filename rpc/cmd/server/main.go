package main

import (
	"fmt"
	"gitlab.sz.sensetime.com/kubersolver/api/student"
	"gitlab.sz.sensetime.com/kubersolver/internal/service"
	"google.golang.org/grpc"
	"log"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for _, addr := range []string{":5001", ":5002"} {
		wg.Add(1)

		go func(addr string) {
			defer func() {
				wg.Done()
			}()

			listen, err := net.Listen("tcp", addr)
			if err != nil {
				log.Fatalln("init listen fail: ", err.Error())
			}

			server := grpc.NewServer()
			student.RegisterStudentManagerServer(server, &service.StudentManager{Addr: addr})

			err = server.Serve(listen)
			if err != nil {
				log.Fatalln("failed to serve:", err.Error())
			}
		}(addr)
	}

	fmt.Println("start serve successfully")

	wg.Wait()
}
