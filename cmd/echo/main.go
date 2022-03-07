package main

import (
	"context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	echo "grpc-example/gen/protos"
	"log"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	echoServer := NewEchoServer()

	echo.RegisterEchoServer(grpcServer, echoServer)
	reflection.Register(grpcServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}

func NewEchoServer() echo.EchoServer {
	return &echoServer{}
}

type echoServer struct {
	echo.UnimplementedEchoServer
}

func (e *echoServer) Echo(ctx context.Context, message *echo.Message) (*echo.Message, error) {
	return &echo.Message{Message: "Parroting: " + message.GetMessage()}, nil
}
