package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/kohdice/go-grpc/pkg/grpc"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not connect.: %v", err)
	}
	defer conn.Close()
	c := pb.NewSampleClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.Hello(ctx, &pb.HelloRequest{Name: "John Doe"})
	if err != nil {
		log.Fatalf("request failed.: %v", err)
	}

	log.Printf("Greeting: %s", r.GetMessage())
}
