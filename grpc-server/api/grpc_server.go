package api

import (
	pb "grpc-server/api/product/v1"
)

type Server struct {
	pb.UnimplementedServiceServer
}
