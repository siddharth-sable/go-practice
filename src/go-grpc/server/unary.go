package main

import (
	"context"

	pb "github.com/siddharth-sable/go-grpc/proto"
)

func (s *helloServer) SayHello(ctx context.Context, req *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{
		Name: "Hello",
	}, nil
}
