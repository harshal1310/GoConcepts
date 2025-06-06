// service-a.go
package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "user-service/pb" // replace with actual path

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatal("Could not connect:", err)
	}
	defer conn.Close()

	client := pb.NewServiceBClient(conn)

	// Call Greet
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp1, err := client.Greet(ctx, &pb.Empty{})
	if err != nil {
		log.Fatal("Greet error:", err)
	}
	fmt.Println("Greet Response:", resp1.GetMessage())

	// Call Add
	resp2, err := client.Add(ctx, &pb.MessageRequest{Message: "hello from A"})
	if err != nil {
		log.Fatal("Add error:", err)
	}
	fmt.Println("Add Response:", resp2.GetMessage())
}
