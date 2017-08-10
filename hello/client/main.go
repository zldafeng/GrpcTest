package main

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/grpclog"
	pb "GrpcTest/proto/hello"
	"context"
)

const (
	Address = "127.0.0.1:50052"
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
	reqBody :=new(pb.HelloRequest)
	reqBody.Name = "gRPC"

	r, err := c.Sayhello(context.Background(),reqBody)
	if err !=nil{
		grpclog.Fatalln(err)
	}

	grpclog.Println(r.Message)
}