package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreetServiceServer
}

func (s *server) SayHello1(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	log.Printf("Received request from: %s", in.GetName())
	return &pb.HelloResponse{
		Message: fmt.Sprintf("Hello, %s!", in.GetName()),
	}, nil
}

func (s *server) SayHello2(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {

	log.Printf("Say Hello2")
	log.Printf("Received request from: %s", in.GetName())
	return &pb.HelloResponse{
		Message: fmt.Sprintf("It is Bmg's response %s", in.GetName()),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &server{})

	log.Println("gRPC Server listening on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
