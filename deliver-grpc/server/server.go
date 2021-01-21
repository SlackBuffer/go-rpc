package main

import (
	"fmt"
	pb "goplay/go-rpc/deliver-grpc/deliver"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc" // 要和 pb.go 引用同一个包
)

type deliverServer struct {
	pb.UnimplementedDeliverServer // 由 gRPC 生成
}

// Deliver 服务端实现
// * 参数是 pb.Deliver_DeliverServer
func (ds *deliverServer) Deliver(stream pb.Deliver_DeliverServer) error {
	go func() {
		for {
			time.Sleep(5 * time.Second)
			stream.Send(&pb.Msg{
				Msg: time.Now().Format(time.ANSIC),
			})
		}
	}()

	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		fmt.Println(in.Msg)
	}
}

func main() {
	// 1. 启动监听
	listener, err := net.Listen("tcp", ":10000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 2. 创建未注册具体服务的 gRPC server
	ds := grpc.NewServer()
	// 3. 用 gRPC 生成的方法注册带有具体实现的 server
	pb.RegisterDeliverServer(ds, &deliverServer{})
	if err := ds.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
