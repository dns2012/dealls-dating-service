package config

import (
	"flag"
	"google.golang.org/grpc"
	"log"
	"net"
)

var (
	GrpcOption         []grpc.ServerOption
	GrpcServerEndpoint = flag.String("grpc-server", ":8080", "grpc server")
	GrpcServer         = grpc.NewServer(GrpcOption...)
)

func grpcOption() []grpc.ServerOption {
	return []grpc.ServerOption{}
}

func RunGrpcServer() {
	server, err := net.Listen("tcp", *GrpcServerEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	go func() {
		GrpcServer.Serve(server)
	}()
}
