package main

import (
	"errors"
	"fmt"
	"goplay/go-rpc/1-rpc/common"
	"log"
	"net"
	"net/rpc"
	"time"
)

type Arith int

func (t *Arith) Multiply(args *common.Args, reply *int) error {
	fmt.Printf("called: %s\n", time.Now().Format("2006-01-02 15:04:05"))
	time.Sleep(10 * time.Second)
	*reply = args.A * args.B
	return nil
}

func (t *Arith) Divide(args *common.Args, quo *common.Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B // 商
	quo.Rem = args.A % args.B // 余数
	return nil
}

func main() {
	rpc.Register(new(Arith))

	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal("Accept error")
		}
		rpc.ServeConn(conn) // 一次只能接收一个客户端的请求，处理完就返回
		// go rpc.ServeConn(conn) //可以启协程来接收并发请求
	}
}
