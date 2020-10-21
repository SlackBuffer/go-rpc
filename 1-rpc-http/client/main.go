package main

import (
	"fmt"
	"goplay/go-rpc/1-rpc-http/common"
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.DialHTTP("tcp", "localhost:1234")
	if err != nil {
		log.Fatal("dial http error:", err)
	}

	// Synchronous call
	args := &common.Args{10, 8}
	var reply int
	err = client.Call("Arith.Multiply", args, &reply)
	if err != nil {
		log.Fatal("arith error:", err)
	}
	fmt.Printf("Arith: %d*%d=%d\n", args.A, args.B, reply)

	// Asynchronous call
	quotient := new(common.Quotient)
	divCall := client.Go("Arith.Divide", args, quotient, nil)
	replyCall := <-divCall.Done // will be equal to divCall
	if replyCall.Error != nil {
		log.Fatal("Arith.Divide error:", err)
	}
	fmt.Printf("Arith: (Arith.Divide(%v)): %v\n", args, replyCall.Reply)
	// check errors, print, etc.
}
