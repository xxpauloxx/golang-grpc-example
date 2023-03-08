package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"

	pb "helloworld/helloworld"
)

type server struct{
    pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello, " + in.GetName() + "!"}, nil
}

func NewServer() *server {
	return &server{}
}

func main() {
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln("Failed to listen:", err)
	}

	s := grpc.NewServer()
    pb.RegisterGreeterServer(s, &server{})

    log.Println("Serving gRPC on 0.0.0.0:8080")
	log.Fatal(s.Serve(lis))
}
