package main

import (
	"context"
	"fmt"
	"gitlab.sz.sensetime.com/kubersolver/api/student"
	"google.golang.org/grpc"
	"log"
)

func main() {
	conn, err := grpc.Dial("127.0.0.1:5001", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("connect err: %v", err)
	}

	defer conn.Close()

	request := student.StringMessage{
		Value: "hello world",
	}
	grpcClient := student.NewStudentManagerClient(conn)
	rsp, err := grpcClient.Echo(context.Background(), &request)

	if err != nil {
		log.Fatalf("request failed, err: %v", err)
	}

	fmt.Println(rsp.Value)
}
