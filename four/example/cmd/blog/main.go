package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"github.com/pkg/errors"
	v1 "github.com/wangbinxiang/go-lesson-four/api/blog/v1"
)

func main() {
	service, err := InitUserService()
	if err != nil {
		log.Panicf("init service fail: %v", err)
	}

	g := grpc.NewServer()
	v1.RegisterUserServerServer(g, service)

	n, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal(errors.Wrap(err, "start server port :8080"))
	}
	log.Println("grpc server will listen :8080")
	g.Serve(n)
}
