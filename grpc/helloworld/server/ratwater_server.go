package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "ratwater.xyz/grpc/helloworld/proto"
	"ratwater.xyz/mod/ratwater"
)

type server struct {
	pb.UnimplementedRatwater_GRPCServer
}

// Add a minimum, just implement the rpc functions defined in the proto file
func (s *server) Hello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
	modResponse := ratwater.Hello(in.String())
	fmt.Printf("UnaryHello() inbound param = %s;\tresponse = %s\n", in.String(), modResponse)
	return &pb.HelloResponse{SResponse: modResponse}, nil
}

func main() {
	port := flag.Int("port", 5051, "The server port")
	flag.Parse()

	socket := fmt.Sprintf("localhost:%d", *port)
	lis, err := net.Listen("tcp", socket)
	if err != nil {
		log.Fatalf("failed to listen on %s : %v", socket, err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterRatwater_GRPCServer(grpcServer, &server{})
	fmt.Printf("About to start grpc server on: %s\n", socket)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v\n", err)
	}
}
