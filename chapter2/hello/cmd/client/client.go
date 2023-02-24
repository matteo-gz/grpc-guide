package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "hello/proto"
	"io"
	"log"
	"time"
)

const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewGreeterClient(conn)
	c1 := C{c}
	c1.unary()
	c1.sStream()
	c1.cStream()
	c1.csStream()
}

type C struct {
	pb.GreeterClient
}

func (c C) unary() {
	fmt.Println("===unary")
	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: "unary"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
func (c C) cStream() {
	fmt.Println("===ClientStream")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	stream, err := c.SayHelloClientStream(ctx)
	if err != nil {
		log.Printf("cStream %v", err)
		return
	}
	sw := []string{"client", "stream"}
	for _, v := range sw {
		err = stream.Send(&pb.HelloRequest{Name: v})
		if err != nil {
			log.Printf("cStream %v \n", err)
			return
		}
	}
	fmt.Printf("client before close time %s \n", time.Now())
	time.Sleep(time.Second)
	res, err1 := stream.CloseAndRecv()
	fmt.Printf("client close recv time %s \n", time.Now())
	fmt.Println(res, err1)
	fmt.Println("over")
}
func (c C) sStream() {
	fmt.Println("===ServerStream")
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	stream, err := c.SayHelloServerStream(ctx, &pb.HelloRequest{Name: "server stream"})
	if err != nil {
		log.Printf("sStream %v", err)
		return
	}
	for {
		res, err1 := stream.Recv()

		if err1 == io.EOF {
			log.Printf("sStream break")
			return
		} else if err1 != nil {
			log.Printf("sStream %v", err1)
			return
		}
		log.Printf("sStream res:%s", res.GetMessage())
	}
}

func (c C) csStream() {
	fmt.Println("===BidirectionalStream")
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()
	stream, err := c.SayHelloBidirectionalStream(ctx)
	if err != nil {
		log.Printf("csStream %v", err)
	}
	sw := []string{"1:1+1", "2:1+2", "3:1+5"}
	ch := make(chan bool)

	go func() {
		for {
			res, err1 := stream.Recv()
			if err1 == io.EOF {
				break
			}
			fmt.Printf("read%s\n", res.GetMessage())
		}
		ch <- true
	}()
	for _, v := range sw {
		time.Sleep(500 * time.Millisecond)
		err2 := stream.Send(&pb.HelloRequest{
			Name: v,
		})
		if err2 != nil {
			log.Printf("send err%v\n", err2)
		}
	}
	err3 := stream.CloseSend()
	log.Printf("close send err%v\n", err3)
	<-ch
	fmt.Println("client over")
	return
}
