package api

import (
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
	pb "grpc-server/api/product/v1"
)

func (server *Server) Hello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	if req.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "argument is required")
	}

	response := &pb.HelloResponse{
		Greet:    "hello " + req.Name,
		SendDate: timestamppb.Now(),
	}
	return response, nil
}
