package main

import (
	"context"
	"fmt"
	"net"

	"google.golang.org/grpc"

	"github.com/arminazimi/basic_grpc/server/echo"
)

type EchoService struct{}

func (e *EchoService) Echo(ctx context.Context, req *echo.EchoRequest) (*echo.EchoResponse, error) {
	return &echo.EchoResponse{
		Response: "echo: " + req.Message,
	}, nil

}

func main() {
	lst, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	srv := &EchoService{}
	echo.RegisterEchoServiceServer(s, srv)
	fmt.Println("is serving")
	err = s.Serve(lst)
	if err != nil {
		panic(err)
	}

}
