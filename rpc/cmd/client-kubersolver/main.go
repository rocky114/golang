package main

import (
	"context"
	"github.com/sercand/kuberesolver/v3"
	"gitlab.sz.sensetime.com/kubersolver/api/student"
	"google.golang.org/grpc"
	"log"
	"net"
	"net/http"
)

func main() {
	kuberesolver.RegisterInClusterWithSchema("kubernetes")
	conn, err := grpc.Dial(
		"kubernetes://default/grpc-server:5001",
		grpc.WithBalancerName("round_robin"),
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Fatalf("connect err: %v", err)
	}

	defer conn.Close()

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		req := student.StringMessage{
			Value: "hello world",
		}
		grpcClient := student.NewStudentManagerClient(conn)
		rsp, err := grpcClient.Echo(context.Background(), &req)

		if err != nil {
			log.Fatalf("request failed, err: %v", err)
		}

		_, _ = writer.Write([]byte(rsp.Value))
		_, _ = writer.Write([]byte("\nip: " + getOutboundIp()))
	})

	err = http.ListenAndServe(":9999", nil)
	if err != nil {
		log.Fatalf("http serve error: %v", err)
	}
}

func getOutboundIp() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP.String()
}
