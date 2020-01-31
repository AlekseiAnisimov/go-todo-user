package main

import (
	"context"
	"fmt"
	pb "github.com/AlekseiAnisimov/go-todo-user/proto"
	"google.golang.org/grpc"
	"time"
)

func main() {
	conn, err := grpc.Dial("localhost:20100", grpc.WithInsecure())
	if err != nil {
		fmt.Print("error 1")
	}
	defer conn.Close()

	c := pb.NewBearerAuthClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	var test *string

	resp, err := c.Check(ctx, &pb.Request{Message:test})
	if err != nil {
		fmt.Print("err request")
	}
	fmt.Print(resp.Message)
}
