package grpc

import (
	"google.golang.org/grpc"
	grpcService "grpc-client/grpc/service/v1"
	"log"
)

var (
	ServiceClient grpcService.ServiceClient
)

func InitGrpcClient() {

	serviceConn, serviceErr := grpc.Dial("localhost:9001", grpc.WithInsecure())
	if serviceErr != nil {
		log.Fatal("cannot connect to products service", serviceErr)
	}
	ServiceClient = grpcService.NewServiceClient(serviceConn)
}
