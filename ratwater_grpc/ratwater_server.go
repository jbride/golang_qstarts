package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "ratwater.xyz/grpc/proto/ratwater"
)

var (
	//port = flag.Int("port", 5051, "The server port")
	port int = 5051
)

type server struct {
	pb.Ratwater_GRPCServer
}

func main() {
	//flag.Parse()
	socket := fmt.Sprintf("localhost:%d", port)
	lis, err := net.Listen("tcp", socket)
	if err != nil {
		log.Fatalf("failed to listen on %s : %v", socket, err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	pb.RegisterRatwater_GRPCServer(grpcServer, &server{})
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
