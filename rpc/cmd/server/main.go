package main

import (
	"context"
	"gitlab.sz.sensetime.com/kubersolver/api/student"
	"google.golang.org/grpc"
	"log"
	"net"
)

type studentManager struct {
	student.UnimplementedStudentManagerServer
}

func main() {
	listen, err := net.Listen("tcp", ":5001")
	if err != nil {
		log.Fatalln("init listen fail: ", err.Error())
	}

	server := grpc.NewServer()
	student.RegisterStudentManagerServer(server, &studentManager{})

	log.Println("gRPC server start...")
	err = server.Serve(listen)
	if err != nil {
		log.Fatalln("failed to serve:", err.Error())
	}
}

func (std *studentManager) Echo(ctx context.Context, req *student.StringMessage) (*student.StringMessage, error) {
	return &student.StringMessage{
		Value: req.GetValue(),
	}, nil
}
