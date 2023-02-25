package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	pb "hello/proto"
	"log"
	"strconv"
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

	// Contact the server and print out its response.
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	a := int32(150)
	in := &pb.Test1{A: &a}
	data, err := proto.Marshal(in)
	if err != nil {
		fmt.Println(err)
	} else {
		tagNum := (1 << 3) | 0
		fmt.Println(tagNum)
		// 1001 0110  0000 0001
		fmt.Println(ten2bin(150))
		fmt.Println("bin", bin2ten("10010110"))
		fmt.Println("bin2", bin2ten("00000001"))
		// 的？这是一个二进制转换成十六进制的过程。首先，将二进制数每4位分割，分别得到1001、1011、0000、0001，然后将每一组二进制数转换成对应的十六进制，依此类推，最后得到9601。
		//fmt.Println(bin2ten(""))
		fmt.Println(ten2bin(int64(tagNum)))
		fmt.Println(bin2ten("1000"))
		fmt.Println("data", data)
	}
	r, err := c.SayHello(ctx, in)
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}

func bin2ten(binaryStr string) int64 {
	value, _ := strconv.ParseInt(binaryStr, 2, 64)
	return value
}
func ten2bin(num int64) string {
	binary := strconv.FormatInt(int64(num), 2)
	return binary
}
