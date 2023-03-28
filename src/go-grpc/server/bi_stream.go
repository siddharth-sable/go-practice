package main

import (
	"io"
	"log"

	pb "github.com/siddharth-sable/go-grpc/proto"
)

func (s *helloServer) SayHelloBidirectionalStreaming(stream pb.GreetService_SayHelloBidirectionalStreamingServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with name: %v", req.Message)

		res := &pb.HelloResponse{
			Name: "Hello" + req.Message,
		}
		if err := stream.Send(res); err != nil {
			return err
		}
	}
}
