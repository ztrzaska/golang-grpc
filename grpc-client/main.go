package main

import (
	"context"
	"fmt"
	"grpc-client/grpc"
	product "grpc-client/grpc/service/v1"
)

func main() {
	fmt.Println("Starting client")

	grpc.InitGrpcClient()
	helloResponse, err := grpc.ServiceClient.Hello(context.Background(), &product.HelloRequest{Name: "Bob"})

	if err != nil {
		fmt.Println("Error connecting to grpc server")
		return
	}

	fmt.Println(helloResponse)
}
