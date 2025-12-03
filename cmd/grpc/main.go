package main

import (
	"context"
	"log"
	"net"

	hellov1 "github.com/techinnovai/hello-go-api/gen/hello/v1"
	"google.golang.org/grpc"
)

type HelloServer struct {
	hellov1.UnimplementedHelloServiceServer
}

func (s *HelloServer) SayHello(ctx context.Context, req *hellov1.HelloRequest) (*hellov1.HelloResponse, error) {
	return &hellov1.HelloResponse{Message: "Hello World"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatal(err)
	}

	server := grpc.NewServer()
	hellov1.RegisterHelloServiceServer(server, &HelloServer{})

	log.Println("gRPC server running on :9090")
	if err := server.Serve(lis); err != nil {
		log.Fatal(err)
	}
}
