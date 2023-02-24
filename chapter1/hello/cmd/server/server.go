package main

import (
	"fmt"
	"google.golang.org/grpc"
	pb "hello/proto"
	"hello/service"
	"net"
)

const (
	port = ":50051"
)

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		fmt.Printf("failed to listen: %v\n", err)
		return
	}
	srv := grpc.NewServer()
	pb.RegisterGreeterServer(srv, &service.Server{})
	fmt.Printf("listen:%s", port)
	if err := srv.Serve(lis); err != nil {
		fmt.Printf("failed to serve: %v\n", err)
		return
	}
}
