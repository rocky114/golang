package main

import (
	"context"
	"fmt"
	"gitlab.sz.sensetime.com/kubersolver/api/student"
	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
	"log"
	"time"
)

const (
	exampleScheme      = "http"
	exampleServiceName = "example.grpc.com"
)

var addrs = []string{"127.0.0.1:5001", "127.0.0.1:5002"}

type exampleResolverBuilder struct{}

func main() {
	resolver.Register(&exampleResolverBuilder{})

	roundRobinConn, err := grpc.Dial(
		fmt.Sprintf("%s:///%s", exampleScheme, exampleServiceName),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`),
		grpc.WithInsecure(),
		grpc.WithBlock(),
	)

	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer roundRobinConn.Close()

	for i := 0; i < 10; i++ {
		clientCall(roundRobinConn)
	}
}

func clientCall(conn grpc.ClientConnInterface) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	request := student.StringMessage{
		Value: "hello world",
	}
	grpcClient := student.NewStudentManagerClient(conn)
	rsp, err := grpcClient.Echo(ctx, &request)

	if err != nil {
		log.Fatalf("request failed, err: %v", err)
	}

	fmt.Println(rsp.Value)
}

func (*exampleResolverBuilder) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	r := &exampleResolver{
		target: target,
		cc:     cc,
		addrsStore: map[string][]string{
			exampleServiceName: addrs,
		},
	}
	r.start()
	return r, nil
}

func (*exampleResolverBuilder) Scheme() string { return exampleScheme }

type exampleResolver struct {
	target     resolver.Target
	cc         resolver.ClientConn
	addrsStore map[string][]string
}

func (r *exampleResolver) start() {
	addrStrs := r.addrsStore[r.target.Endpoint]
	addrs := make([]resolver.Address, len(addrStrs))
	for i, s := range addrStrs {
		addrs[i] = resolver.Address{Addr: s}
	}
	r.cc.UpdateState(resolver.State{Addresses: addrs})
}
func (*exampleResolver) ResolveNow(o resolver.ResolveNowOptions) {}
func (*exampleResolver) Close()                                  {}
