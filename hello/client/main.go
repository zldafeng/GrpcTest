package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "GrpcTest/proto/hello"
	"context"
	"log"
)

const (
	Address = "127.0.0.1:8888"
)

func main() {
	// 连接
	conn,err :=grpc.Dial(Address,grpc.WithInsecure())

	if err !=nil{
		grpclog.Fatalln(err)
	}

	defer conn.Close()

	// 初始化客户端
	c := pb.NewHelloClient(conn)

	// 调用方法
	req := &pb.HelloRequest{Name: "gRPC"}
	r, err := c.SayHello(context.Background(),req)
	if err !=nil{
		grpclog.Fatalln(err)
	}
	log.Println(r.Message)
}