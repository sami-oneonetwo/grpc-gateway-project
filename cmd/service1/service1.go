package main

import (
	"context"
	"log"
	"net"

	pb "github.com/sami-oneonetwo/grpc-gateway-project/api/service1"
	"google.golang.org/grpc"
)

// Service1Server implements the Service1 gRPC service.
type Service1Server struct {
	pb.UnimplementedService1Server
}

// SayHello handles the SayHello gRPC call.
func (s *Service1Server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterService1Server(grpcServer, &Service1Server{})

	log.Println("Service1 gRPC server is running on port 50051...")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
