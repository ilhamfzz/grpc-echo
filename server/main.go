package main

import (
	"context"
	pb "go-grpc-echo/proto/echo"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

type server struct {
	pb.UnimplementedEchoServer
}

func (s *server) UnaryEcho(ctx context.Context, in *pb.EchoRequest) (*pb.EchoResponse, error) {
	log.Println("---Calling Unary Echo RPC---")
	log.Println("incoming message: ")
	log.Println("message: ", in.Message)

	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, status.Error(codes.DataLoss, "failed to get metadata")
	}

	if value, ok := md["timestamp"]; ok {
		log.Println("timestamp metadata: ")
		for i, v := range value {
			log.Printf("%v. %v\n", i+1, v)
		}
	}

	return &pb.EchoResponse{Message: in.Message}, nil
}

func main() {
	list, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	log.Println("Starting server on port :9090")
	if err := s.Serve(list); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
