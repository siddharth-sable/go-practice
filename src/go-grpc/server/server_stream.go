package main

import (
	"log"
	"time"

	pb "github.com/siddharth-sable/go-grpc/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NamesList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got request with Names: %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Name: "Hello" + name,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
