package main

import (
	"context"
	pb "go-grpc-echo/proto/echo"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/metadata"
)

func CallingUnaryEchoRPC(c pb.EchoClient, message string) {
	log.Println("---Calling Unary Echo RPC---")
	log.Println("sending message...")

	// metadata
	md := metadata.Pairs("timestamp", time.Now().Format(time.StampNano))

	// create context with metadata
	ctx := metadata.NewOutgoingContext(context.Background(), md)

	// ctx := context.Background()
	// req := &pb.EchoRequest{Message: message}

	res, err := c.UnaryEcho(ctx, &pb.EchoRequest{Message: message})
	if err != nil {
		log.Fatalf("Error while calling UnaryEcho RPC: %v", err)
	}

	log.Println("Response from UnaryEcho:")
	log.Println("response: ", res.Message)
}

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer conn.Close()

	client := pb.NewEchoClient(conn)
	CallingUnaryEchoRPC(client, "Hello from client!")
	CallingUnaryEchoRPC(client, "Hello again from client!")
}
