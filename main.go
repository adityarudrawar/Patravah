package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, req *pb.TestRequest) (*pb.TestResponse, error) {
  return &pb.TestResponse{Message: "Hello, " + req.Name}, nil
}

func main() {
  lis, err := net.Listen("tcp", ":50051")
  if err != nil {
    log.Fatalf("failed to listen: %v", err)
  }
  s := grpc.NewServer()
  pb.RegisterMyServiceServer(s, &server{})
  log.Println("Server started")
  if err := s.Serve(lis); err != nil {
    log.Fatalf("failed to serve: %v", err)
  }
}
