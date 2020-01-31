package main

import (
	"fmt"
	"net"
	"context"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "github.com/AlekseiAnisimov/go-todo-user/proto"
)

type server struct {
}

func (s *server) Check(ctx context.Context, in *pb.Request) (*pb.Response, error) {
	var test *string
	return &pb.Response{Message: test}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":20100")
	if err != nil {
		fmt.Print("error 1")
	}
	fmt.Print("start")

	rpcserv := grpc.NewServer()

	pb.RegisterBearerAuthServer(rpcserv, &server{})
	reflection.Register(rpcserv)

	err = rpcserv.Serve(listener)
	if err != nil {
		fmt.Print("error 2")
	}
}