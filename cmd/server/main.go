package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/kohdice/go-grpc/pkg/grpc"
)

type server struct {
	pb.UnimplementedSampleServer
}

func (s *server) Hello(_ context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterSampleServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
