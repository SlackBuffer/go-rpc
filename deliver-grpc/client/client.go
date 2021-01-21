package main

import (
	"context"
	pb "goplay/go-rpc/deliver-grpc/deliver"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:10000", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("dial err: %v", err)
	}
	defer conn.Close()
	dc := pb.NewDeliverClient(conn)
	ctx := context.Background()
	d, err := dc.Deliver(ctx)
	if err != nil {
		log.Fatalf("call Deliver error: %v", err)
	}
	d.Send(&pb.Msg{Msg: "1st"})
	for {
		msg, err := d.Recv()
		if err != nil {
			log.Println(err)
			break
		}
		log.Println(msg)
	}
}
