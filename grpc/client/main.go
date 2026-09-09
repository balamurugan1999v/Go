package main

import (
	"context"
	"log"
	"time"

	pb "grpc-demo/proto"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewGreetServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello1(ctx, &pb.HelloRequest{Name: "Balamurugan"})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}

	res1, err1 := client.SayHello2(ctx, &pb.HelloRequest{Name: "Bmg"})

	if err1 != nil {
		log.Fatal("Exception occured while sayHello2 called")
	}
	log.Printf("Server Response: %s", res.GetMessage())
	log.Printf("Server SayHello2 response: %s", res1.GetMessage())
}
