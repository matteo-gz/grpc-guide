package service

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/proto"
	pb "hello/proto"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.Test1) (*pb.HelloReply, error) {
	fmt.Println("in:", in)
	data, err := proto.Marshal(in)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(data)
	}
	return &pb.HelloReply{Message: "Hello " + fmt.Sprintf("%d", in.GetA())}, nil
}
