package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"

	"github.com/arminazimi/basic_grpc/server/echo"
)

func main() {
	ctx := context.Background()
	conn, err := grpc.DialContext(ctx, "localhost:8080", grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	e := echo.NewEchoServiceClient(conn)
	res, err := e.Echo(ctx, &echo.EchoRequest{Message: "Hello"})
	if err != nil {
		panic(err)
	}
	fmt.Println("got from server", res.Response)
}
