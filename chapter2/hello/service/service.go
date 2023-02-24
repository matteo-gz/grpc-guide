package service

import (
	"context"
	"fmt"
	pb "hello/proto"
	"io"
	"log"
	"time"
)

type Server struct {
	pb.UnimplementedGreeterServer
}

func (s *Server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	fmt.Println("===SayHello")
	return &pb.HelloReply{Message: "Hello " + in.Name}, nil
}
func (s *Server) SayHelloServerStream(in *pb.HelloRequest, stream pb.Greeter_SayHelloServerStreamServer) error {
	fmt.Println("===SayHelloServerStream")
	fmt.Println(in.Name)
	var err error
	sw := []string{"1", "2", "3"}
	for _, v := range sw {
		err = stream.Send(&pb.HelloReply{
			Message: v,
		})
		if err != nil {
			fmt.Printf("SSayHello %v", err)
		}
	}
	fmt.Println("over")
	return nil
}
func (s *Server) SayHelloClientStream(stream pb.Greeter_SayHelloClientStreamServer) error {
	fmt.Println("===SayHelloClientStream")
	for {
		res, err1 := stream.Recv()
		if err1 == io.EOF {
			log.Printf("sStream break")
			fmt.Printf("server receive client close time %s\n", time.Now())
			time.Sleep(time.Second)
			err1 = stream.SendAndClose(&pb.HelloReply{Message: "sStream close"})
			fmt.Printf("server close time %s\n", time.Now())
			if err1 != nil {
				log.Printf("sStream %v", err1)
			}
			break
		} else if err1 != nil {
			log.Printf("sStream %v", err1)
			break
		}
		log.Printf("sStream res:%s \n", res.GetName())
	}
	return nil
}

func (s *Server) SayHelloBidirectionalStream(stream pb.Greeter_SayHelloBidirectionalStreamServer) error {
	fmt.Println("===SayHelloBidirectionalStream")
	for {
		res, err := stream.Recv()
		fmt.Printf("recv %s\n", res.GetName())
		if err == io.EOF {
			log.Printf("cs res:%v \n", err)
			break
		}
		time.Sleep(500 * time.Millisecond)
		err1 := stream.Send(&pb.HelloReply{
			Message: fmt.Sprintf("s:%v", res.GetName()),
		})
		if err1 != nil {
			log.Printf("cs send:%v \n", err1)
		}
	}
	log.Printf("cs leave \n")
	return nil
}
