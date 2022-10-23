package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/rpc"
)

type Args struct {
	A, B int
}

type MuliplyService struct{}

func (t *MuliplyService) Do(args *Args, reply *int) error {
	log.Println("inside MuliplyService")
	*reply = args.A * args.B
	return nil
}

func main() {
	service := new(MuliplyService)
	rpc.Register(service)
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatal("listen error:", e)
	}

	fmt.Println("server is running...")
	http.Serve(l, nil)
}
