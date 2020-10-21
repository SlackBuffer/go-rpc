package main

import (
	"errors"
	"goplay/go-rpc/1-rpc/common"
	"log"
	"net"
	"net/http"
	"net/rpc"
	"time"
)

type Arith int

func (t *Arith) Multiply(args *common.Args, reply *int) error {
	time.Sleep(time.Second)
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
	rpc.HandleHTTP()
	listener, err := net.Listen("tcp", ":1234")
	if err != nil {
		log.Fatal("listen error:", err)
	}
	http.Serve(listener, nil)
}
