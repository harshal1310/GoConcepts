package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "user-service/pb" // replace with actual path

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedServiceBServer
}

func (s *server) Greet(ctx context.Context, req *pb.Empty) (*pb.MessageResponse, error) {
	return &pb.MessageResponse{Message: "Hello from Service B!"}, nil
}

func (s *server) Add(ctx context.Context, req *pb.MessageRequest) (*pb.MessageResponse, error) {
	if req.GetMessage() == "" {
		return &pb.MessageResponse{Message: "Error: empty message"}, nil
	}
	return &pb.MessageResponse{Message: "Received: " + req.GetMessage()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterServiceBServer(grpcServer, &server{})
	fmt.Println("Service B listening on :50051")
	log.Fatal(grpcServer.Serve(lis))
}
